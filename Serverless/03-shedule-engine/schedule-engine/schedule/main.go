package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ncostamagna/stack_develop_go/Serverless/03-shedule-engine/schedule-engine/schedule/engine"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

func main() {

	// Setup Logger
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", os.Getenv("APP_NAME"),
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	// // Load .env
	// err := godotenv.Load()
	// if err != nil {
	// 	level.Error(logger).Log("exit", err)
	// 	os.Exit(-1)
	// }

	// Services
	// ctx := context.Background()
	var es engine.Service
	es = engine.NewService(logger)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println("Initial app")
	// var httpAddr = flag.String("http", ":"+os.Getenv("APP_PORT"), "http listen address")
	// mux := http.NewServeMux()
	ce := engine.MakeEndpoints(es)

	handler := engine.NewAWSLambdaHandler(ce)
	lambda.StartHandler(handler)
	// mux.Handle("/v1/schedule/", engine.MakeHandler(ctx, ce))

	// http.Handle("/", accessControl(mux))
	// http.Handle("/metrics", promhttp.Handler())

	// go func() {
	// 	fmt.Println("listening on port", *httpAddr)
	// 	errs <- http.ListenAndServe(*httpAddr, nil)
	// }()

	level.Error(logger).Log("exit", <-errs)
}

// func accessControl(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

// 		if r.Method == "OPTIONS" {
// 			return
// 		}

// 		h.ServeHTTP(w, r)
// 	})
// }

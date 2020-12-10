package engine

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/transport/awslambda"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHandler -
func MakeHandler(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()

	scheduleHandler := httptransport.NewServer(
		endpoints.Schedule,
		decodeScheduleReq,
		encodeResponse,
	)
	r.Handle("/v1/schedule/", scheduleHandler).Methods("POST")

	return r
}

// NewAWSLambdaHandler -
func NewAWSLambdaHandler(endpoints Endpoints) *awslambda.Handler {
	fmt.Println("Test")
	return awslambda.NewHandler(endpoints.Schedule, decodeAWSLambdaRequest, encodeAWSLambdaResponse)
}

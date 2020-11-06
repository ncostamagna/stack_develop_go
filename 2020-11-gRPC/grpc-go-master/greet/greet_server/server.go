package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/katoozi/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet Function Was invoked with: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := &greetpb.GreetResponse{
		Result: fmt.Sprintf("Hello %s\n", firstName),
	}
	return result, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes Function Was invoked with: %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	fmt.Println("LongGreet Function Was invoked with a streaming request")

	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			res := &greetpb.LongGreetResponse{
				Result: result,
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			log.Fatalf("Error while receive client stream: %v", err)
		}
		firstName := req.GetGreeting().GetFirstName()
		result += "Hello " + firstName + "! "
	}
}

func (*server) GreetEveryone(stream greetpb.GreetService_GreetEveryoneServer) error {
	fmt.Println("GreetEveryone Function Was invoked with a streaming request")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
			return err
		}
		firstName := req.Greeting.GetFirstName()
		result := &greetpb.GreetEveryoneResponse{
			Result: "Hello " + firstName,
		}
		err = stream.Send(result)
		if err != nil {
			log.Fatalf("Error while sending response: %v", err)
			return err
		}
	}
}

func (*server) GreetWithDeadline(ctx context.Context, req *greetpb.GreetWithDeadlineRequest) (*greetpb.GreetWithDeadlineResponse, error) {
	fmt.Printf("GreetWithDeadline Function Was invoked with: %v\n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			// the client canceled the request
			fmt.Println("Client Canceld the Request!")
			return nil, status.Error(
				codes.Canceled,
				"Client canceld the request!",
			)
		}
		time.Sleep(1 * time.Second)
	}
	firstName := req.GetGreeting().GetFirstName()
	result := &greetpb.GreetWithDeadlineResponse{
		Result: fmt.Sprintf("Hello %s\n", firstName),
	}
	return result, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	tls := true
	opts := []grpc.ServerOption{}
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Error while loading ssl: %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serv: %v", err)
	}
}

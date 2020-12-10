package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleHello(name string) string {
	return name
}

func main() {
	fmt.Println("Example")
	lambda.Start(handleHello)
}

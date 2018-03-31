package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

// DemoEvent input
type DemoEvent struct {
	ID int `json:"id"`
}

// DemoResult output
type DemoResult struct {
	Message string `json:"result"`
}

// Handler lamda entry point
func Handler(event DemoEvent) (DemoResult, error) {
	return DemoResult{Message: fmt.Sprintf("Received ID [%d]", event.ID)}, nil
}

func main() {
	lambda.Start(Handler)
}

package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

func Handler(request Request) (Response, error) {
	fmt.Println(request)
	return Response{
		Message: "Hola mundo",
		Ok:      true,
	}, nil
}

var (
	start = lambda.Start
)

func main() {
	start(Handler)
}

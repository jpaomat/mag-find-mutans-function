package main

import (
	"fmt"

	"mag-stadistics-dna-processed-function/src/config/response"
	stadistics "mag-stadistics-dna-processed-function/src/controllers/stadisticsController"

	"github.com/aws/aws-lambda-go/lambda")

type Request struct {
	ID float64 `json:"id"`
}

var (
	start = lambda.Start
)

func Handler(request Request) (*response.Response, error) {
	fmt.Println(request)

	resp := stadistics.GetStadisticsDnaProcessed()
	fmt.Println(resp)
	
	return &response.Response{
		Message:    "OK",
		StatusCode: 200,
		Body: response.BodyStruct{
			Count_mutant_dna: 0,
			Count_human_dna:  0,
			Ratio:            0,
		},
	}, nil
}

func main() {
	start(Handler)
}

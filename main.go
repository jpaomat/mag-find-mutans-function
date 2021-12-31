package main

import (
	"fmt"
	"mag-stadistics-dna-processed-function/src/config/response"
	stadistics "mag-stadistics-dna-processed-function/src/controllers/stadisticsController"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {}

var (
	start = lambda.Start
)

func Handler(request events.APIGatewayProxyRequest) (*response.Response, error) {
	fmt.Println("Log 1 (CL 18-main) -> Input data to mag-stadistics-dna-proccesed-function lambda: ", request)

	respStadistics := stadistics.GetStadisticsDnaProcessed()

	fmt.Println("Log 7 (CL 22-main) -> Response to mag-stadistics-dna-proccesed-function lambda: ", respStadistics)
	return &response.Response{
		Message:    "OK",
		StatusCode: 200,
		Body: response.BodyStruct{
			Count_mutant_dna: respStadistics.Count_mutant_dna,
			Count_human_dna: respStadistics.Count_human_dna,
			Ratio: respStadistics.Ratio,
		},
	}, nil
}

func main() {
	start(Handler)
}

package main

import (
	"fmt"
	"errors"
	"mag-stadistics-dna-processed-function/src/config/response"
	stadistics "mag-stadistics-dna-processed-function/src/controllers/stadisticsController"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"mag-stadistics-dna-processed-function/src/utils"
)

type Request struct {}

var (
	start = lambda.Start
)
const (
	errDefault = "Error with the connection"
)

var (
	logger         = utils.Logger
)

func Handler(request events.APIGatewayProxyRequest) (*response.Response, error) {
	fmt.Println("Log 1 (CL 18-main) -> Input data to mag-stadistics-dna-proccesed-function lambda: ", request)

	respStadistics, errStadistics:= stadistics.GetStadisticsDnaProcessed()
	if errStadistics != nil {
		return nil, errors.New(errStadistics.Message)
	}

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

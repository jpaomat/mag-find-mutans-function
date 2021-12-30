package utils

import (
	"encoding/json"
	"fmt"
	"log"
	credentialsdb "mag-stadistics-dna-processed-function/src/config/credentialsDb"
	errormanager "mag-stadistics-dna-processed-function/src/config/errorManager"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

var (
	region       string = "us-east-1"
	versionStage string = "AWSCURRENT"
	errDefault   string = "Error with the connection"
)

func Logger(errorIn string, errorOut string, status int, errorException string) *errormanager.ErrorManager {
	errorCode := time.Now().Unix()
	log.Printf("ErrorInput: %s:: {%d}", errorIn, errorCode)
	log.Printf("ErrorOutput: %s:: {%d}", errorOut, errorCode)
	if errorException != "" {
		log.Printf("ErrorException: %s:: {%d}", errorException, errorCode)
	}
	return &errormanager.ErrorManager{
		Message: fmt.Sprintf("%s :: %d", errorOut, errorCode),
		Status:  status,
	}
}

func GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, *errormanager.ErrorManager) {
	svc := secretsmanager.New(
		session.New(),
		aws.NewConfig().WithRegion(region),
	)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String(versionStage),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return nil, Logger("Error to connect with DB", errDefault, http.StatusInternalServerError, err.Error())
	}
	return result, nil
}

func SecretManagerDB(secretValue *secretsmanager.GetSecretValueOutput) *credentialsdb.SecretData {
	var secretString string
	if secretValue.SecretString != nil {
		secretString = *secretValue.SecretString
	}

	var secretDatabase *credentialsdb.SecretData
	err := json.Unmarshal([]byte(secretString), &secretDatabase)
	if err != nil {
		log.Print("Error to get DB credentials")
		panic(err.Error())
	}

	return secretDatabase
}

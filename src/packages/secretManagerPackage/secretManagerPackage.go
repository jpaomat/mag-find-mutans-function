package secretmanagerpackage

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type Secret struct {
	secretName string
}

func New(sn string) Secret {
	mySecretName := Secret{secretName: sn}
	return mySecretName
}

type SecretData struct {
	Hostname string  `json:"host"`
	UserName string  `json:"username"`
	Password string  `json:"password"`
	Database string  `json:"dbInstanceIdentifier"`
	Port     float64 `json:"port"`
}

var (
	secretName   string = ""
	region       string = "us-east-1"
	versionStage string = "AWSCURRENT"
)

func (mySecretName Secret) GetSecretVal() SecretData {

	secretName = mySecretName.secretName
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
		panic(err.Error())
	}

	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	}

	var secretData SecretData
	err = json.Unmarshal([]byte(secretString), &secretData)
	if err != nil {
		panic(err.Error())
	}

	return secretData
}

package secretmanagerpackage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SecretName struct {
	secretName string
}

func New(sn string) SecretName {
	mySecretName := SecretName{secretName: sn}
	return mySecretName
}

// type SecretData struct {
// 	Hostname string  `json:"host"`
// 	UserName string  `json:"username"`
// 	Password string  `json:"password"`
// 	Database string  `json:"dbInstanceIdentifier"`
// 	Port     float64 `json:"port"`
// }

var (
	secretName   string = ""
	region       string = "us-east-1"
	versionStage string = "AWSCURRENT"
)

func (mySecretName SecretName) GetSecretVal() *secretsmanager.GetSecretValueOutput {

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

	// var secretString string
	// if result.SecretString != nil {
	// 	secretString = *result.SecretString
	// }

	// var secretDatabase SecretData
	// err = json.Unmarshal([]byte(secretString), &secretDatabase)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// return secretDatabase
	return result
}

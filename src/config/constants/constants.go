package constants

import (
	"fmt"
	"mag-stadistics-dna-processed-function/src/utils"
	"os"
)

type ConstantsDb struct {
	maxOpenDbConn string
	masIdleDbConn string
	maxDbLifetime string
}

func BuildConstants() *ConstantsDb {
	return &ConstantsDb{
		maxOpenDbConn: os.Getenv("MAXOPENDBCONN"),
		masIdleDbConn: os.Getenv("MASIDLEDBCONN"),
		maxDbLifetime: os.Getenv("MAXDBLIFETIME"),
	}
}

func GetMysqlConnectionString() string {
	secretCredentialsDb, err := utils.GetSecretValue(os.Getenv("SECRETDBNAME"))
	credentialsDb := utils.SecretManagerDB(secretCredentialsDb)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", credentialsDb.UserName, credentialsDb.Password, credentialsDb.Hostname, credentialsDb.Port, credentialsDb.Database)
	if err != nil {
		panic(err)
	}
	return connectionString
}

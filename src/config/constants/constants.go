package constants

import (
	"fmt"
	"mag-stadistics-dna-processed-function/src/utils"
	"strconv"
	"time"
)

var (
	castToInt     = strconv.Atoi
	parseDuration = time.ParseDuration
)

func GetMysqlConnectionString() string {
	secretCredentialsDb, err := utils.GetSecretValue("/rds_db/mysql")
	credentialsDb := utils.SecretManagerDB(secretCredentialsDb)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", credentialsDb.UserName, credentialsDb.Password, credentialsDb.Hostname, credentialsDb.Port, credentialsDb.Database)
	if err != nil {
		panic(err)
	}
	return connectionString
}

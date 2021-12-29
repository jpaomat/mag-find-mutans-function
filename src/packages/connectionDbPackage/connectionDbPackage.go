package connectiondbpackage

import (
	"encoding/json"
	"fmt"

	sm "mag-stadistics-dna-processed-function/src/packages/secretManagerPackage"

	"database/sql"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
)

type SecretNameDb struct {
	secretName string
}

func NewValue(sn string) SecretNameDb {
	secretNameDb := SecretNameDb{secretName: sn}
	return secretNameDb
}

type SecretData struct {
	Hostname string `json:"host"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Database string `json:"dbInstanceIdentifier"`
	Port     int    `json:"port"`
}

func (secretNameDb SecretNameDb) GetConnectionDb() *sql.DB {
	getSecretDb := sm.New(secretNameDb.secretName)
	secretDataDb := getSecretDb.GetSecretVal()
	credentialsDb := SecretManager(secretDataDb)
	sqlConnection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", credentialsDb.UserName, credentialsDb.Password, credentialsDb.Hostname, credentialsDb.Port, credentialsDb.Database)
	fmt.Println("sqlConnection", sqlConnection)
	db, err := sql.Open("mysql", sqlConnection)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing

	return db
}

func SecretManager(secretValue *secretsmanager.GetSecretValueOutput) SecretData {
	var secretString string
	if secretValue.SecretString != nil {
		secretString = *secretValue.SecretString
	}

	var secretDatabase SecretData
	err := json.Unmarshal([]byte(secretString), &secretDatabase)
	if err != nil {
		panic(err.Error())
	}

	return secretDatabase
}

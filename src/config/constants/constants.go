package constants

import (
	"fmt"
	"log"
	"mag-stadistics-dna-processed-function/src/utils"
	"os"
	"strconv"
	"time"
)

type Constants struct {
	port                  string
	mySQLConnectionString string
	maxOpenDbConn         string
	masIdleDbConn         string
	maxDbLifetime         string
}

const (
	maxOpenDbConnDefault = 10
	masIdleDbConnDefault = 5
	maxDbLifetimeDefault = 5 * time.Minute

	ACCESS_ID = "accessId"
)

var (
	castToInt     = strconv.Atoi
	parseDuration = time.ParseDuration
)

func BuildConstants() *Constants {
	return &Constants{
		mySQLConnectionString: "",
		maxOpenDbConn:         os.Getenv("MAXOPENDBCONN"),
		masIdleDbConn:         os.Getenv("MASIDLEDBCONN"),
		maxDbLifetime:         os.Getenv("MAXDBLIFETIME"),
	}
}

func (c *Constants) GetPort() string {
	return c.port
}
func (c *Constants) GetMaxOpenDbConn() int {
	maxOpenDbConn, err := castToInt(c.maxOpenDbConn)
	if err != nil {
		log.Print("not maxOpenDbConn from env (default: 10)")
		maxOpenDbConn = maxOpenDbConnDefault
	}
	return maxOpenDbConn
}
func (c *Constants) GetMasIdleDbConn() time.Duration {
	masIdleDbConn, err := parseDuration(c.masIdleDbConn)
	if err != nil {
		log.Print("not masIdleDbConn from env (default: 5ms)")
		masIdleDbConn = time.Duration(masIdleDbConnDefault)
	}
	return masIdleDbConn
}
func (c *Constants) GetMaxDbLifetime() time.Duration {
	maxDbLifetime, err := parseDuration(c.maxDbLifetime)
	if err != nil {
		log.Print("not maxDbLifetime from env (default: 5m)")
		maxDbLifetime = time.Duration(maxDbLifetimeDefault)
	}
	return maxDbLifetime

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

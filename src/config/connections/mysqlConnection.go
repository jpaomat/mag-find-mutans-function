package connections

import (
	"database/sql"
	errormanager "mag-stadistics-dna-processed-function/src/config/errorManager"
	"mag-stadistics-dna-processed-function/src/utils"
	"net/http"
	"time"
)

const (
	errDefault = "Error with the connection"
)

type MySQLConnection struct { // info para lograr la conexion
	connectionString string
	maxOpenDbConn    int
	maxIdleDbConn    time.Duration
	maxDbLifeTime    time.Duration
}

var (
	openConnection = sql.Open
	logger         = utils.Logger
)

func BuildMySQLConnection(connectionString string, maxOpenDbConn int, maxIdleDbConn time.Duration, maxDbLifeTime time.Duration) *MySQLConnection {
	return &MySQLConnection{
		connectionString: connectionString,
		maxOpenDbConn:    maxOpenDbConn,
		maxIdleDbConn:    maxIdleDbConn,
		maxDbLifeTime:    maxDbLifeTime * time.Minute,
	}
}

// func GetConnectDBMysql(connectionString string, maxOpenDbConn int, maxIdleDbConn time.Duration, maxDbLifeTime time.Duration) (*sql.DB, *errormanager.ErrorManager) {
// 	if connectionString == "" {
// 		return nil, logger("Missing the string connection", errDefault, http.StatusInternalServerError, "")
// 	}
// 	db, err := openConnection("mysql", connectionString)
// 	if err != nil {
// 		return nil, logger("Error to connect with DB", errDefault, http.StatusInternalServerError, err.Error())
// 	}
// 	db.SetConnMaxLifetime(maxDbLifeTime)
// 	db.SetMaxOpenConns(maxOpenDbConn)
// 	db.SetMaxIdleConns(maxOpenDbConn)
// 	return db, nil
// }

func GetConnectDBMysql(connectionString string) (*sql.DB, *errormanager.ErrorManager) {
	if connectionString == "" {
		return nil, logger("Missing the string connection", errDefault, http.StatusInternalServerError, "")
	}
	db, err := openConnection("mysql", connectionString)
	if err != nil {
		return nil, logger("Error to connect with DB", errDefault, http.StatusInternalServerError, err.Error())
	}
	db.SetConnMaxLifetime(10)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	return db, nil
}

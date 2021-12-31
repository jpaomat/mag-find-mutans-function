package connections

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
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

func (mConn *MySQLConnection) GetConnectDBMysql() (*sql.DB, *errormanager.ErrorManager) {
	if mConn.connectionString == "" {
		return nil, logger("Missing the string connection", errDefault, http.StatusInternalServerError, "")
	}
	db, err := openConnection("mysql", mConn.connectionString)
	if err != nil {
		return nil, logger("Error to connect with DB", errDefault, http.StatusInternalServerError, err.Error())
	}
	db.SetConnMaxLifetime(mConn.maxDbLifeTime)
	db.SetMaxOpenConns(mConn.maxOpenDbConn)
	db.SetMaxIdleConns(mConn.maxIdleConns)
	return db, nil
}

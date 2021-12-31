package connections

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	errormanager "mag-stadistics-dna-processed-function/src/config/errorManager"
	"mag-stadistics-dna-processed-function/src/utils"
	"net/http"
)

const (
	errDefault = "Error with the connection"
)

type MySQLConnection struct { // info para lograr la conexion
	connectionString string
}

var (
	openConnection = sql.Open
	logger         = utils.Logger
)

func BuildMySQLConnection(connectionString string) *MySQLConnection {
	return &MySQLConnection{
		connectionString: connectionString,
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
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5)
	return db, nil
}

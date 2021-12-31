package getDnaSequencesService

import (
	"database/sql"
	"mag-stadistics-dna-processed-function/src/config/connections"
	"mag-stadistics-dna-processed-function/src/config/constants"
	errormanager "mag-stadistics-dna-processed-function/src/config/errorManager"
	"mag-stadistics-dna-processed-function/src/utils"
	"net/http"
)

const (
	errDefault = "Error with the connection"
)

var (
	logger         = utils.Logger
)

func GetDataDnaSequences() (*sql.Rows, *errormanager.ErrorManager) {
	connectionDb, errDto := connections.BuildMySQLConnection(
		constants.GetMysqlConnectionString(),
		constants.GetMaxDbLifetime(),
		constants.GetMaxOpenDbConn(),
		constants.GetMasIdleDbConn(),
	).GetConnectDBMysql()
	if errDto != nil {
		return nil, errDto
	}
	defer connectionDb.Close()

	rows, err := connectionDb.Query("SELECT MUTANT FROM mutants_general.DNA_VERIFICATION_MUTANTS")
	if err != nil {
		return nil, logger("Error to get data DB", errDefault, http.StatusInternalServerError, err.Error())
	}
	return rows, nil
}
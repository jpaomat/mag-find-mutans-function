package getDnasequencesService

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
	connectionDb, errDto := connections.GetConnectDBMysql(
			constants.GetMysqlConnectionString(),
		)
	if errDto != nil {
		panic(errDto)
	}
	defer connectionDb.Close()

	rows, err := connectionDb.Query("SELECT MUTANT FROM mutants_general.DNA_VERIFICATION_MUTANTS")
	// if there is an error inserting, handle it
	if err != nil {
		return nil, logger("Error to get data DB", errDefault, http.StatusInternalServerError, err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer rows.Close()
	return rows, nil
}
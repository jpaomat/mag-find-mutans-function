package stadisticscontroller

import (
	// "database/sql"
	// "mag-stadistics-dna-processed-function/src/config/connections"
	// "mag-stadistics-dna-processed-function/src/config/constants"
	"fmt"
	"mag-stadistics-dna-processed-function/src/config/response"
	"mag-stadistics-dna-processed-function/src/services/getDnaSequencesService"
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

func GetStadisticsDnaProcessed() (*response.BodyStruct, *errormanager.ErrorManager) {
	// connectionDb, errDto := loadConnection()
	// if errDto != nil {
	// 	panic(errDto)
	// }
	// defer connectionDb.Close()

	// rows, err := connectionDb.Query("SELECT MUTANT FROM mutants_general.DNA_VERIFICATION_MUTANTS")
	// // if there is an error inserting, handle it
	// if err != nil {
	// 	panic(err.Error())
	// }
	rows, errData := getDnaSequencesService.GetDataDnaSequences()
	if errData != nil {
		return nil, errData
	}
	// be careful deferring Queries if you are using transactions
	defer rows.Close()
	count_mutant_dna:=0
	count_human_dna:=0
	for rows.Next() {
		var mutant string
        errorScan := rows.Scan(&mutant)
        if errorScan != nil {
			return nil, logger("Error to get data DB", errDefault, http.StatusInternalServerError, errorScan.Error())
        }
		if mutant == "1" {
			count_mutant_dna ++
		}
		count_human_dna++
    }
	ratio := float64(count_mutant_dna)/float64(count_human_dna)
	return &response.BodyStruct{
		Count_mutant_dna: count_mutant_dna,
		Count_human_dna: count_human_dna,
		Ratio: fmt.Sprintf("%.1f", ratio),
	}, nil
}

// func loadConnection() (*sql.DB, *errormanager.ErrorManager) {
// 	connectionDb, errDto := connections.GetConnectDBMysql(
// 			constants.GetMysqlConnectionString(),
// 		)
// 		if errDto != nil {
// 			panic(errDto)
// 		}
// 	return connectionDb, errDto
// }

// func getDataDnaSequences() (*sql.Rows, *errormanager.ErrorManager) {
// 	connectionDb, errDto := connections.GetConnectDBMysql(
// 			constants.GetMysqlConnectionString(),
// 		)
// 	if errDto != nil {
// 		panic(errDto)
// 	}
// 	defer connectionDb.Close()

// 	rows, err := connectionDb.Query("SELECT MUTANT FROM mutants_general.DNA_VERIFICATION_MUTANTS")
// 	// if there is an error inserting, handle it
// 	if err != nil {
// 		return nil, logger("Error to get data DB", errDefault, http.StatusInternalServerError, err.Error())
// 	}
// 	// be careful deferring Queries if you are using transactions
// 	defer rows.Close()
// 	return rows, nil
// }

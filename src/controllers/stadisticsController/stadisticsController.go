package stadisticscontroller

import (
	"database/sql"
	"fmt"
	"mag-stadistics-dna-processed-function/src/config/connections"
	"mag-stadistics-dna-processed-function/src/config/constants"
	"mag-stadistics-dna-processed-function/src/config/response"
	errormanager "mag-stadistics-dna-processed-function/src/config/errorManager"
)

func GetStadisticsDnaProcessed() *response.BodyStruct {
	// // clsConnectiob := db.NewValue("/rds_db/mysql")
	connectionDb, errDto := loadConnection()
	if errDto != nil {
		panic(errDto)
	}
	defer connectionDb.Close()
	fmt.Println("connection DB", connectionDb)

	rows, err := connectionDb.Query("SELECT MUTANT FROM mutants_general.DNA_VERIFICATION_MUTANTS")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Data table", rows)
	// be careful deferring Queries if you are using transactions
	defer rows.Close()
	count_mutant_dna:=0
	count_human_dna:=0
	for rows.Next() {
		var mutant string
        err = rows.Scan(&mutant)
        if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
        }
		if mutant == "1" {
			count_mutant_dna ++
		}
		count_human_dna++
    }
	ratio := float64(count_mutant_dna)/float64(count_human_dna)
	// ratio = (count_human_dna.compareTo(BigDecimal.ZERO) != 0) ? count_mutant_dna.divide(count_human_dna, 2, RoundingMode.UNNECESSARY) : new BigDecimal("0.00");
	err = rows.Err()
	return &response.BodyStruct{
		Count_mutant_dna: count_mutant_dna,
		Count_human_dna: count_human_dna,
		Ratio: Sprintf("%.1f"),
	}
}

func loadConnection() (*sql.DB, *errormanager.ErrorManager) {
	fmt.Println("loadconnection")
	connectionDb, errDto := connections.GetConnectDBMysql(
			constants.GetMysqlConnectionString(),
		)
		if errDto != nil {
			panic(errDto)
		}
	return connectionDb, errDto
}

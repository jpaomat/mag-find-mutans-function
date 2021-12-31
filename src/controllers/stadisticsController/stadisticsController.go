package stadisticscontroller

import (
	"database/sql"
	"fmt"
	"mag-stadistics-dna-processed-function/src/config/connections"
	"mag-stadistics-dna-processed-function/src/config/constants"
	errormanager "mag-stadistics-dna-processed-function/src/config/errorManager"
	"strconv"
	"time"
)

type proProjectRepoImpl struct {
	constants constants.Constants
	// sqlTool   *sqlTools.SqlTool
}

var (
	buildMySQLConnection = connections.BuildMySQLConnection
)

var (
	castToInt     = strconv.Atoi
	parseDuration = time.ParseDuration
)

type Dnas struct {
    ID   int    `json:"id"`
    Dna string `json:"dna"`
    Mutant string `json:"mutant"`
}


func GetStadisticsDnaProcessed() string {
	// // clsConnectiob := db.NewValue("/rds_db/mysql")
	connectionDb, errDto := loadConnection()
	if errDto != nil {
		panic(errDto)
	}
	defer connectionDb.Close()
	fmt.Println("connection DB", connectionDb)

	rows, err := connectionDb.Query("SELECT ID, DNA, MUTANT FROM mutants_general.DNA_VERIFICATION_MUTANTS")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Data table", rows)
	// be careful deferring Queries if you are using transactions
	defer rows.Close()
	for rows.Next() {
        var dnas Dnas
        err = rows.Scan(&dnas.ID, &dnas.Dna, &dnas.Mutant)
        if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
        }
		fmt.Println("Select ejecutado", &dnas.ID, &dnas.Dna, &dnas.Mutant)
    }
	err = rows.Err()
	return "Select ejecutado"
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

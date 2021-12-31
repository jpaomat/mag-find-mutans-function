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
    ID   int
    Dna string
    Mutant string
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
	dnas := Dnas{}
	arrayDnas := []Dnas{}
	count_mutant_dna:=0
	count_human_dna:=0
	for rows.Next() {
		var id int
		var dna, mutant string
        err = rows.Scan(&id, &dna, &mutant)
        if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
        }
		fmt.Println("Select dnasss",dna,mutant)
		dnas.ID = id
		dnas.Dna = dna
		dnas.Mutant = mutant
		if mutant == "1" {
			count_mutant_dna ++
		} else {
			count_human_dna++
		}
		arrayDnas = append(arrayDnas, dnas)
    }
	fmt.Println("Select ejecutado",arrayDnas)
	err = rows.Err()
	return fmt.Sprintf("resut mutant %d- result human %d", count_mutant_dna, count_human_dna)
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

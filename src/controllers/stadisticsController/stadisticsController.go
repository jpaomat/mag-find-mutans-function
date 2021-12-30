package stadisticscontroller

import (
	"fmt"
	"mag-stadistics-dna-processed-function/src/config/connections"
	"mag-stadistics-dna-processed-function/src/config/constants"
	"mag-stadistics-dna-processed-function/src/controllers/stadisticsController"
)

func GetStadisticsDnaProcessed() string {
	// // clsConnectiob := db.NewValue("/rds_db/mysql")
	// connectionString := constants.GetMysqlConnectionString()
	connectionDb, errDto := connectionDb := connections.GetConnectDBMysql(
		constants.GetMysqlConnectionString(),
		constants.GetMaxOpenDbConn(),
		constants.GetMasIdleDbConn(),
		constants.GetMaxDbLifetime(),
	)
	defer connectionDb.Close()

	resulSql, err := connectionDb.Query("SELECT * FROM mutants_general.DNA_VERIFICATION_MUTANTS")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Data table", resulSql)
	defer resulSql.Close()
	// be careful deferring Queries if you are using transactions
	return "Select ejecutado"
}

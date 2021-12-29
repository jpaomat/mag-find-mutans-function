package getstadisticsdnapackage

import (
	"fmt"
	db "mag-stadistics-dna-processed-function/src/packages/connectionDbPackage"
)

func GetStadisticsDnaProcessed() string {
	clsConnectiob := db.NewValue("/rds_db/mysql")
	connectionDb := clsConnectiob.GetConnectionDb()
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

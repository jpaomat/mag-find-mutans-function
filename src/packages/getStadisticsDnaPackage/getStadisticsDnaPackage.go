package getstadisticsdnapackage

import (
	"fmt"
	db "mag-stadistics-dna-processed-function/src/packages/connectionDbPackage"
)

func GetStadisticsDnaProcessed() string {
	clsConnectiob := db.NewValue("/rds_db/mysql")
	connectionDb := clsConnectiob.GetConnectionDb()

	resulSql, err := connectionDb.Query("SELECT * FROM mutants_general.DNA_VERIFICATION_MUTANTS")
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Data table", resulSql)
	// be careful deferring Queries if you are using transactions
	defer resulSql.Close()
	return "Select ejecutado"
}

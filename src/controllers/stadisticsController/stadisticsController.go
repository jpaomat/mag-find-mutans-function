package stadisticscontroller

import (
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
	rows, errData := getDnaSequencesService.GetDataDnaSequences()
	if errData != nil {
		return nil, errData
	}
	count_mutant_dna:=0
	count_human_dna:=0
	for rows.Next() {
		var mutant string
        errorScan := rows.Scan(&mutant)
		fmt.Println("...", mutant)
        if errorScan != nil {
			return nil, logger("Error to get data DB", errDefault, http.StatusInternalServerError, errorScan.Error())
        }
		if mutant == "1" {
			count_mutant_dna ++
		}
		count_human_dna++
    }
	ratio := float64(count_mutant_dna)/float64(count_human_dna)
	defer rows.Close()
	return &response.BodyStruct{
		Count_mutant_dna: count_mutant_dna,
		Count_human_dna: count_human_dna,
		Ratio: fmt.Sprintf("%.1f", ratio),
	}, nil
}


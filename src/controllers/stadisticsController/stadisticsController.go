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
	countMutantDna:=0
	countHumanDna:=0
	for rows.Next() {
		var mutant string
        errorScan := rows.Scan(&mutant)
        if errorScan != nil {
			return nil, logger("Error to get data DB", errDefault, http.StatusInternalServerError, errorScan.Error())
        }
		if mutant == "1" {
			countMutantDna ++
		}
		countHumanDna++
    }
	ratio := float64(countMutantDna)/float64(countHumanDna)
	defer rows.Close()
	return &response.BodyStruct{
		Count_mutant_dna: countMutantDna,
		Count_human_dna: countHumanDna,
		Ratio: fmt.Sprintf("%.1f", ratio),
	}, nil
}


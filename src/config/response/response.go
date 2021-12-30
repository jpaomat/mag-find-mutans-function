package response

type BodyStruct struct {
	Count_mutant_dna int     `json:"countMutants"`
	Count_human_dna  int     `json:"countHumans"`
	Ratio            float64 `json:"ratio"`
}

type Response struct {
	Message    string     `json:"message"`
	StatusCode int        `json:"statusCode"`
	Body       BodyStruct `json:"body"`
}

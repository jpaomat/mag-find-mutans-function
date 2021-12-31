package response

type BodyStruct struct {
	Count_mutant_dna int     `json:"count_mutant_dna"`
	Count_human_dna  int     `json:"count_human_dna"`
	Ratio            string `json:"ratio"`
}

type Response struct {
	Message    string     `json:"message"`
	StatusCode int        `json:"statusCode"`
	Body       BodyStruct `json:"body"`
}

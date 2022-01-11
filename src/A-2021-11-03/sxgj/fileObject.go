package sxgj

type FileObject struct {
	FileNumber     string `json:"file_number"`
	FileName       string `json:"file_name"`
	Uri            string `json:"uri"`
	AlgorithmType  int    `json:"algorithm_type"`
	Hash           string `json:"hash"`
	Summary        string `json:"summary"`
	TermOfValidity string `json:"term_of_validity"`
}

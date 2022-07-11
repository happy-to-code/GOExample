package sxgj

type FileObject struct {
	FileNumber     string `1.json:"file_number"`
	FileName       string `1.json:"file_name"`
	Uri            string `1.json:"uri"`
	AlgorithmType  int    `1.json:"algorithm_type"`
	Hash           string `1.json:"hash"`
	Summary        string `1.json:"summary"`
	TermOfValidity string `1.json:"term_of_validity"`
}

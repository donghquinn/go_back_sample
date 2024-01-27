package types

type DataResResult struct {
	result any
}

type ResponseType struct {
	ResCode int 	`json:"resCode" xml:"resCode"`
	DataRes any 	`json:"dataRes" xml:"dataRes"`
}
package types

type DataResResult struct {
	result any
}

type ResponseType  struct {
	resCode int
	dataRes DataResResult
	errMsg []string
}
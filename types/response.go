package types

type DataResResult struct {
	result any
}

type ResponseType struct {
	ResCode int 	
	DataRes any 	
	ErrMsg []string 
}
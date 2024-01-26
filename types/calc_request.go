package types

type CalculateRequest struct {
	Num1 int `form:"num1" json:"num1" xml:"num1" binding:"required"`
	Num2 int `form:"num2" json:"num2" xml:"num2" binding:"required"`
}
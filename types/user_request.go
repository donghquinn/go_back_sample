package types

type UserLoginRequest struct {
	Email string `json:"email" xml:"email" binding:"required"`
	Password string `json:"password" xml:"password" binding:"required"`
}
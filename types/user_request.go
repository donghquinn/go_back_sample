package types

type UserLoginRequest struct {
	Email string `json:"email" xml:"email" binding:"required,email"`
	Password string `json:"password" xml:"password" binding:"required,min=0"`
}

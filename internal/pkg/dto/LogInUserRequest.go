package dto

type LogInUserRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,e164"`
	Password    string `json:"password" binding:"required,min=6,max=15"`
}

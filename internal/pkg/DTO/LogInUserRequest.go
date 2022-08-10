package DTO

type LogInUserRequest struct {
	PhoneNumber string `json:"phoneNumber" validate:"required,e164"`
	Password    string `json:"password" validate:"required,min=6,max=15"`
}

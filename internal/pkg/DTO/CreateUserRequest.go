package DTO

type CreateUserRequest struct {
	Name        string `json:"name" validate:"required,min=4,max=15"`
	PhoneNumber string `json:"phoneNumber" validate:"required,e164"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=6,max=15"`
}

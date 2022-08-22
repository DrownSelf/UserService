package DTO

type CreateUserRequest struct {
	Name        string `json:"name" binding:"required,min=4,max=15"`
	PhoneNumber string `json:"phoneNumber" binding:"required,e164"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6,max=15"`
}

package DTO

type ChangeUserRequest struct {
	PhoneNumber    string `json:"phoneNumber" binding:"required,e164"`
	NewPhoneNumber string `json:"newPhoneNumber" binding:"required,e164"`
	NewEmail       string `json:"newEmail" binding:"required,email"`
	NewName        string `json:"newName" binding:"required,min=6,max=15"`
}

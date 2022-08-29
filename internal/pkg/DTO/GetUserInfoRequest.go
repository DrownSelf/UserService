package DTO

type GetUserInfoRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,e164"`
}

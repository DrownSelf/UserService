package dto

type GetUserInfoRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,e164"`
}

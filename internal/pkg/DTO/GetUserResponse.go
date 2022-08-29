package DTO

type GetUserResponse struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Rating      int    `json:"rating"`
}

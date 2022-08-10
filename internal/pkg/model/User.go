package model

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

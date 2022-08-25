package model

import "time"

type User struct {
	Id          int        `json:"id" `
	Name        string     `json:"name"`
	PhoneNumber string     `json:"phoneNumber"`
	Email       string     `json:"email" `
	Password    string     `json:"password"`
	Rating      *int       `json:"rating"`
	CreatedAt   *time.Time `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	IsDeleted   bool       `json:"isDeleted"`
}

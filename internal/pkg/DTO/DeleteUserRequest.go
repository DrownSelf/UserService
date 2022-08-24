package DTO

type DeleteUserRequest struct {
	Id int `json:"id" validate:"required"`
}

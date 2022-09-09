package dto

type DeleteUserRequest struct {
	Id int `json:"id" validate:"required"`
}

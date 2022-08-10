package DTO

type GetUserResponse struct {
	Id    int    `json:"id"`
	Token string `json:"jwt"`
}

package handlers

type RateRideRequest struct {
	Id     string `json:"id"`
	Rating int32  `json:"rating"`
}

type RateUserFromOrderRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Rating      int32  `json:"rating"`
}

type MakeOrderRequest struct {
	From     string `json:"from"`
	To       string `json:"to"`
	TaxiType string `json:"taxiType"`
}

type LogInUserRequest struct {
	PhoneNumber string `json:"phoneNumber" binding:"required,e164"`
	Password    string `json:"password" binding:"required,min=6,max=15"`
}

type ChangeUserRequest struct {
	NewPhoneNumber string `json:"newPhoneNumber" binding:"required,e164"`
	NewEmail       string `json:"newEmail" binding:"required,email"`
	NewName        string `json:"newName" binding:"required,min=6,max=15"`
}

type CreateUserRequest struct {
	Name        string `json:"name" binding:"required,min=4,max=15"`
	PhoneNumber string `json:"phoneNumber" binding:"required,e164"`
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6,max=15"`
}

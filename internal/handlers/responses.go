package handlers

type UserRideResponse struct {
	Id                string `json:"id"`
	From              string `json:"from"`
	To                string `json:"to"`
	DriverName        string `json:"driverName"`
	DriverPhoneNumber string `json:"driverPhoneNumber"`
	TaxiType          string `json:"taxiType"`
}

type GetUserResponse struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber string  `json:"phoneNumber"`
	Rating      float64 `json:"rating"`
}

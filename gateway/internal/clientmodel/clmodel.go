package clientmodel

type ClientRegisterRequest struct {
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Phone		string	`json:"phone"`
	City        string  `json:"city"`
	Street      string  `json:"street"`
	Home_number string  `json:"home_number"`
	Bank_card   string  `json:"bank_card"`
	Balance     float32 `json:"balance"`
}

type ClientLogin struct{
	Id string `json:"id"`
	Email string `json:"email"`
}

type LoginResponse struct{
	Status string `json:"status"`
	Token string `json:"token"`
}

package domain

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Register struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_Name"`
	LastName    string `json:"last_name"`
}

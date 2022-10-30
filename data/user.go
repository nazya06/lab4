package data

type User struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Age         int32  `json:"age"`
	Surname     string `json:"surname,omitempty"`
}

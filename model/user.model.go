package model

type User struct {
	Id          string `json:"id"`
	CreatedBy   string `json:"createdBy"`
	CreatedDate string `json:"createdDate"`
	CreatedFrom string `json:"createdFrom"`
	Email       string `json:"email"`
	FullName    string `json:"fullName"`
}

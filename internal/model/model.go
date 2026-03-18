package model

type User struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password int    `json:"password"`
}

package model

type User struct {
	User_id  int    `json:"id"`
	Login    string `json:"login"`
	Password int    `json:"password"`
}

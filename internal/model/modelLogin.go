package model

type LoginResponse struct {
	Role string `json:"role"`
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

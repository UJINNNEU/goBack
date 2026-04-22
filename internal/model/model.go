package model

import "encoding/json"

type User struct {
	ID       int    `json:"user_id"`
	Login    string `json:"loginZ"`
	Password int    `json:"passwordZ"`
}

type TestFull struct {
	Test
	Questions json.RawMessage `json:"Questions"`
}

type Test struct {
	Id          int    `json:"Id"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}
package model

import "encoding/json"

type User struct {
	ID       int    `json:"user_id"`
	Login    string `json:"loginZ"`
	Password int    `json:"passwordZ"`
}

type TestFull struct {
	Test
	Questions []Question `json:"Question"`
	a         json.RawMessage `json: "Test"`
}

type Test struct {
	Test_id          int    `json:"Test_id"`
	Test_title       string `json:"Test_title"`
	Test_description string `json: "Test_description"`
}

type Question struct {
	Question_id    int    `json: "Question_id"`
	Question_title string `json: "Question_title"`
}

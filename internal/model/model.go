package model

type Test struct {
	id_test          int
	version_test     int
	title_test       string
	description_test string
	duration         int
	questions        []Question
}

type Question struct {
	Question_id    int
	Index          int
	Question_title string
}

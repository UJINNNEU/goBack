package test

import "database/sql"

type TestStorage struct {
	db *sql.DB
}

func New(db *sql.DB) *TestStorage {
	return &TestStorage{
		db: db}
}
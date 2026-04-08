package serviceT

import (
	"backend/internal/repository/storage/test"
)

type TestService struct {
	test *test.TestStorage
}

func New(testStorage *test.TestStorage) *TestService {
	return &TestService{
		test: testStorage,
	}
}
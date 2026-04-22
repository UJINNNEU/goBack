package serviceT

import (
	"backend/internal/repository/storage/test"
)

// TODO локальные интерфейсфы с маленькой буквы (инкапсулируем)
type TestService struct {
	test *test.TestStorage
}

// TODO принимать интерфейс (проверь в коде везде)
func New(testStorage *test.TestStorage) *TestService {
	return &TestService{
		test: testStorage,
	}
}

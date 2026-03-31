package service

import (
	"backend/internal/model"
	"backend/internal/repository"
	"context"
)

type TestService interface {
	GetTestById(ctx context.Context, id int) (model.TestFull, error)
	GetAvailableTests(ctx context.Context, user_id int) ([]model.TestFull, error)
	//getAllTest(ctx context.Context) []model.TestFull

}

type testService struct {
	testRepo repository.TestRepository
}

func NewTestService(testRepo repository.TestRepository) TestService {
	return &testService{
		testRepo: testRepo,
	}
}
func (t *testService) GetTestById(ctx context.Context, id int) (model.TestFull, error) {
	return t.testRepo.GetTestById(ctx,id)
}

func(t *testService) GetAvailableTests(ctx context.Context, user_id int) ([]model.TestFull, error)  {
	return t.testRepo.GetAvailableTests(ctx, user_id)
}
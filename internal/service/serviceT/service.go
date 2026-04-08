package serviceT

import (
	"backend/internal/model"
	"context"
)

func (t *TestService) GetTestById(ctx context.Context, id int) (model.TestFull, error) {
	return t.test.GetTestById(ctx, id)
}

func (t *TestService) GetAvailableTests(ctx context.Context, user_id int) ([]model.TestFull, error) {
	tests, err := t.test.GetAvailableTestsId(ctx, user_id)
	var avaiilableTests []model.TestFull

	for _, value := range tests {
		test, err := t.GetTestById(ctx, value)
		if err != nil {
			return nil, err
		}

		if test.Id != 0 && test.Description != "" && test.Title != "" {
			avaiilableTests = append(avaiilableTests, test)
		}
	}
	return avaiilableTests, err
}

package test

import (
	"backend/internal/model"
	"context"
	"fmt"
)

func (s *TestStorage) GetTestById(ctx context.Context, id int) (model.TestFull, error) {
	var testfull model.TestFull

	query := `SELECT 
    		t.id,
    		t.title_test,
    		t.description_test,
    	  	t.questions_in_test
		FROM tests t
		WHERE t.id = $1;`

	//TODO сделать везде с контекстом
	rows, err := s.db.QueryContext(ctx, query, id)
	if err != nil {

		return testfull, err
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(
			&testfull.Test.Id,
			&testfull.Title,
			&testfull.Description,
			&testfull.Questions,
		); err != nil {
			return testfull, err

		}
	}
	fmt.Println(testfull)
	return testfull, nil
}

func (s *TestStorage) GetAvailableTestsId(ctx context.Context, user_id int) ([]int, error) {

	// доступные тесты
	availableTestsId := []int{}
	rows, err := s.db.Query(`SELECT t.id
		FROM tests t
		JOIN test_rules r ON r.id_test = t.id
		LEFT JOIN answers a 
    		ON a.test_id = t.id
    		AND a.user_id = $1
		GROUP BY t.id, r.retake_interval, r.retake_type
		HAVING 
    		MAX(a.date) IS NULL
    		OR MAX(a.date) <= CURRENT_DATE - 
        (r.retake_interval || ' ' || r.retake_type)::INTERVAL;`,
		user_id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var test_id int

		if err := rows.Scan(&test_id); err != nil {
			return nil, err
		}
		availableTestsId = append(availableTestsId, test_id)
	}
	return availableTestsId, nil
}

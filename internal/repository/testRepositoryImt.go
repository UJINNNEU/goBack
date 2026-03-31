package repository

import (
	"backend/internal/model"
	"context"
	"database/sql"
	"fmt"
)

type testPostgres struct {
	db *sql.DB
}

func NewTestRepository(db *sql.DB) TestRepository {
	return &testPostgres{db: db}
}

func (r *testPostgres) GetTestById(ctx context.Context, id int) (model.TestFull, error) {

	rows, err := r.db.Query(
		`SELECT 
    		t.id_test,
    		t.title_test,
    		t.description_test,
    		q.id_question,
    		q.title_question
		FROM tests t
		LEFT JOIN questions q ON t.id_test = q.test_id
		WHERE t.id_test = $1;`, id)

	//defer rows.Close()
	var sliceQuestion []model.Question
	var testfull model.TestFull
	var testSet bool = false

	if err != nil {
		return testfull, err
	}

	for rows.Next() {

		var test model.Test
		var question model.Question

		err := rows.Scan(
			&test.Test_id, &test.Test_title, &test.Test_description,
			&question.Question_id, &question.Question_title,
		)

		if err != nil {
			return testfull, err
		}
		if !testSet {
			testSet = !testSet
			testfull.Test = test
		}

		sliceQuestion = append(sliceQuestion, question)
	}

	if err := rows.Err(); err != nil {
		return testfull, err
	}

	fmt.Println(testfull)
	testfull.Questions = sliceQuestion

	return testfull, nil
}

func (r *testPostgres) GetAvailableTests(ctx context.Context, user_id int) ([]model.TestFull, error) {

	// доступные тесты
	availableTests := []model.TestFull{}
	availableTestsId := []int{}
	rows, err := r.db.Query(`SELECT t.id_test
		FROM tests t
		JOIN test_rules r ON r.id_test = t.id_test
		LEFT JOIN answers a 
    		ON a.id_test = t.id_test 
    		AND a.user_id = $1
		GROUP BY t.id_test, r.retake_interval, r.retake_type
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

	for _, test_id := range availableTestsId {
		testfull, err := r.GetTestById(ctx, test_id)
		if err != nil {
			return nil, err
		}
		availableTests = append(availableTests, testfull)
	}

	return availableTests, nil
}

/*func (r *testPostgres) getAllTest(ctx context.Context) []model.TestFull {
	return []model.TestFull{}
}*/

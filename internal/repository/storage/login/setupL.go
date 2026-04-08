package login

import(
	"database/sql"
)
type LoginRepository struct {
	db *sql.DB
}
func New( db *sql.DB) (*LoginRepository) {
	return &LoginRepository{
		db: db,
	}
}
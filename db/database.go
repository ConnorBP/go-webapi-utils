package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func db_get_user_info(db *sql.DB, user_id int64) (string, string, error) {
	r := db.QueryRow("SELECT username,uuid4 FROM users where id=?;", user_id)
	result := make([]string, 0)
	err := r.Scan(&result)
	if err == sql.ErrNoRows {
		// new_uuid := uuid.New()
		// _, err = db.Exec("UPDATE users set uuid4 = '?', ", new_uuid)
	} else if err != nil {
		return "", "", err
	}

	// success
	return result[0], result[1], nil
}

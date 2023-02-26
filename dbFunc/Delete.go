package dbFunc

import "database/sql"

func Delete(db *sql.DB, id int) (bool, error) {
	res, err := db.Exec("DELETE FROM todos WHERE id = ($1)", id)
	if err != nil {
		return false, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	if n > 0 {
		return true, nil
	}
	return false, nil
}

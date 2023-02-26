package dbFunc

import (
	"database/sql"
	"todolistapi/structs"
)

func GetByID(db *sql.DB, id int) (*structs.Todo, bool, error) {
	rows, err := db.Query("SELECT * FROM todos WHERE id = ($1)", id)
	if err != nil {
		return nil, false, err
	}

	if rows.Next() {
		res := structs.Todo{}
		rows.Scan(&res.ID, &res.Name, &res.Description, &res.Start, &res.Finish, &res.IsDone)
		return &res, true, nil
	} else {
		return nil, false, nil
	}

}

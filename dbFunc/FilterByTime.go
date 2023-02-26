package dbFunc

import (
	"database/sql"
	"todolistapi/structs"
)

func FilterByTime(db *sql.DB, after string, before string) ([]structs.Todo, error) {
	var res []structs.Todo
	var rows *sql.Rows
	var err error

	if after != "" && before != "" {
		rows, err = db.Query("SELECT * FROM todos WHERE start <= ($1) AND finish IS NOT NULL AND finish >= ($2)", before, after)
	} else if after != "" {
		rows, err = db.Query("SELECT * FROM todos WHERE finish >= ($1)", after)
	} else if before != "" {
		rows, err = db.Query("SELECT * FROM todos WHERE start <= ($1)", before)
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var todo structs.Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.Start, &todo.Finish, &todo.IsDone)
		res = append(res, todo)
	}

	return res, nil
}

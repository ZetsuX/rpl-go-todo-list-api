package dbFunc

import (
	"database/sql"
	"todolistapi/structs"
)

func GetAll(db *sql.DB) ([]structs.Todo, error) {
	var res []structs.Todo

	rows, err := db.Query("SELECT * FROM todos ORDER BY id ASC")
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

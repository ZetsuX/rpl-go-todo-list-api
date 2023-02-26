package dbFunc

import (
	"database/sql"
	"todolistapi/structs"
)

func Edit(db *sql.DB, id int, todo structs.Todo) (*structs.Todo, bool, error) {
	var rows *sql.Rows
	var err error

	if todo.Finish == "" {
		rows, err = db.Query("UPDATE todos SET name = ($1), description = ($2), start = ($3), finish = NULL, isDone = ($4) WHERE id = ($5) RETURNING id, name, description, start, finish, isDone", todo.Name, todo.Description, todo.Start, todo.IsDone, id)
	} else {
		rows, err = db.Query("UPDATE todos SET name = ($1), description = ($2), start = ($3), finish = ($4), isDone = ($5) WHERE id = ($6) RETURNING id, name, description, start, finish, isDone", todo.Name, todo.Description, todo.Start, todo.Finish, todo.IsDone, id)
	}

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

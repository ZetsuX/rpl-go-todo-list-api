package dbFunc

import (
	"database/sql"
	"todolistapi/structs"
)

func Insert(db *sql.DB, todo structs.Todo) (*structs.Todo, error) {
	var rows *sql.Rows
	var err error

	if todo.Finish == "" {
		rows, err = db.Query("INSERT INTO todos (name, description, start, isDone) VALUES ($1, $2, $3, $4) RETURNING id, name, description, start, finish, isDone", todo.Name, todo.Description, todo.Start, false)
	} else {
		rows, err = db.Query("INSERT INTO todos (name, description, start, finish, isDone) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, description, start, finish, isDone", todo.Name, todo.Description, todo.Start, todo.Finish, false)
	}

	if err != nil {
		return nil, err
	}

	rows.Next()
	res := structs.Todo{}

	rows.Scan(&res.ID, &res.Name, &res.Description, &res.Start, &res.Finish, &res.IsDone)
	return &res, nil
}

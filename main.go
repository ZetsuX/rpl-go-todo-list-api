package main

import (
	"fmt"
	"net/http"

	"todolistapi/dbFunc"
	"todolistapi/handleFunc"

	"github.com/darahayes/go-boom"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Postman Docs : https://documenter.getpostman.com/view/25087235/2s93CNLsa9

	db, err := dbFunc.Connect()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			boom.NotFound(w, "Not found")
			return
		}

		switch r.Method {
		case http.MethodGet:
			err = handleFunc.GetTodos(w, r, db)
		case http.MethodPost:
			err = handleFunc.AddNewTodo(w, r, db)
		case http.MethodPut:
			err = handleFunc.EditTodo(w, r, db)
		case http.MethodDelete:
			err = handleFunc.DeleteTodo(w, r, db)
		default:
			boom.MethodNotAllowed(w, "Method not allowed")
		}

		if err != nil {
			fmt.Println(err)
		}
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/search" {
			boom.NotFound(w, "Not found")
			return
		}

		if r.Method == http.MethodGet {
			err = handleFunc.SearchTodos(w, r, db)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			boom.MethodNotAllowed(w, "Method not allowed")
		}
	})

	http.HandleFunc("/filter", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/filter" {
			boom.NotFound(w, "Not found")
			return
		}

		if r.Method == http.MethodGet {
			err = handleFunc.FilterTodos(w, r, db)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			boom.MethodNotAllowed(w, "Method not allowed")
		}
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

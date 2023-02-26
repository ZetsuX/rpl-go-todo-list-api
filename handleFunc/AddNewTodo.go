package handleFunc

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"todolistapi/dbFunc"
	"todolistapi/structs"

	"github.com/darahayes/go-boom"
)

func AddNewTodo(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	err := r.ParseForm()
	if err != nil {
		boom.Internal(w, err.Error())
		return err
	}

	r.ParseMultipartForm(0)
	name := r.FormValue("name")
	description := r.FormValue("description")
	start := r.FormValue("start")
	finish := r.FormValue("finish")

	newTodo := structs.Todo{
		Name:        name,
		Description: description,
		Start:       start,
		Finish:      finish,
		IsDone:      false,
	}

	if name == "" {
		boom.BadRequest(w, "'name' field is not filled")
		return errors.New("'name' field is not filled")
	}

	if start == "" {
		boom.BadRequest(w, "'start' field is not filled")
		return errors.New("'start' field is not filled")
	} else if start != "" && finish != "" {
		startTime, err := time.Parse(time.RFC3339, start)
		if err != nil {
			boom.BadRequest(w, err.Error())
		}

		finishTime, err := time.Parse(time.RFC3339, finish)
		if err != nil {
			boom.BadRequest(w, err.Error())
		}

		if finishTime.Before(startTime) {
			boom.BadRequest(w, "'start' is further than 'finish'")
			return errors.New("'start' is further than 'finish'")
		}
	}

	res, err := dbFunc.Insert(db, newTodo)
	if err != nil {
		boom.BadRequest(w, err.Error())
		return err
	}

	jsonMap := map[string]interface{}{
		"success": true,
		"data":    res,
	}

	b, err := json.Marshal(jsonMap)
	if err != nil {
		boom.Internal(w, err.Error())
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, string(b))
	return nil
}

package handleFunc

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todolistapi/dbFunc"
	"todolistapi/structs"

	"github.com/darahayes/go-boom"
)

func EditTodo(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	keys, success := r.URL.Query()["id"]

	if !success || len(keys[0]) < 1 {
		boom.BadRequest(w, "Url param 'id' is missing")
		return errors.New("url param 'id' is missing")
	}

	id, err := strconv.Atoi(keys[0])
	if err != nil {
		boom.BadRequest(w, err.Error())
		return err
	}

	err = r.ParseForm()
	if err != nil {
		boom.Internal(w, err.Error())
		return err
	}

	r.ParseMultipartForm(0)
	name := r.FormValue("name")
	description := r.FormValue("description")
	start := r.FormValue("start")
	finish := r.FormValue("finish")
	tmp := r.FormValue("isDone")

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

	isDone := false
	if tmp == "t" || tmp == "true" {
		isDone = true
	}

	newTodo := structs.Todo{
		Name:        name,
		Description: description,
		Start:       start,
		Finish:      finish,
		IsDone:      isDone,
	}

	res, check, err := dbFunc.Edit(db, id, newTodo)
	if err != nil {
		boom.BadRequest(w, err.Error())
		return err
	}

	var jsonMap map[string]interface{}
	if check {
		jsonMap = map[string]interface{}{
			"success": true,
			"data":    res,
		}
	} else {
		jsonMap = map[string]interface{}{
			"success": false,
			"data":    nil,
			"message": "No data found with the inserted id",
		}
	}

	b, err := json.Marshal(jsonMap)
	if err != nil {
		boom.Internal(w, err.Error())
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(b))
	return nil
}

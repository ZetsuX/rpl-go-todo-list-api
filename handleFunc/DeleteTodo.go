package handleFunc

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"todolistapi/dbFunc"

	"github.com/darahayes/go-boom"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
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

	var jsonMap map[string]interface{}
	check, err := dbFunc.Delete(db, id)
	if err != nil {
		boom.Internal(w, err.Error())
		return err
	}

	if check {
		jsonMap = map[string]interface{}{
			"success": true,
		}
	} else {
		jsonMap = map[string]interface{}{
			"success": false,
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

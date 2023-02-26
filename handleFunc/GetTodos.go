package handleFunc

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"todolistapi/dbFunc"

	"github.com/darahayes/go-boom"
)

func GetTodos(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	keys, success := r.URL.Query()["id"]
	var err error
	var jsonMap map[string]interface{}

	if !success || len(keys[0]) < 1 {
		res, err := dbFunc.GetAll(db)
		if err != nil {
			boom.Internal(w, err.Error())
			return err
		}

		jsonMap = map[string]interface{}{
			"data":    res,
			"success": true,
		}

	} else {
		id, err := strconv.Atoi(keys[0])
		if err != nil {
			boom.BadRequest(w, err.Error())
			return err
		}

		res, check, err := dbFunc.GetByID(db, id)
		if err != nil {
			boom.Internal(w, err.Error())
			return err
		}

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

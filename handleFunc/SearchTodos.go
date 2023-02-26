package handleFunc

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"todolistapi/dbFunc"

	"github.com/darahayes/go-boom"
)

func SearchTodos(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	keys, success := r.URL.Query()["q"]

	if !success || len(keys[0]) < 1 {
		boom.BadRequest(w, "Url param 'q' is missing")
		return errors.New("url param 'q' is missing")
	}

	res, err := dbFunc.SearchByQuery(db, keys[0])
	if err != nil {
		boom.Internal(w, err.Error())
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
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(b))
	return nil
}

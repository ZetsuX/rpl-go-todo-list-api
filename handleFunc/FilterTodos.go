package handleFunc

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"todolistapi/dbFunc"

	"github.com/darahayes/go-boom"
)

func FilterTodos(w http.ResponseWriter, r *http.Request, db *sql.DB) error {
	var after, before string

	keys, success := r.URL.Query()["after"]
	if success && len(keys[0]) > 1 {
		after = keys[0]
	}

	keys, success = r.URL.Query()["before"]
	if success && len(keys[0]) > 1 {
		before = keys[0]
	}

	if before == "" && after == "" {
		boom.BadRequest(w, "Url param 'after' and 'before' are missing")
		return errors.New("url param 'after' and 'before' are missing")
	} else if before != "" && after != "" {
		afterTime, err := time.Parse(time.RFC3339, after)
		if err != nil {
			boom.BadRequest(w, err.Error())
		}

		beforeTime, err := time.Parse(time.RFC3339, before)
		if err != nil {
			boom.BadRequest(w, err.Error())
		}

		if beforeTime.Before(afterTime) {
			boom.BadRequest(w, "Url param 'after' is further than 'before'")
			return errors.New("url param 'after' is further than 'before'")
		}
	}

	res, err := dbFunc.FilterByTime(db, after, before)
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

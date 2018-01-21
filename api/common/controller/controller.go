package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Dlacreme/httpd/back/wesult"
)

// Final will provide the final output
func Final(w http.ResponseWriter, res wesult.Result) {
	if res.Error != nil {
		w.WriteHeader(res.Error.Code)
		w.Write([]byte(res.Error.Message))
		return
	}
	json.NewEncoder(w).Encode(res.Output)
}

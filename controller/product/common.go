package product

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func writeResponse(res []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal(body, x); err != nil {
			return
		}
	}
}

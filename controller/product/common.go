package product

import (
	"encoding/json"
	"github.com/deniarianto1606/go-store/product/ports/findbycode"
	"github.com/deniarianto1606/go-store/product/ports/save"
	"io/ioutil"
	"net/http"
)

// Handler product handler
type Handler interface {
	FindByCode(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	findByCode findbycode.UseCase
	save       save.UseCase
}

func NewHandler(findByCode findbycode.UseCase, save save.UseCase) Handler {
	return &handler{findByCode: findByCode, save: save}
}

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

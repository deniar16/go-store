package product

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

func (h *handler) FindByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	product, err := h.findByCode.FindByCode(code)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	res, _ := json.Marshal(product)
	writeResponse(res, w)
}

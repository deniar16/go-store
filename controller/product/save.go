package product

import (
	"encoding/json"
	"github.com/deniarianto1606/go-store/product/domain"
	"net/http"
)

func (h *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := &domain.Product{}
	ParseBody(r, product)
	b := h.save.Save(product)
	res, _ := json.Marshal(b)
	writeResponse(res, w)
}

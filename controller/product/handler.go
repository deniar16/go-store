package product

import (
	"github.com/deniarianto1606/go-store/product/ports/findbycode"
	"github.com/deniarianto1606/go-store/product/ports/save"
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

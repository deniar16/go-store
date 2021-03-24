package product

import (
	"encoding/json"
	"github.com/deniarianto1606/go-store/product/domain"
	"github.com/deniarianto1606/go-store/product/ports"
	"github.com/deniarianto1606/go-store/product/ports/findbycode"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

// Handler product handler
type Handler interface {
	FindByCode(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	productService ports.ProductService
	findByCode     findbycode.UseCase
}

func NewHandler(service ports.ProductService, findByCode findbycode.UseCase) Handler {
	return &handler{productService: service, findByCode: findByCode}
}

func (p *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := &domain.Product{}
	ParseBody(r, product)
	b := p.productService.Save(product)
	res, _ := json.Marshal(b)
	writeResponse(res, w)
}

func (p *handler) FindByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	product, err := p.findByCode.FindByCode(code)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	res, _ := json.Marshal(product)
	writeResponse(res, w)
}

func writeResponse(res []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

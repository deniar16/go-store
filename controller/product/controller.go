package product

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	prod "github.com/deniarianto1606/go-store/product"
)

type ProductHandler interface {
	//Get(w http.ResponseWriter, r *http.Request)
	CreateProduct(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	productService prod.ProductService
}

func NewHandler(service prod.ProductService) ProductHandler {
	return &handler{productService: service}
}

func (p *handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := &prod.Product{}
	ParseBody(r, product)
	b:= p.productService.Save(product)
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
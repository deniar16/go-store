package order

import (
	"encoding/json"
	"github.com/deniarianto1606/go-store/order/domain"
	"github.com/deniarianto1606/go-store/order/ports"
	"github.com/go-chi/chi"
	"io/ioutil"
	"net/http"
)

type OrderHandler interface {
	FindByCode(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	orderService ports.OrderService
}

func NewHandler(service ports.OrderService) OrderHandler {
	return &handler{orderService: service}
}

func (p *handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	product := &domain.Order{}
	ParseBody(r, product)
	b:= p.orderService.Save(product)
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (p *handler) FindByCode(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	product, err := p.orderService.FindByCode(code)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
	res,_ := json.Marshal(product)
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

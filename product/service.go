package product

type ProductService interface {
	FindByCode(code string) (*Product, error)
	Save(product *Product) error
}

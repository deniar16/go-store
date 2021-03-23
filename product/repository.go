package product

type ProductRepository interface {
	FindByCode(code string) (*Product, error)
	Save(product *Product) error
}
package pantrygo

type Basket[V any] interface {
	Create(payload map[string]*V) error
	Update(payload map[string]*V) error
	Get() (map[string]*V, error)
	RemoveBasket() error
	DeleteByID(id string) error
	GetByID(id string) (*V, error)
}

func ToList[V any](basketData map[string]*V) []*V {

	var contents []*V

	for _, data := range basketData {
		contents = append(contents, data)
	}

	return contents
}

func NewBasketClient[V any](key, basketName string) Basket[V] {
	return &basketImpl[V]{
		key:        key,
		basketName: basketName,
	}
}

package pantrygo

import (
	"fmt"
	"os"
	"testing"
)

const pantryID = "edab71ed-482a-4cc6-a57b-a86acc86742f"

type Product struct {
	ID    int
	Name  string
	Price int
}

func (p *Product) String() string {
	return fmt.Sprintf("ID: %d, Name %s, price %d", p.ID, p.Name, p.Price)
}

func TestCreateBasket(t *testing.T) {

	b := NewBasketClient[Product](pantryID, "products")

	contents := make(map[string]*Product)
	err := b.Create(contents)

	if err != nil {
		t.Error(err)
	}

}

func TestUpdateBasket(t *testing.T) {
	b := NewBasketClient[Product](pantryID, "products")

	prod := &Product{
		ID:    1,
		Name:  "my Product",
		Price: 30,
	}

	req := make(map[string]*Product)
	req["1"] = prod
	req["2"] = &Product{
		ID:    2,
		Name:  "New Product",
		Price: 15,
	}

	err := b.Update(req)

	if err != nil {
		t.Error(err)
	}

}

func TestGetBasket(t *testing.T) {

	b := NewBasketClient[Product](pantryID, "products")

	prods, err := b.Get()

	if err != nil {
		t.Error(err)
	}

	for key, val := range prods {
		fmt.Printf("%s -> %s\n", key, val)
	}

}

func TestDelete(t *testing.T) {
	b := NewBasketClient[Product](pantryID, "products")
	err := b.RemoveBasket()

	if err != nil {
		t.Error(err)
	}

}

func TestDeleteByID(t *testing.T) {
	b := NewBasketClient[Product](pantryID, "products")
	err := b.DeleteByID("2")

	if err != nil {
		t.Error(err)
	}
}

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

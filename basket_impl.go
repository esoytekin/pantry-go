package pantrygo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const basePath = "https://getpantry.cloud"
const apiVersion = "1"
const timeout = time.Second * 5

type basketImpl[V any] struct {
	key        string
	basketName string
}

func (b *basketImpl[V]) getBasketURL() string {

	return fmt.Sprintf("%s/apiv%s/pantry/%s/basket/%s", basePath, apiVersion, b.key, b.basketName)
}

func (b *basketImpl[V]) Create(payload map[string]*V) error {
	basketURL := b.getBasketURL()
	resp := basketUtils.request(http.MethodPost, basketURL, payload)

	if resp.error != nil {
		return resp.error
	}

	respStr := string(resp.body)

	if strings.Contains(respStr, createBasketMsg) {
		return nil
	}

	return errx.ErrUnknown
}

func (b *basketImpl[V]) Update(payload map[string]*V) error {
	basketURL := b.getBasketURL()
	resp := basketUtils.request(http.MethodPut, basketURL, payload)

	if resp.error != nil {
		if errors.Is(resp.error, errx.ErrPantry) {
			time.Sleep(timeout)
			return b.Update(payload)

		}
		return resp.error
	}

	return nil
}

func (b *basketImpl[V]) Get() (map[string]*V, error) {
	basketURL := b.getBasketURL()
	resp := basketUtils.request(http.MethodGet, basketURL, nil)

	if resp.error != nil {
		if errors.Is(resp.error, errx.ErrPantry) {
			time.Sleep(timeout)
			return b.Get()
		}
		return nil, resp.error
	}

	var data map[string]*V
	err := json.Unmarshal(resp.body, &data)

	return data, err
}

func (b *basketImpl[V]) RemoveBasket() error {
	basketURL := b.getBasketURL()

	resp := basketUtils.request(http.MethodDelete, basketURL, nil)

	if resp.error != nil {
		if errors.Is(resp.error, errx.ErrPantry) {
			time.Sleep(timeout)
			return b.RemoveBasket()
		}
		return resp.error
	}

	return nil
}

func (b *basketImpl[V]) DeleteByID(id string) error {

	data, err := b.Get()

	if err != nil {
		return err
	}

	item := data[id]

	if item == nil {
		return errx.ErrResourceNotFound
	}

	delete(data, id)

	return b.Create(data)

}

func (b *basketImpl[V]) GetByID(id string) (*V, error) {
	data, err := b.Get()

	if err != nil {
		return nil, err
	}

	item := data[id]

	if item == nil {
		return nil, errx.ErrResourceNotFound
	}

	return item, nil
}

package pantrygo

import (
	"errors"
	"io/ioutil"
	"net/http"
)

type responseWrapper struct {
	body       []byte
	response   *http.Response
	statusCode int
	error
}

func (r *responseWrapper) withError(err error) *responseWrapper {
	r.error = err
	return r
}

func (r *responseWrapper) withResponse(resp *http.Response) *responseWrapper {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.error = err
		return r
	}

	if resp.StatusCode != 200 {
		respBody := string(body)

		if basketUtils.isPantryError(respBody) {
			r.error = errx.ErrPantry
		} else {
			r.error = errors.New(respBody)
		}
		return r
	}

	r.body = body
	r.response = resp
	r.statusCode = resp.StatusCode
	return r
}

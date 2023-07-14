package pantrygo

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

const pantryMsg = "Pantry is a free service"
const createBasketMsg = "Your Pantry was updated with basket"

var basketUtils = &utils{}

type utils struct{}

func (utils) isPantryError(msg string) bool {

	return strings.Contains(msg, pantryMsg)

}

func (u utils) request(method, url string, payload interface{}) *responseWrapper {
	r := &responseWrapper{}
	var payloadReader *bytes.Reader = bytes.NewReader(make([]byte, 0))

	if payload != nil {
		jsonBody, _ := json.Marshal(payload)
		payloadReader = bytes.NewReader(jsonBody)
	}

	requestURL := url
	hR, err := http.NewRequest(method, requestURL, payloadReader)
	hR.Header.Set("Content-Type", "application/json")

	if err != nil {
		r.withError(errx.ErrUnknown)
	}

	resp, err := http.DefaultClient.Do(hR)

	if err != nil {
		r.withError(errx.ErrUnknown)
	}

	defer resp.Body.Close()

	return r.withResponse(resp)

}

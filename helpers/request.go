package helpers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

// JSONRequest :
func JSONRequest(_ context.Context, url, method string, body interface{}, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}

	var bBody []byte
	var err error
	if body != nil {
		bBody, err = json.Marshal(&body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bBody))
	req.Header.Add("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Add(k, v)
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return response, nil
}

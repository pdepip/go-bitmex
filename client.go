/*
   client.go
       Wrapper for the Bitmex Exchange Api

   Author:
       Pat DePippo <pat@dcrypt.io>
*/
package bitmex

import (
    "fmt"
    "net/http"
    "encoding/json"
    "bytes"
)

type Client struct {
	key        string
	secret     string
	httpClient *http.Client
}

func NewClient(key, secret string) *Client {
	return &Client{
		key:        key,
		secret:     secret,
		httpClient: &http.Client{},
	}
}

func (c *Client) do(method, resource string, params, result interface{}) (res *http.Response, err error) {

	// Build url
	fullUrl := fmt.Sprintf("%s/%s", ApiUrl, resource)

	// Parse params
	var data []byte
	body := bytes.NewReader(make([]byte, 0))

	if params != nil {
		data, err = json.Marshal(params)
		if err != nil {
			return res, err
		}

		body = bytes.NewReader(data)
	}

	// Build Request
	req, err := http.NewRequest(method, fullUrl, body)
	if err != nil {
		return
	}

	// Handle Auth

	// Make request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
	} else {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(result)
	}
	return
}

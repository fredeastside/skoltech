package devices

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//Client structure
type Client struct {
	url  string
	port string
}

//NewClient client constructor
func NewClient(url, port string) *Client {
	return &Client{url, port}
}

//Send data to partner url
func (c *Client) Send(d *Device) (*http.Response, error) {
	err := d.Validate()
	if err != nil {
		return nil, err
	}
	err = d.Transform()
	if err != nil {
		return nil, err
	}

	jsonString, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	request, err := http.NewRequest(
		"POST", c.url+":"+c.port, bytes.NewBuffer(jsonString),
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return response, nil
}

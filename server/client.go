package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//Client structure used to connect to the API server
type Client struct {
	hostname   string
	port       int
	httpClient *http.Client
}

//NewClient returns a new client configured to communicate on a server with the given hostname and port
func NewClient(hostname string, port int) *Client {
	return &Client{
		hostname:   hostname,
		port:       port,
		httpClient: &http.Client{},
	}
}

//GetProduct function retrieves a product json from the server API
func (c *Client) GetProduct(productID int) (*Product, error) {
	body := bytes.Buffer{}
	path := fmt.Sprintf("product/%d", productID)
	req, err := http.NewRequest("GET", c.requestPath(path), &body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		respBody := new(bytes.Buffer)
		_, err := respBody.ReadFrom(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("got a non 200 status code: %v", resp.StatusCode)
		}
		return nil, fmt.Errorf("got a non 200 status code: %v - %s", resp.StatusCode, respBody.String())
	}

	p := &Product{}
	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s:%v/%s", c.hostname, c.port, path)
}

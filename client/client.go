package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/vpovarna/go-mux-api/server"
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
func (c *Client) GetProduct(productID int) (*server.Product, error) {
	path := fmt.Sprintf("product/%d", productID)
	body, err := c.httpRequest(path, "GET", bytes.Buffer{})

	if err != nil {
		return nil, err
	}

	defer body.Close()
	product := &server.Product{}
	err = json.NewDecoder(body).Decode(product)

	if err != nil {
		return nil, err
	}
	return product, nil
}

//GetAllProducts function will return all products created on the server
func (c *Client) GetAllProducts() (*[]server.Product, error) {
	body, err := c.httpRequest("products", "GET", bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	products := []server.Product{}
	err = json.NewDecoder(body).Decode(&products)
	if err != nil {
		return nil, err
	}

	return &products, nil
}

//NewProduct function will create a new product through the server API
func (c *Client) NewProduct(product *server.Product) error {
	buffer := bytes.Buffer{}
	err := json.NewEncoder(&buffer).Encode(product)

	if err != nil {
		return err
	}

	_, err = c.httpRequest("product", "POST", buffer)

	if err != nil {
		return err
	}

	return nil
}

//DeleteProduct function will delete a product using server API
func (c *Client) DeleteProduct(productID int) error {
	path := fmt.Sprintf("product/%d", productID)

	_, err := c.httpRequest(path, "DELETE", bytes.Buffer{})

	if err != nil {
		return err
	}

	return nil
}

//UpdateProduct updates a product by ID through server API call
func (c *Client) UpdateProduct(product *server.Product) error {
	buffer := bytes.Buffer{}
	err := json.NewEncoder(&buffer).Encode(product)

	if err != nil {
		return err
	}

	path := fmt.Sprintf("product/%d", product.ID)
	_, err = c.httpRequest(path, "PUT", buffer)

	if err != nil {
		return err
	}
	return nil
}

func (c *Client) httpRequest(path string, method string, body bytes.Buffer) (closer io.ReadCloser, err error) {
	req, err := http.NewRequest(method, c.requestPath(path), &body)

	if err != nil {
		return nil, err
	}

	switch method {
	case "GET":
	case "DELETE":
	default:
		req.Header.Add("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}

func (c *Client) requestPath(path string) string {
	return fmt.Sprintf("%s:%v/%s", c.hostname, c.port, path)
}

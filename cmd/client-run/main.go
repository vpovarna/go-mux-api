package main

import (
	"fmt"
	"os"

	"github.com/vpovarna/go-mux-api/client"
	"github.com/vpovarna/go-mux-api/server"
)

var a server.App

const purgeTableQuery = `DELETE FROM products`
const resetPrimaryKey = `ALTER SEQUENCE products_id_seq RESTART WITH 1`

func main() {
	fmt.Println("Client run... ")

	fmt.Println("Cleannig up the DB!")
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
	clearDBTable()

	c := client.NewClient("http://localhost", 18010)
	getAllProducts(c)

	products := []server.Product{{ID: 1, Name: "Client Test Product", Price: 11.11}, {ID: 2, Name: "Seccond Client Test Product", Price: 12.11}}

	for _, product := range products {
		fmt.Printf("---Creating tests product: %v ---\n", product)
		err := c.NewProduct(&product)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Successfully created product: %v \n", product)
	}

	lastProduct := products[len(products)-1]
	fmt.Printf("---Geting product with id=%d ---\n", lastProduct.ID)

	p, err := c.GetProduct(lastProduct.ID)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(*p)

	getAllProducts(c)

	fmt.Printf("---Deleting product with id=%d ---\n", lastProduct.ID)
	err = c.DeleteProduct(lastProduct.ID)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Product with id: %d has been deleted successfully \n", lastProduct.ID)
	}

	getAllProducts(c)

	fmt.Println("---Updating Product---")
	seccondProduct := server.Product{ID: 1, Name: "Updated Client Product", Price: 12.11}
	err = c.UpdateProduct(&seccondProduct)
	if err != nil {
		fmt.Println(err.Error())
	}

	getAllProducts(c)

}

func getAllProducts(c *client.Client) {
	fmt.Println("---Geting all the products---")
	allProducts, err := c.GetAllProducts()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(*allProducts)
}

func clearDBTable() {
	a.DB.Exec(purgeTableQuery)
	a.DB.Exec(resetPrimaryKey)
}

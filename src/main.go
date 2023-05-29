package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// product represents data about a record product.
type product struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProductByID)
	router.POST("/products", postProducts)

	router.Run("localhost:8080")
}

// products slice to seed record product data.
var products = []product{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getProducts responds with the list of all products as JSON.
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, products)
}

// postProducts adds an product from JSON received in the request body.
func postProducts(c *gin.Context) {
	var newProduct product

	// Call BindJSON to bind the received JSON to
	// newProduct.
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}

	// Add the new product to the slice.
	products = append(products, newProduct)
	c.IndentedJSON(http.StatusCreated, newProduct)
}

// getProductByID locates the product whose ID value matches the id
// parameter sent by the client, then returns that product as a response.
func getProductByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of products, looking for
	// an product whose ID value matches the parameter.
	for _, a := range products {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

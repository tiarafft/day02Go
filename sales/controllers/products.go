package controllers

import (
	"log"
	"net/http"
	"time"
	// jis
	db "github.com/dtc/sales/connections"
	"github.com/dtc/sales/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

// CreateProduct : function to create product
func CreateProduct(c echo.Context) error {
	product := new(models.Product)
	if err1 := c.Bind(product); err1 != nil {
		return err1
	}
	// prepare sql statement
	con := db.GetConnection()
	sql := "INSERT INTO products(name, division, bundling, description, technology, warranty, demo, portofolio, note, created_at, updated_at) VALUES( ?, ?, ?, ?, ?, ?, ?, ?,?,?,?)"

	// Create a prepared SQL statement
	stmt, err := con.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		log.Fatal(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()

	// Replace the '?' in our prepared statement with 'argument'
	result, err2 := stmt.Exec(product.Name, product.Division, product.Bundling, product.Description, product.Technology, product.Warranty, product.Demo, product.Portofolio, product.Note, time.Now(), time.Now())

	// fmt.Println(result.LastInsertId.(string))
	// Exit if we get an error
	if err2 != nil {
		log.Fatal(err2)
	}
	lastInsert, err3 := result.LastInsertId()

	if err3 != nil {
		log.Fatal(err3)
	} else {
		// TODO: oop based on count pricing input
		pricing := &models.Pricing{
			ProductID: lastInsert,
		}
		models.AddPrPricing(pricing)
	}

	return c.JSON(http.StatusCreated, lastInsert)
}

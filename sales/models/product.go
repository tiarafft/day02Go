package models

import (
	"time"

	db "github.com/dtc/sales/connections"
)

// Product : is a struct for creating product at actual database
type Product struct {
	ID          int8      `json:"id" form:"id"`
	Name        string    `json:"name" form:"name"`
	Division    string    `json:"division" form:"division"`
	Bundling    int8      `json:"bundling" form:"bundling"`
	Description string    `json:"description" form:"description"`
	Technology  string    `json:"technology" form:"technology"`
	Warranty    string    `json:"warranty" form:"warranty"`
	Demo        string    `json:"demo" form:"demo"`
	Portofolio  string    `json:"portofolio" form:"portofolio"`
	Note        string    `json:"note" form:"note"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProductsCategories : using as present data from joining table
type ProductsCategories struct {
	Categories []Category `json:"categories"`
	Products   []Product  `json:"products"`
}

// ProductsMarkets : using as present data from joining table
type ProductsMarkets struct {
	Markets  []Market  `json:"markets"`
	Products []Product `json:"products"`
}

// Pricing : join table on database
type Pricing struct {
	ProductID int8   `json:"product_id"`
	PriceID   int8   `json:"price_id" form:"price_id"`
	Amount    string `json:"amount" form:"amount"`
}

// ProductPrices get
type ProductPrices struct {
	ID       int8      `json:"id"`
	Prices   []Price   `json:"prices"`
	Products []Product `json:"products"`
	Amount   []Pricing `json:"amount"`
}

// AddPrPricing : adding product and price
func AddPrPricing(*Pricing) (int64, error) {
	con := db.GetConnection()

	//
	sql := "INSERT INTO products_prices(product_id, pricing_id, amount) VALUES( ?, ?, ?)"

	// Create a prepared SQL statement
	stmt, err := con.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Make sure to cleanup after the program exits
	defer stmt.Close()
	price := &Pricing{}
	// Replace the '?' in our prepared statement with 'argument'
	result, err2 := stmt.Exec(price.ProductID, price.PriceID, price.Amount)

	// Exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()

}

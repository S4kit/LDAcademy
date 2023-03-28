package Models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	DB "GoAssig/GoAssig/Initializers"
)

type Product struct {
	ID          int    `json:"id"`
	ShopID      int    `json:"shop_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Categories  []int  `json:"categories"`
}

func AddProduct(product Product) (int64, error) {

	result, err := DB.Connect.Exec("INSERT INTO Product (ShopID, Name, Description, Categories) VALUES (?, ?, ?,?);", product.ShopID, product.Name, product.Description, product.Categories)
	if err != nil {
		return 0, fmt.Errorf("Couldn't add the Product : %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Couldn't add the Product: %v", err)
	}
	return id, nil
}
func ShowAllProducts() ([]Product, error) {
	var prod []Product
	result, err := DB.Connect.Query("SELECT * FROM Product")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find any Products %v", err)
	}
	defer result.Close()
	for result.Next() {
		var prd Product
		if err := result.Scan(&prd.ID, &prd.ShopID, &prd.Name, &prd.Description, &prd.Categories); err != nil {
			return nil, fmt.Errorf("Couldn't find any Products %v", err)
		}
		prod = append(prod, prd)
	}
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("Couldn't find any Products %v", err)
	}
	return prod, nil

}
func ShowProduct(id int) (Product, error) {
	var prod Product
	result := DB.Connect.QueryRow("SELECT * FROM Product WHERE ID=?", id)
	if err := result.Scan(&prod.ID, &prod.ShopID, &prod.Name, &prod.Description, prod.Categories); err != nil {
		if err == sql.ErrNoRows {
			return prod, fmt.Errorf("Couldn't find any Products %v", err)
		}
		return prod, fmt.Errorf("Couldn't find any Products %v", err)
	}
	return prod, nil
}
func EditProduct(id int, product Product) (bool, error) {

	result, err := DB.Connect.Query("UPDATE Products SET Name = ?, Description = ? , Categories = ? WHERE ShopID = ?;", product.Name, product.Description, product.Categories, id)

	if err != nil {

		return false, fmt.Errorf("Couldn't Edit any Products %v", err)
	}
	defer result.Close()
	return true, nil

}
func DeleteProduct(id int) (bool, error) {
	result, err := DB.Connect.Query("DELETE FROM Products WHERE ID=?;", id)

	if err != nil {

		return false, fmt.Errorf("Couldn't Delete any Products %v", err)
	}
	defer result.Close()
	return true, nil
}

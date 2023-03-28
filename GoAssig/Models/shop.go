package Models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	DB "GoAssig/GoAssig/Initializers"
)

type Shop struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func AddShop(shop Shop) (int64, error) {
	result, err := DB.Connect.Exec("INSERT INTO Shop (Name,Address ) VALUES (?, ?);", shop.Name, shop.Address)
	if err != nil {
		return 0, fmt.Errorf("Couldn't add the Shop : %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Couldn't add the Shop: %v", err)
	}
	return id, nil
}
func ShowAllShops() ([]Shop, error) {
	var shop []Shop
	result, err := DB.Connect.Query("SELECT * FROM Shop")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find any shop %v", err)
	}
	defer result.Close()
	for result.Next() {
		var shop1 Shop
		if err := result.Scan(&shop1.ID, &shop1.Name, &shop1.Address); err != nil {
			return nil, fmt.Errorf("Couldn't find any shop %v", err)
		}
		shop = append(shop, shop1)
	}
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("Couldn't find any shop %v", err)
	}
	return shop, nil
}

func ShowShop(id int) (Shop, error) {
	var shop Shop
	result := DB.Connect.QueryRow("SELECT * FROM Shop WHERE ID=?;", id)
	if err := result.Scan(&shop.ID, &shop.Name, &shop.Address); err != nil {
		if err == sql.ErrNoRows {
			return shop, fmt.Errorf("Couldn't find any Shop %v", err)
		}
		return shop, fmt.Errorf("Couldn't find any Shop %v", err)
	}
	return shop, nil
}

func EditShop(id int, shop Shop) (bool, error) {

	result, err := DB.Connect.Query("UPDATE Shop SET Name = ?, Address = ? WHERE ID=?;", shop.Name, shop.Address, id)

	if err != nil {
		return false, fmt.Errorf("Couldn't Edit any Shop %v", err)
	}
	defer result.Close()
	return true, nil
}

func DeleteShop(id int) (bool, error) {
	result, err := DB.Connect.Query("DELETE FROM Shop WHERE ID=?;", id)

	if err != nil {

		return false, fmt.Errorf("Couldn't Delete any Shop %v", err)
	}
	defer result.Close()
	return true, nil
}

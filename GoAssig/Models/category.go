package Models

import (
	DB "GoAssig/GoAssig/Initializers"
	"fmt"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ShowAllCat() ([]Category, error) {
	var cat []Category
	result, err := DB.Connect.Query("SELECT * FROM Category")
	if err != nil {
		return nil, fmt.Errorf("Couldn't find any Category %v", err)
	}
	defer result.Close()
	for result.Next() {
		var cate Category
		if err := result.Scan(&cate.ID, &cate.Name, &cate.Description); err != nil {
			return nil, fmt.Errorf("Couldn't find any Category %v", err)
		}
		cat = append(cat, cate)
	}
	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("Couldn't find any Category %v", err)
	}
	return cat, nil
}

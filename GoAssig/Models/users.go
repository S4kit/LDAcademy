package Models

import (
	DB "GoAssig/GoAssig/Initializers"
	"database/sql"
	"fmt"
)

type User struct {
	ID       string `form:"ID"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

func NewUser(user User) (int64, error) {

	result, err := DB.Connect.Exec("INSERT INTO User (Email,Password ) VALUES (?, ?);", user.Email, user.Password)
	if err != nil {
		return 0, fmt.Errorf("Couldn't add the User : %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Couldn't add the User: %v", err)
	}
	return id, nil
}

func AskUser(email string) (User, bool, error) {
	var user User
	result := DB.Connect.QueryRow("SELECT * FROM User WHERE Email=?;", email)
	if err := result.Scan(&user.ID, &user.Email, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return user, false, fmt.Errorf("Couldn't find any User %v", err)
		}
		return user, false, fmt.Errorf("Couldn't find any User %v", err)
	}
	return user, true, nil
}

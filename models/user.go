package models

import (
	"errors"
	"rest-api-project/db"
	"rest-api-project/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(email, password) values (?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := res.LastInsertId()

	u.Id = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password from users WHERE email = ?"

	var retreivePassword string
	row := db.DB.QueryRow(query, u.Email)
	err := row.Scan(&u.Id, &retreivePassword)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPassordHash(u.Password, retreivePassword)
	if !passwordIsValid {
		return errors.New("credentials Invalid")
	}
	return nil
}

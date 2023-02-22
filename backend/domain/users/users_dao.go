package users

import (
	"go-react-auth-demo/datasource/mysql/users_db"
	"go-react-auth-demo/utils/errors"
)

var (
	queryInsertUser     = "INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?);"
	queryGetUserByEmail = "SELECT id, first_name, last_name, email, password FROM users WHERE email = ?;"
	queryGetUserById    = "SELECT id, first_name, last_name, email FROM users WHERE id = ?;"
)

func (user *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)

	if saveErr != nil {
		return errors.NewInternalServerError("database error")
	}

	userId, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	user.ID = userId

	return nil

}

func (user *User) GetByEmail() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUserByEmail)
	if err != nil {
		return errors.NewInternalServerError("invalid email")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password); getErr != nil {
		return errors.NewInternalServerError("database errror")
	}

	return nil
}

func (user *User) GetById() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email); getErr != nil {
		return errors.NewInternalServerError("database error")
	}

	return nil
}

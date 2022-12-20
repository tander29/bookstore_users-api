package users

import (
	"bookstore_users-api/datasources/mysql/users_db"
	"bookstore_users-api/utils/date_utils"
	"bookstore_users-api/utils/errors"
	mysql_utils "bookstore_users-api/utils/mysql"
	"fmt"
)

// in this system data access object is only place used to access database
const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(? , ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)
	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (u *User) Save() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	u.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error trying to save user %s", err.Error()))
	}
	u.Id = userId

	return nil
}

func (u *User) Update() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

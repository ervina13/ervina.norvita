package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

type UserRepo struct {
	db *sql.DB
}

type RegisterRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

type UserRegis struct {
	db *sql.DB
}

func NewUserRegis(db *sql.DB) *UserRegis {
	return &UserRegis{db: db}
}

func (u *UserRepo) FecthUser() ([]User, error) {
	var sqlStmt string
	var users []User

	sqlStmt = "SELECT id, name, username, email, password, role, created_at FROM users;"

	rows, err := u.db.Query(sqlStmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var user User
	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
		)

		if err != nil {
			//panic(err)
			return nil, errors.New("No Data")
		}

		users = append(users, user)
	}

	return users, err
}

func (u *UserRepo) QueryUser(username string) User {
	var sqlStmt string
	var user User

	sqlStmt = `SELECT id, name, username, email, password, role FROM users WHERE username = ?;`

	row := u.db.QueryRow(sqlStmt, username)
	_ = row.Scan(&user.ID, &user.Username, &user.Password, &user.Role)

	return user
}

func (u *UserRepo) Login(username string, password string) (*string, error) {
	var user User
	err := u.db.QueryRow("SELECT username FROM users WHERE username = ? AND password = ?", username, password).Scan(&user.Username)
	if err != nil {
		return nil, errors.New("Login Failed")
	}
	return &user.Username, nil
}

func (u *UserRegis) Register(username string, password string, role string) (*string, error) {
	var user User
	_, err := u.db.Exec("INSERT INTO users (username, password,role) VALUES (?, ?,?)", username, password, role)
	if err != nil {
		fmt.Println(err, "error")
		return nil, errors.New("register failed")
	}
	return &user.Username, nil
}

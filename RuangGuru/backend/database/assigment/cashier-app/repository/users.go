package repository

import (
	"database/sql"
	//"errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) FetchUserByID(id int64) (User, error) {
	return User{}, nil // TODO: replace this
}

func (u *UserRepository) FetchUsers() ([]User, error) {
	return []User{}, nil // TODO: replace this
}

func (u *UserRepository) Login(username string, password string) (*string, error) {
	// err := u.db.Save(userDBName, updated)
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	// var loggedInUser *string

	// // load users
	// users, err := u.LoadOrCreate()
	// if err != nil {
	// 	return loggedInUser, err
	// }

	// user := searchUserByUsername(users, username)
	// if user == nil || user.Password != password {
	// return loggedInUser, errors.New("Login Failed")
	// }
	return nil, nil // TODO: replace this
}

func (u *UserRepository) InsertUser(username string, password string, role string, loggedin bool) error {
	return nil // TODO: replace this
}

func (u *UserRepository) FetchUserRole(username string) (*string, error) {
	return nil, nil // TODO: replace this
}

package repository

import (
	"fmt"

	"github.com/ruang-guru/playground/backend/golang-http-server/assignment/cashier-app/db"
)

type UserRepository struct {
	db db.DB
}

func NewUserRepository(db db.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) LoadOrCreate() ([]User, error) {
	return []User{}, nil // TODO: replace this
}

func (u *UserRepository) SelectAll() ([]User, error) {
	//return []User{}, nil // TODO: replace this
	records, _ := u.db.Load("users")
	result := make([]User, 0)
	var isLoggedIn bool
	for i := 1; i < len(records); i++ {
		if records[i][2] == "false" {
			isLoggedIn = false
		} else {
			isLoggedIn = true
		}
		user := User{
			Username: records[i][0],
			Password: records[i][1],
			Loggedin: isLoggedIn,
			Role:     records[i][3],
		}
		result = append(result, user)
	}
	return result, nil
}

func (u UserRepository) Login(username string, password string) (*string, error) {
	//return nil, nil // TODO: replace this
	records, _ := u.db.Load("users")
	updated := [][]string{
		{"username", "password", "loggedin", "role"},
	}
	var isUsernameAndPasswordFound bool
	var result *string
	for i := 1; i < len(records); i++ {
		records[i][2] = "false"
		if records[i][0] == username && records[i][1] == password {
			records[i][2] = "true"
			isUsernameAndPasswordFound = true
			result = &records[i][0]
		}
		updated = append(updated, []string{
			records[i][0],
			records[i][1],
			records[i][2],
			records[i][3],
		})
	}
	if isUsernameAndPasswordFound == false {
		return nil, fmt.Errorf("Login Failed")
	}
	err := u.db.Save("users", updated)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (u *UserRepository) Save(users []User) error {
	return nil // TODO: replace this
}

func (u *UserRepository) GetUserRole(username string) (*string, error) {
	// TODO: answer here
	records, _ := u.db.Load("users")
	for i := 1; i < len(records); i++ {
		if records[i][0] == username {
			return &records[i][3], nil
		}
	}
	return nil, fmt.Errorf("not found")

}

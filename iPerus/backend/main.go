package main

import (
	api "IPERPUS/controller"
	"IPERPUS/repository"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	usersRepo := repository.NewUserRepo(db)
	booksRepo := repository.NewBooksRepository(db)
	regisRepo := repository.NewUserRegis(db)
	addBooks := repository.AddNewBooks(db)
	findWriter := repository.FindaBook(db)
	profilUserRepo := repository.NewProfilUserRepository(db)
	profiladminRepo := repository.NewProfiladminRepository(db)
	editUsername := repository.UpdateaBook(db)

	mainApi := api.NewController(*usersRepo, *booksRepo, *regisRepo, *addBooks, *findWriter, *profilUserRepo, *profiladminRepo, *editUsername)
	mainApi.Start()
}

package repository

import (
	"database/sql"
	"errors"
	"fmt"
	//"github.com/klauspost/compress/gzhttp/writer"
	//"errors"
)

type BooksRepo struct {
	db *sql.DB
}

//type buku struct{Book}
func NewBooksRepository(db *sql.DB) *BooksRepo {
	return &BooksRepo{db: db}
}

type AddBooks struct {
	db *sql.DB
}

func AddNewBooks(db *sql.DB) *AddBooks {
	return &AddBooks{db: db}
}

type Finder struct {
	db *sql.DB
}

func FindaBook(db *sql.DB) *Finder {
	return &Finder{db: db}
}

type UpdateBook struct {
	db *sql.DB
}

func UpdateaBook(db *sql.DB) *UpdateBook {
	return &UpdateBook{db: db}
}
func (u *BooksRepo) ReadList() ([]Book, error) {
	rows, err := u.db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var sbuku []Book

	for rows.Next() {
		var pbuku Book
		err := rows.Scan(&pbuku.BookTitle, &pbuku.Writer, &pbuku.Tahun, &pbuku.Kategori, &pbuku.Link)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		sbuku = append(sbuku, pbuku)
	}
	return sbuku, nil
}

func (u *AddBooks) AddBooksList(title string, writer string, tahun string, kategori string, link string) (*string, error) {
	var buku Book
	_, err := u.db.Exec("INSERT INTO books (title,writer,tahun,kategori,link) VALUES(?,?,?,?,?)", title, writer, tahun, kategori, link)
	if err != nil {
		fmt.Println(err, "error")
		return nil, errors.New("Failed to add book")
	}
	return &buku.BookTitle, nil
}

func (u *Finder) FindBookBYWriter(writer string) ([]Book, error) {

	rows, err := u.db.Query("SELECT *FROM books WHERE writer LIKE '%" + writer + "%'")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sbuku []Book

	for rows.Next() {
		var pbuku Book
		err := rows.Scan(&pbuku.BookTitle, &pbuku.Writer, &pbuku.Tahun, &pbuku.Kategori, &pbuku.Link)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		sbuku = append(sbuku, pbuku)
	}
	return sbuku, nil

}
func (u *Finder) FindBookBYYear(tahun string) ([]Book, error) {

	rows, err := u.db.Query("SELECT *FROM books WHERE tahun LIKE '%" + tahun + "%'")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sbuku []Book

	for rows.Next() {
		var pbuku Book
		err := rows.Scan(&pbuku.BookTitle, &pbuku.Writer, &pbuku.Tahun, &pbuku.Kategori, &pbuku.Link)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		sbuku = append(sbuku, pbuku)
	}
	return sbuku, nil

}

func (u *Finder) FindBookBYKategori(kategori string) ([]Book, error) {

	rows, err := u.db.Query("SELECT *FROM books WHERE kategori LIKE '%" + kategori + "%'")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sbuku []Book

	for rows.Next() {
		var pbuku Book
		err := rows.Scan(&pbuku.BookTitle, &pbuku.Writer, &pbuku.Tahun, &pbuku.Kategori, &pbuku.Link)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		sbuku = append(sbuku, pbuku)
	}
	return sbuku, nil

}

func (u *BooksRepo) FetchByTitle(title string) ([]Book, error) {
	rows, err := u.db.Query("SELECT *FROM books WHERE title LIKE '%" + title + "%'")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sbuku []Book

	for rows.Next() {
		var pbuku Book
		err := rows.Scan(&pbuku.BookTitle, &pbuku.Writer, &pbuku.Tahun, &pbuku.Kategori, &pbuku.Link)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		sbuku = append(sbuku, pbuku)
	}
	return sbuku, nil

}

func (u *UpdateBook) UpdateBookTitles(title string) ([]Book, error) {
	//var edit Book
	rows, err := u.db.Query("UPDATE books SET title =? WHERE title =?")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var sbuku []Book

	for rows.Next() {
		var pbuku Book
		err := rows.Scan(&pbuku.BookTitle, &pbuku.Writer, &pbuku.Tahun, &pbuku.Kategori, &pbuku.Link)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		sbuku = append(sbuku, pbuku)
	}
	return sbuku, nil

}

package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
	CREATE TABLE books (
		title 		varchar(255) not null primary key,
		writer 		varchar(255) not null,
		tahun 		varchar(255) not null,
		kategori 	varchar(255) not null,
		link		varchar(255) not null
	);
	
	CREATE TABLE users (
		id integer not null primary key AUTOINCREMENT,
   		username varchar(255) not null,
    	password varchar(255) not null,
		role	varchar(255) not null
	);

	INSERT INTO books(title, writer,tahun,kategori,link) VALUES
    ('Harry Potter dan Batu Bertuah', 'J.K. Rowling','1997','Novel' ,'https://drive.google.com/file/d/1fRhMmDwT8AtUz9SZCuFKtwPqkbpVb1Jf/view?usp=sharing');

	INSERT INTO users(username, password,role) VALUES
    ('vina', '1313','admin'),
    ('kevin', '1212','user')`)

	if err != nil {
		panic(err)
	}
}

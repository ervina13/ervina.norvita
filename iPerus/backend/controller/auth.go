package api

import (
	//	"IPERPUS/repository"
	"IPERPUS/repository"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	//	"github.com/prometheus/client_golang/api"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `db:"role"`
}
type ResponsesTrue struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
type RegisterResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `db:"role"`
}
type AuthErrorResponse struct {
	Error string `json:"error"`
}

var jwtKey = []byte("key")

type Claims struct {
	Username string
	Role     string
	ID       int64
	jwt.StandardClaims
}
type Bookss struct {
	//	ID        int    `json:"id"`
	Title    string `json:"title"`
	Writer   string `json:"writer"`
	Tahun    string `json:"tahun"`
	Kategori string `json:"kategori"`
	Link     string `json:"link"`
}

type BookListSuccess struct {
	AllBooks []repository.Book `json: "books"`
}

func (api *controller) login(w http.ResponseWriter, req *http.Request) {

	var user User

	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := api.userRepository.Login(user.Username, user.Password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	claims := &Claims{
		Username:       *res,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
		Path:  "/",
	})
	json.NewEncoder(w).Encode(ResponsesTrue{Username: *res, Token: tokenString})

}

func (api *controller) signup(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := api.userRegister.Register(user.Username, user.Password, user.Role)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	claims := &Claims{
		Username:       *res,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: tokenString,
		Path:  "/",
	})
	json.NewEncoder(w).Encode(ResponsesTrue{Username: user.Username, Token: tokenString})

}

type AddToBookList struct {
	BookTitle string `json:"title`
}
type AddToBookListsuccess struct {
	BookTitle string `db:"title"`
	Writer    string `db:"writer"`
	Tahun     string `db:"tahun"`
	Kategori  string `db:"kategori"`
	Link      string `db:"link"`
}

func (api *controller) bookList(w http.ResponseWriter, r *http.Request) {
	//	var requestList AddToBookList
	buku, err := api.booksRepository.ReadList()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(r)

	// fetch, err := api.booksRepository.FetchByTitle(requestList.BookTitle)
	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	encoder.Encode(AuthErrorResponse{Error: err.Error()})
	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
	// encoder.Encode(AddToBookListsuccess{
	// 	BookTitle: fetch.BookTitle,

	// })

	defer func() {
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
	}()

	encoder.Encode(BookListSuccess{AllBooks: buku})
	return
}

func (api *controller) findBooksWriter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	writer := r.FormValue("writer")
	buku, err := api.findingBook.FindBookBYWriter(writer)
	if buku == nil {
		encoder.Encode(BookListSuccess{AllBooks: []repository.Book{}})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(r)

	defer func() {
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
	}()
	fmt.Println(buku)
	encoder.Encode(BookListSuccess{AllBooks: buku})

	return
}
func (api *controller) findBooksTahun(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	tahun := r.FormValue("tahun")
	buku, err := api.findingBook.FindBookBYYear(tahun)

	if buku == nil {
		encoder.Encode(BookListSuccess{AllBooks: []repository.Book{}})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(r)

	defer func() {
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
	}()
	fmt.Println(buku)
	encoder.Encode(BookListSuccess{AllBooks: buku})

	return
}

func (api *controller) findBooksKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	kategori := r.FormValue("kategori")
	buku, err := api.findingBook.FindBookBYKategori(kategori)

	if buku == nil {
		encoder.Encode(BookListSuccess{AllBooks: []repository.Book{}})
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(r)

	defer func() {
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
	}()
	fmt.Println(buku)
	encoder.Encode(BookListSuccess{AllBooks: buku})

	return
}

func (api *controller) addBooks(w http.ResponseWriter, r *http.Request) {
	var book Bookss

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := api.addBooksList.AddBooksList(book.Title, book.Writer, book.Tahun, book.Kategori, book.Link)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(res)

}

func (api *controller) dashboardUser(w http.ResponseWriter, r *http.Request) {
	showBookList, err := api.booksRepository.ReadList()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(r)

	defer func() {
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
	}()

	encoder.Encode(BookListSuccess{AllBooks: showBookList})
	return
}

func (api *controller) dashboadAdmin(w http.ResponseWriter, r *http.Request) {
	var book Bookss

	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := api.addBooksList.AddBooksList(book.Title, book.Writer, book.Tahun, book.Kategori, book.Link)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(*res)

}

func (api *controller) UpdateBooksTitle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	buku, err := api.editBookTitle.UpdateBookTitles(title)
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(r)

	defer func() {
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{Error: err.Error()})
			return
		}
	}()
	fmt.Println(buku)
	encoder.Encode(BookListSuccess{AllBooks: buku})

	return
}

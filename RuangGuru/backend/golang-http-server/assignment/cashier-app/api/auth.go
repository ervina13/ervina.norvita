package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginSuccessResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type AuthErrorResponse struct {
	Error string `json:"error"`
}

// Jwt key yang akan dipakai untuk membuat signature
var jwtKey = []byte("key")

//var jwtCookieKey = []byte("token")
//var claimsCtxKey = []byte("claims")

// Struct claim digunakan sebagai object yang akan di encode oleh jwt
// jwt.StandardClaims ditambahkan sebagai embedded type untuk provide standart claim yang biasanya ada pada JWT
type Claims struct {
	Username string
	Role     string
	jwt.StandardClaims
}

func (api *API) login(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	var user User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := api.usersRepo.Login(user.Username, user.Password)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		encoder.Encode(AuthErrorResponse{Error: err.Error()})
		return
	}

	userRole, err := api.usersRepo.GetUserRole(*res)

	// Task: 1. Deklarasi expiry time untuk token jwt
	//       2. Buat claim menggunakan variable yang sudah didefinisikan diatas
	//       3. expiry time menggunakan time millisecond

	// TODO: answer here
	expTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: user.Username,
		Role:     *userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}
	// Task: Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai

	// TODO: answer here
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Task: 1. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
	//       2. return internal error ketika ada kesalahan ketika pembuatan JWT string

	// TODO: answer here
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Task: Set token string kedalam cookie response

	// TODO: answer here
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenStr,
		Expires: expTime,
	})
	// Task: Return response berupa username dan token JWT yang sudah login

	json.NewEncoder(w).Encode(LoginSuccessResponse{Username: user.Username, Token: tokenStr}) // TODO: replace this
}

func (api *API) logout(w http.ResponseWriter, req *http.Request) {
	api.AllowOrigin(w, req)
	encoder := json.NewEncoder(w)

	token, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// return unauthorized ketika token kosong
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// return bad request ketika field token tidak ada
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if token.Value == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	c := http.Cookie{
		Name:   "token",
		MaxAge: -1,
	}
	http.SetCookie(w, &c)

	encoder.Encode(AuthErrorResponse{Error: ""}) // TODO: replace this
}

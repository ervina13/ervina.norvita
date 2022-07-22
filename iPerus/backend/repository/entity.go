package repository

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Token    string `db:"token"`
}
type Book struct {
	//ID        int64  `db:"id"`
	BookTitle string `db:"title"`
	Writer    string `db:"writer"`
	Tahun     string `db:"tahun"`
	Kategori  string `db:"kategori"`
	Link      string `db:"link"`
}
type ProfilUser struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	Gender    string `db:"gender"`
	Birthdate string `db:"birthdate"`
	Address   string `db:"address"`
	NoHp      string `db:"nohp"`
	UserID    int64  `db:"user_id"`
}

type Profiladmin struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Address  string `db:"address"`
	NoHp     string `db:"nohp"`
	Username string `db:"username"`
	UserID   int64  `db:"user_id"`
}

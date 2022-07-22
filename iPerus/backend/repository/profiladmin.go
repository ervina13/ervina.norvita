package repository

import (
	"database/sql"
	"errors"
)

type ProfiladminReposistory struct {
	db *sql.DB
}

func NewProfiladminRepository(db *sql.DB) *ProfiladminReposistory {
	return &ProfiladminReposistory{db: db}
}

func (p *ProfiladminReposistory) FecthProfil() ([]Profiladmin, error) {
	var sqlStmt string
	var profils []Profiladmin

	sqlStmt = "SELECT id, name, email, address, nohp, instansi, user_id FROM profil;"

	rows, err := p.db.Query(sqlStmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var profil Profiladmin
	for rows.Next() {
		err := rows.Scan(
			&profil.ID,
			&profil.Name,
			&profil.Email,
			&profil.Address,
			&profil.NoHp,
			&profil.UserID,
		)

		if err != nil {
			return nil, errors.New("No Data")
		}

		profils = append(profils, profil)
	}

	return profils, err
}

func (p *ProfiladminReposistory) FetchProfilByUserID(userID int64) (Profiladmin, error) {
	var profil Profiladmin
	var sqlStmt string

	sqlStmt = `SELECT p.id, p.name, p.email, p.address, p.nohp, p.instansi, u.username, p.user_id
	FROM profil p
	JOIN users u ON p.user_id = u.id
	WHERE p.user_id = ?
	LIMIT 1;`

	row := p.db.QueryRow(sqlStmt, userID)
	err := row.Scan(
		&profil.ID,
		&profil.Name,
		&profil.Email,
		&profil.Address,
		&profil.NoHp,
		&profil.Username,
		&profil.UserID,
	)
	if err != nil {
		return profil, err
	}

	return profil, nil
}

func (p *ProfiladminReposistory) InsertProfiladmin(profil Profiladmin) error {
	var sqlStmt string

	sqlStmt = "INSERT INTO profil (name, email, address, nohp, instansi, user_id) VALUES (?, ?, ?, ?, ?, ?);"

	_, err := p.db.Exec(sqlStmt, profil.Name, profil.Email, profil.Address, profil.NoHp, profil.UserID)

	if err != nil {
		panic(err)
	}

	return nil
}

func (p *ProfiladminReposistory) UpdateProfiladmin(profil Profiladmin) error {
	var sqlStmt string

	p.FecthProfil()

	sqlStmt = "UPDATE profil SET name = ?, email = ?, address = ?, nohp = ?, instansi = ? WHERE id = ?;"

	_, err := p.db.Exec(sqlStmt, profil.ID)

	if err != nil {
		panic(err)
	}

	return nil
}

func (p *ProfiladminReposistory) FetchProfilID(userID int64) (Profiladmin, error) {
	var sqlStmt string
	var profil Profiladmin

	sqlStmt = "SELECT id FROM profil WHERE user_id = ?;"

	row := p.db.QueryRow(sqlStmt, userID)
	err := row.Scan(&profil.ID)

	return profil, err
}
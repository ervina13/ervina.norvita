package main

import (
	"encoding/json"
	"fmt"
)

// Dari contoh yang telah diberikan, cobalah untuk melakukan encode struct menjadi json.
// Lengkapi function EncodeToJson agar dapat mengembalikan nilai byte hasil dari encode objek Leaderboard.
// Modifikasi struct UserRank sehingga field Name menjadi name ,field Rank menjadi rank, dan field Email tidak ikut untuk diencode.

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	Name  string `json:"name"`
	Email string `json:"-"`
	Rank  int    `json:"rank"`
}

type Leaderboard struct {
	Users []*UserRank
}

func EncodeToJson(leaderboard Leaderboard) ([]byte, error) {
	// TODO: answer here
	jsonData, _ := json.Marshal(leaderboard{
		Name:  "vina",
		Email: "eeeee@gmail.com",
		Rank:  2
	})
	jsonString := string(jsonData)
	fmt.Println(jsonString)
}

package seed

import (
	"log"

	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/config"
	"github.com/afrizal423/Golang-Perpustakaan-Restful-API/structs"
)

var jenbuk = []structs.Jenis_Buku{
	structs.Jenis_Buku{
		JenisBuku: "History",
		Deskripsi: "Ini adalah jenis buku history",
	},
	structs.Jenis_Buku{
		JenisBuku: "Cerpen",
		Deskripsi: "Ini adalah jenis buku Cerpen",
	},
	structs.Jenis_Buku{
		JenisBuku: "Kepercayaan",
		Deskripsi: "Ini adalah jenis buku Kepercayaan",
	},
	structs.Jenis_Buku{
		JenisBuku: "Politik",
		Deskripsi: "Ini adalah jenis buku Politik",
	},
	structs.Jenis_Buku{
		JenisBuku: "Ujian",
		Deskripsi: "Ini adalah jenis buku Ujian",
	},
}

var penbuk = []structs.Penulis_Buku{
	structs.Penulis_Buku{
		PenulisBuku:   "Penulis 1",
		AlamatPenulis: "Kota 1",
		EmailPenulis:  "test1@test.com",
	},
	structs.Penulis_Buku{
		PenulisBuku:   "Penulis 2",
		AlamatPenulis: "Kota 2",
		EmailPenulis:  "test2@test.com",
	},
	structs.Penulis_Buku{
		PenulisBuku:   "Penulis 3",
		AlamatPenulis: "Kota 3",
		EmailPenulis:  "test3@test.com",
	},
	structs.Penulis_Buku{
		PenulisBuku:   "Penulis 4",
		AlamatPenulis: "Kota 4",
		EmailPenulis:  "test4@test.com",
	},
	structs.Penulis_Buku{
		PenulisBuku:   "Penulis 5",
		AlamatPenulis: "Kota 5",
		EmailPenulis:  "test5@test.com",
	},
}

var PubBuk = []structs.Penerbit_Buku{
	structs.Penerbit_Buku{
		PenerbitBuku:   "Penerbit 1",
		AlamatPenerbit: "Kota 1",
		EmailPenerbit:  "publisher1@test.com",
	},
	structs.Penerbit_Buku{
		PenerbitBuku:   "Penerbit 2",
		AlamatPenerbit: "Kota 2",
		EmailPenerbit:  "publisher2@test.com",
	},
	structs.Penerbit_Buku{
		PenerbitBuku:   "Penerbit 3",
		AlamatPenerbit: "Kota 3",
		EmailPenerbit:  "publisher3@test.com",
	},
	structs.Penerbit_Buku{
		PenerbitBuku:   "Penerbit 4",
		AlamatPenerbit: "Kota 4",
		EmailPenerbit:  "publisher4@test.com",
	},
	structs.Penerbit_Buku{
		PenerbitBuku:   "Penerbit 5",
		AlamatPenerbit: "Kota 5",
		EmailPenerbit:  "publisher5@test.com",
	},
}

var dtbuku = []structs.Buku{
	structs.Buku{
		ISBN:            "INIISBNKE1",
		IDKategoriJenis: 1,
		Judul:           "Ini buku ke 1",
		IDPenulisBuku:   1,
		IDPenerbitBuku:  1,
		ThnTerbit:       "2020",
		StokBuku:        10,
		RakBuku:         "A01",
		DeskripsiBuku:   "Ini deskripsi kali ya, bagian buku 1",
	},
	structs.Buku{
		ISBN:            "INIISBNKE2",
		IDKategoriJenis: 2,
		Judul:           "Ini buku ke 2",
		IDPenulisBuku:   2,
		IDPenerbitBuku:  2,
		ThnTerbit:       "2020",
		StokBuku:        10,
		RakBuku:         "A02",
		DeskripsiBuku:   "Ini deskripsi kali ya, bagian buku 2",
	},
	structs.Buku{
		ISBN:            "INIISBNKE3",
		IDKategoriJenis: 3,
		Judul:           "Ini buku ke 3",
		IDPenulisBuku:   3,
		IDPenerbitBuku:  3,
		ThnTerbit:       "2020",
		StokBuku:        10,
		RakBuku:         "A03",
		DeskripsiBuku:   "Ini deskripsi kali ya, bagian buku 3",
	},
	structs.Buku{
		ISBN:            "INIISBNKE4",
		IDKategoriJenis: 4,
		Judul:           "Ini buku ke 4",
		IDPenulisBuku:   4,
		IDPenerbitBuku:  4,
		ThnTerbit:       "2020",
		StokBuku:        10,
		RakBuku:         "A04",
		DeskripsiBuku:   "Ini deskripsi kali ya, bagian buku 4",
	},
	structs.Buku{
		ISBN:            "INIISBNKE5",
		IDKategoriJenis: 5,
		Judul:           "Ini buku ke 5",
		IDPenulisBuku:   5,
		IDPenerbitBuku:  5,
		ThnTerbit:       "2020",
		StokBuku:        10,
		RakBuku:         "A05",
		DeskripsiBuku:   "Ini deskripsi kali ya, bagian buku 5",
	},
}

var db = config.Db

func Seed_JenBuk() {
	// masukkan ke db
	for i, _ := range jenbuk {
		err := db.Debug().Model(&structs.Jenis_Buku{}).Create(&jenbuk[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data: %v", err)
		}
	}
}

func Seed_PenBuk() {
	// masukkan ke db
	for i, _ := range penbuk {
		err := db.Debug().Model(&structs.Penulis_Buku{}).Create(&penbuk[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data: %v", err)
		}
	}
}

func Seed_PubBuk() {
	// masukkan ke db
	for i, _ := range PubBuk {
		err := db.Debug().Model(&structs.Penerbit_Buku{}).Create(&PubBuk[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data: %v", err)
		}
	}
}

func Seed_Buku() {
	// masukkan ke db
	for i, _ := range PubBuk {
		err := db.Debug().Model(&structs.Buku{}).Create(&dtbuku[i]).Error
		if err != nil {
			log.Fatalf("tidak bisa seed data: %v", err)
		}
	}
}

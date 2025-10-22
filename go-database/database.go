package belajar_golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

/**
Dengan pooling:

- Sekali buka koneksi, disimpan di memori.
- Saat ada query baru, ambil koneksi yang idle (menganggur).
- Setelah query selesai, koneksi dikembalikan ke pool.

benar-benar mirip dengan konsep pooling di goroutines dimana data bisa
digunakan dan dikembalikan lalu digunakan ulang

**/

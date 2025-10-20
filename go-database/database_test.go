package belajar_golang_database

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {

}

/**
	package database/sql merupakan seuah interface yang
	mendefinisakn kontrak yang harus dimiliki driver agar bisa digunakan untuk berinteraksi dengan dbms
**/

// dengan adanya package database maka kita bisa menggunakan func yang sama untuk menjalankan database apapun

// karena prinsip dari interface sendiri adalah polimorfism

// yang berbeda hanyalah nama driver dan kode sqlnya

// buktinya kita menggunakan metode dari package sql seperti sql.Open dll

func TestOpenConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// gunakan DB
}

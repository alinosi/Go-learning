package belajar_golang_database

import (
	"fmt"
	"testing"
)

func TestPooling(t *testing.T){
	fmt.Println("test dimulai")
}

/**
Dengan pooling:

- Sekali buka koneksi, disimpan di memori.
- Saat ada query baru, ambil koneksi yang idle (menganggur).
- Setelah query selesai, koneksi dikembalikan ke pool.

benar-benar mirip dengan konsep pooling di goroutines dimana data bisa 
digunakan dan dikembalikan lalu digunakan ulang

**/
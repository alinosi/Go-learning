package belajar_golang_web

import (
	"net/http"
	"testing"
)

// menjalankan web server built-in di golang

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080", // definisikan IP dan port yang akan digunakan layanan web server
	}

	// banyak attribut lain yang digunakan oleh server, untuk sementara gunakan addr dulu agar web server aktif

	err := server.ListenAndServe() // method ini akan mengaktifkan web server
	if err != nil {
		panic(err)
	}
}

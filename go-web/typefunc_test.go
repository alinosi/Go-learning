package belajar_golang_web

import (
	"fmt"
	"testing"
)

type mahasiswa interface {
	ambilnama()
}

type Unsri struct {
	name string
}

func (u *Unsri) changename(s string) {
	u.name = s
}

type ambilnim func(a string)

func (a ambilnim) ambilnama() {
	fmt.Println("func sebagai ambilnim berhasil dilakukan")
}

func milikMahasiswa(m mahasiswa) {
	m.ambilnama()
}

func TestType(t *testing.T) {
	var andhika Unsri
	andhika.changename("daniel")

	fmt.Println(andhika)
}

package main_test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

//go:embed version.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("logo_new.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/file_1.txt
//go:embed files/file_2.txt
//go:embed files/file_3.txt
//go:embed files/file_4.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/file_1.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/file_2.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/file_3.txt")
	fmt.Println(string(c))

	d, _ := files.ReadFile("files/file_4.txt")
	fmt.Println(string(d))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}

// in-one case 

//go:embed *.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
    dirEntries, _ := path.ReadDir(".") // titik artinya direktori utama (root dari embed)
    for _, entry := range dirEntries {
        if !entry.IsDir() {
            fmt.Println(entry.Name())
            file, _ := path.ReadFile(entry.Name()) // tidak perlu "files/"
            fmt.Println(string(file))
        }
    }
}
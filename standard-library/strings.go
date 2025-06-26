package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Nobynobygon", "Noby"))                           // apakah string mengandung kata "Noby"
	fmt.Println(strings.Split("Nobynobygon", " "))                                 // memecah setiap kata yang dipisahkan space menjadi slice
	fmt.Println(strings.ToLower("Nobynobygon"))                                    // konversi ke huruf kecil
	fmt.Println(strings.ToUpper("Nobynobygon"))                                    // konversi ke huruf besar
	fmt.Println(strings.Trim("      Nobynobygon      ", " "))                      // menghapus karakter pilihan di awal dan akhir kata
	fmt.Println(strings.ReplaceAll("Nobynobygon Noby nobygon", "Noby", "alinosi")) // mengganti setiap kata Noby menjadi alinosi
}
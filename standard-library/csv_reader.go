package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "sir,noby,bitzer\n" +
		"sir,hikaru,tcahaya\n" +
		"sir anton nugroho"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF { // EOF = end file(file selesai dibaca/sudah di baris terakhir)
			break
		}

		fmt.Println(record)
	}
}

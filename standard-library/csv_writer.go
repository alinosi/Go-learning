package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"sir", "noby", "bitzer"})
	_ = writer.Write([]string{"sir", "hikaru", "tcahaya"})
	_ = writer.Write([]string{"sir", "anton", "nugroho"})

	writer.Flush()
}

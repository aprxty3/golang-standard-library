package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	name := "Dwi Aji Prasetyo"

	encodeBase64 := base64.StdEncoding.EncodeToString([]byte(name))
	fmt.Println(encodeBase64)

	decodeBase64, err := base64.StdEncoding.DecodeString(encodeBase64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(decodeBase64))
	}

	fmt.Println("=====================================")

	csvReader := "Halo,Nama,Saya \n" +
		"Dwi,Aji,Prasetyo \n" +
		"Saya,Sedang,Belajar,Go"

	reader := csv.NewReader(strings.NewReader(csvReader))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}

	fmt.Println("=====================================")

	csvWriter := csv.NewWriter(os.Stdout)

	csvWriter.Write([]string{"Halo", "Nama", "Saya"})
	csvWriter.Write([]string{"Dwi", "Aji", "Prasetyo"})
	csvWriter.Write([]string{"Saya", "Sedang", "Belajar", "Go"})

	csvWriter.Flush()
}

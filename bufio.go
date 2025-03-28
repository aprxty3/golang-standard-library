package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader("this a long string \nwith multiple lines \nand many words")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))

	}

	fmt.Println("=====================================")

	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Halo, Nama, Saya \n")
	_, _ = writer.WriteString("Dwi, Aji, Prasetyo \n")
	_, _ = writer.WriteString("Saya, Sedang, Belajar, Go \n")

	writer.Flush()
}

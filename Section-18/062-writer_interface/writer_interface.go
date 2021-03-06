package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	filePath := "./writes/test.go"

	n, err := fileWriter(filePath)
	if err != nil {
		log.Fatalln("error:", err)
	}

	fmt.Printf("Wrote %v bytes to %s\n", n, filePath)

	fmt.Fprintln(os.Stdout, "Hello there")

	io.WriteString(os.Stdout, "Hello there")
}

func fileWriter(path string) (int, error) {
	file, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	bs, err := fileReader("./writer_interface.go")
	if err != nil {
		return 0, err
	}

	n, err := file.Write(bs)
	if err != nil {
		return n, err
	}

	return n, nil
}

func fileReader(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	bs := make([]byte, 1024)

	m, err := file.Read(bs)
	if err != nil {
		return bs[:m], err
	}

	return bs[:m], nil
}

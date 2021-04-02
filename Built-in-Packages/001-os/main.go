package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// read main.go file
	readFileOS("./main.go", 2048)

	// read file.html file
	readFileOS("./file.html", 1048)

	// read main.go file
	// readFileIO("./main.go")

	// read file.html file
	// readFileIO("./file.html")

	// read main.go
	// readFileChunk("./main.go")

	// read file.html file
	// readFileChunk("./file.html")

	// read main.go
	// readFileBufio("./main.go")

	// read file.html
	// readFileBufio("./file.html")
}

func readFileOS(path string, size int32) {
	myFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()

	bs := make([]byte, size)
	n, err := myFile.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Size: %v Bytes\nContent:\n%s\n", n, bs)
	fmt.Println()
}

func readFileIO(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Size: %v Bytes\nContent:\n%s\n", len(content), content)
	fmt.Println()
}

func readFileChunk(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// declare chunk size
	const maxSize = 4

	// create a buffer
	b := make([]byte, maxSize)

	content := []byte{}

	for {
		// read content to buffer
		readTotal, err := file.Read(b)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		content = append(content, b[:readTotal]...)
	}

	fmt.Printf("Size: %v Bytes\nContent:\n%s\n", len(content), content)
}

func readFileBufio(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	content := []byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = append(content, []byte(fmt.Sprintf("%v\n", scanner.Text()))...)
	}

	fmt.Printf("Size: %v Bytes\nContent:\n%s\n", len(content), content)
}

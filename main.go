package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// read a go file

func main() {

	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)

	words := make(map[string]int)
	for scan.Scan() {
		words[scan.Text()]++
	}

	fmt.Println(words)

}

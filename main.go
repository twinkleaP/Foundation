package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	fmt.Println(words) // print map

	type data struct {
		key   string
		value int
	}

	var v []data

	for key, value := range words {
		v = append(v, data{key, value})
	}

	sort.Slice(v, func(i int, j int) bool {

		return v[i].value > v[j].value

	})
	for _, s := range v {
		fmt.Println(s.key, "appears", s.value, "times")
	}

}

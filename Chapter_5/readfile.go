package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("Chapter_5\\data.txt")
	fmt.Println(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	fmt.Println(scanner)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

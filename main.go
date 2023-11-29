package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func popcount(str string, words []string) (string, error) {
	var v string

	for _, word := range words {
		if lev(word, str) < 2 {
			v = word
		}
	}

	if v == "" {
		return "", errors.New("No word was found")
	}

	return v, nil
}

func main() {
	flag.Parse()

	word := flag.Arg(0)
	if word == "" {
		log.Fatal("No word is provided")
	}

	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	words := strings.Split(string(data), "\n")

	v, err := popcount(word, words)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("maybe you wanted %s instead of %s\n", v, word)
}

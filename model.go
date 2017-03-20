package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	model := load("model.txt")
	fmt.Println(len(model))
}

type WordVector struct {
	word    string
	vectors []float32
}

func load(filename string) []WordVector {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	result := []WordVector{}

	scanner := bufio.NewScanner(file)
	firstLine := true
	for scanner.Scan() {
		var line = strings.Split(scanner.Text(), " ")
		if firstLine {
			fmt.Println("words " + line[0])
			fmt.Println("vectors " + line[1])
			firstLine = false
		} else {
			wv := parseLine(line)
			result = append(result, wv)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func parseLine(line []string) WordVector {
	vector := []float32{}
	for i := 1; i < len(line); i++ {
		val, _ := strconv.ParseFloat(line[i], 32)
		vector = append(vector, float32(val))
	}
	return WordVector{line[0], vector}
}

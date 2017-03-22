package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

type WordVector struct {
	word    string
	vectors []float32
}

type WordVectors []WordVector

func main() {
	model := load("model.txt")

	found, vector := model.FindVector("whale")

	if found {
		fmt.Println("found", vector.word)
		vector.Add(vector)
	}


	fmt.Println(len(model))
}


func load(filename string) WordVectors {
	file, err := os.Open(filename)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	result := WordVectors{}

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


func (vectors WordVectors) FindVector(word string) (bool, WordVector){
	for _,vector := range vectors{
		if vector.word == word {
			return true, vector
		}
	}
	return false, WordVector{word, []float32{}}
}

func (vector2 WordVector) Add(vector1 WordVector) WordVector {
	returnArray := []float32{}
	for i := 0; i < len(vector1.vectors); i++{
		returnArray = append(returnArray, vector1.vectors[i] + vector2.vectors[i])
	}
	return WordVector{ vector1.word + " + " + vector2.word, returnArray}
}

func (vector2 WordVector) Subtract(vector1 WordVector) WordVector {
	returnArray := []float32{}
	for i := 0; i < len(vector1.vectors); i++{
		returnArray = append(returnArray, vector1.vectors[i] - vector2.vectors[i])
	}
	return WordVector{ vector1.word + " - " + vector2.word, returnArray}
}

func (vector2 WordVector) Distance(vector1 WordVector) float32 {
	distance := float64(0)
	for i := 0; i < len(vector1.vectors); i++{
		distance += math.Pow(float64(vector1.vectors[i] - vector2.vectors[i]), float64(2))
	}
	return float32(math.Sqrt(distance))
}

func (vectors WordVectors) Nearest(vector WordVector) (bool, WordVector){
	if (len(vectors) == 0){
		return false, vector
	}

	shortestDistance := math.MaxFloat64
	shortestVector := vector

	for _,vector2 := range vectors{
		if vector2.word == vector2.word {
			continue
		}

		thisDistance := float64(vector.Distance(vector2))
		if (thisDistance < shortestDistance){
			shortestDistance = thisDistance
			shortestVector = vector2
		}
	}
	return true, shortestVector
}

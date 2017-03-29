package main

import (
  "fmt"
)

func main() {
	model := Load("glove.6B.300d.txt")

	_, king := model.FindVector("king")
	_, man := model.FindVector("man")
	_, queen := model.FindVector("queen")
	newVector := king.Subtract(man).Add(queen)
	_, nearest := model.Nearest(newVector)
	fmt.Println(nearest.word)

	_, face := model.FindVector("spoon")
	_, nearestToFace := model.Nearest(face)
	fmt.Println(nearestToFace.word)

}

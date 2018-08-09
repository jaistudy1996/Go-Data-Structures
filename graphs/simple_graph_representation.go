package main

import ("fmt")

func main() {
	createMatrix(6, 6)
}

func createMatrix(vertices, edges int) [][]int {
	return make([][]int, vertices)
}

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

	file, err := os.Open("prob67/0067_triangle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var triangle [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		numbers := make([]int, len(parts))
		for i, s := range parts {
			numbers[i], _ = strconv.Atoi(s)
		}
		triangle = append(triangle, numbers)
	}

	for row := len(triangle) - 2; row >= 0; row-- {
		for i := 0; i < len(triangle[row]); i++ {
			triangle[row][i] += max(triangle[row+1][i], triangle[row+1][i+1])
		}
	}

	fmt.Println(triangle[0][0])
}

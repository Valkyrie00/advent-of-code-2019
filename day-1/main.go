package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := reader("inputs.txt")

	var fuelSum, totalSum int
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Panicln("Error mass coversion", err)
		}

		fuelSum += calculate(mass)
		reduce(mass, &totalSum)
	}

	// Part 1
	fmt.Printf("PART 1 - Fuel sum: %v \n", fuelSum)

	// Part 2
	fmt.Printf("PART 2 - Total Fuel: %v \n", totalSum)
}

func reduce(mass int, totalSum *int) {
	reduced := calculate(mass)
	if reduced < 0 {
		return
	}

	*totalSum += reduced
	reduce(reduced, totalSum)
}

func calculate(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

// Input reader
func reader(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return bufio.NewScanner(file)
}

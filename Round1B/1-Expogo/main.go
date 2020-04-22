package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getNextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	i, _ := strconv.Atoi(string(scanner.Bytes()))
	return i
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	nbCase := getNextInt(scanner)

	for caseId := 1; caseId <= nbCase; caseId++ {
		x := int64(getNextInt(scanner))
		y := int64(getNextInt(scanner))

		result, err := findBestSolution(x, y)

		if err != nil {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", caseId)
		} else {
			fmt.Printf("Case #%d: %s\n", caseId, result)
		}

	}
}

func findBestSolution(x, y int64) (string, error) {

	if x == 0 && y == 0 {
		return "", nil
	}

	absSum := int64(math.Abs(float64(x)) + math.Abs(float64(y)))
	if Even(absSum) {
		return "", fmt.Errorf("no solution")
	}

	if Even(x) {
		newX := x / 2
		if newX == 0 && y == 1 {
			return "N", nil
		}
		if newX == 0 && y == -1 {
			return "S", nil
		}
		if (Even(newX) && Odd((y+1)/2)) || (Odd(newX) && Even((y+1)/2)) {
			res, err := findBestSolution(newX, (y+1)/2)
			return "S" + res, err
		} else if (Even(newX) && Odd((y-1)/2)) || (Odd(newX) && Even((y-1)/2)) {
			res, err := findBestSolution(newX, (y-1)/2)
			return "N" + res, err
		}
	}
	if Even(y) {
		newY := y / 2
		if newY == 0 && x == 1 {
			return "E", nil
		}
		if newY == 0 && x == -1 {
			return "W", nil
		}
		if (Even(newY) && Odd((x+1)/2)) || (Odd(newY) && Even((x+1)/2)) {
			res, err := findBestSolution((x+1)/2, newY)
			return "W" + res, err
		} else if (Even(newY) && Odd((x-1)/2)) || (Odd(newY) && Even((x-1)/2)) {
			res, err := findBestSolution((x-1)/2, newY)
			return "E" + res, err
		}
	}
	return "", fmt.Errorf("no solution")
}

func Even(number int64) bool {
	return number%2 == 0
}

func Odd(number int64) bool {
	return !Even(number)
}

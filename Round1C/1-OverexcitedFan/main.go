package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getNextInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	i, _ := strconv.Atoi(string(scanner.Bytes()))
	return i
}

func getNextString(scanner *bufio.Scanner) string {
	scanner.Scan()
	i := string(scanner.Bytes())
	return i
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	nbCase := getNextInt(scanner)

	for caseId := 1; caseId <= nbCase; caseId++ {
		result, err := findSolution(scanner)

		if err != nil {
			fmt.Printf("Case #%d: IMPOSSIBLE\n", caseId)
		} else {
			fmt.Printf("Case #%d: %s\n", caseId, result)
		}

	}
}

func abs(n int) int {
	y := n >> 31
	return (n ^ y) - y
}

func distance(x, y int) int {
	return abs(x) + abs(y)
}

func findSolution(scanner *bufio.Scanner) (string, error) {

	//4 4 SSSS

	x := getNextInt(scanner)
	y := getNextInt(scanner)
	path := getNextString(scanner)

	for i, d := range path {
		switch d {
		case 'N':
			y += 1
		case 'S':
			y -= 1
		case 'E':
			x += 1
		case 'W':
			x -= 1
		}
		//fmt.Println("Test :", i,  distance(x, y), x, y)
		if distance(x, y) <= i+1 {
			return strconv.Itoa(i + 1), nil
		}
	}

	//	return strings.Trim(strings.Replace(fmt.Sprint(myTab), " ", ",", -1), "[]"), nil
	return "", fmt.Errorf("no solution")
}

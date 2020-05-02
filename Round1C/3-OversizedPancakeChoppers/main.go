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

func getNextInt64(scanner *bufio.Scanner) int64 {
	scanner.Scan()
	i, _ := strconv.ParseInt(string(scanner.Bytes()), 10, 64)
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

func findSolution(scanner *bufio.Scanner) (string, error) {

	nbSlices := getNextInt(scanner)
	nbDiners := getNextInt(scanner)

	maxSize := int64(0)
	masSizeDouble := int64(0)
	maxSame := 0
	sliceBySize := make(map[int64]int)
	for index := 0; index < nbSlices; index++ {
		sliceSize := getNextInt64(scanner)
		nbSame, ok := sliceBySize[sliceSize]
		if !ok {
			nbSame = 1
			sliceBySize[sliceSize] = 1
		} else {
			nbSame++
			sliceBySize[sliceSize] = nbSame
		}
		if nbSame > maxSame {
			maxSame = nbSame
			masSizeDouble = sliceSize
		}
		if sliceSize > maxSize {
			maxSize = sliceSize
		}
	}

	if maxSame >= nbDiners {
		return "0", nil
	}

	if nbDiners == 2 {
		return "1", nil
	}

	if maxSame == 2 && masSizeDouble != maxSize {
		return "1", nil
	}

	for size, _ := range sliceBySize {
		for otherSize, _ := range sliceBySize {
			if (otherSize%2) == int64(0) && size == otherSize/2 {

				return "1", nil
			}
		}
	}

	return "2", nil
}

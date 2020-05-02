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

func findSolution(scanner *bufio.Scanner) (string, error) {

	maxPossible := make(map[string]int)
	carNotUse := make(map[string]int)

	power := getNextInt(scanner)
	//	fmt.Println("Power",power)

	for index := 0; index < 10000; index++ {
		val := getNextString(scanner)
		res := getNextString(scanner)

		//fmt.Println(val, res)
		firstInt, _ := strconv.Atoi(string(val[0]))
		firstCar := string(res[0])

		for _, c := range res {
			_, ok := maxPossible[string(c)]
			if !ok {
				carNotUse[string(c)] = 0
			}
		}

		delete(carNotUse, firstCar)
		max, ok := maxPossible[firstCar]
		if !ok {
			maxPossible[firstCar] = 9
			max = 9
		}

		if len(val) == 1 {
			if max > firstInt {
				maxPossible[firstCar] = firstInt
			}
		} else if val != "-1" && len(res) == power {
			if max > firstInt {
				maxPossible[firstCar] = firstInt
			}
		}

	}

	byNumber := make(map[int]string)
	for k, v := range maxPossible {
		byNumber[v] = k
	}

	lostCar := ""
	for k := range carNotUse {
		lostCar = k
	}

	solution := ""
	for i := 0; i < 10; i++ {
		v, ok := byNumber[i]
		if !ok {
			v = lostCar
		}
		solution += v
	}
	return solution, nil
}

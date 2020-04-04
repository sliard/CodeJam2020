package main

import (
	"fmt"
	"os"
)

func main() {
	var nbCase int
	var b int
	if _, err := fmt.Scanln(&nbCase, &b); err != nil {
		panic(err)
	}

	for caseId := 1; caseId <= nbCase; caseId++ {
		fmt.Fprintf(os.Stderr, "@@@@@@@@@@@@@@@@@@@@ CASE : %d\n\r", caseId)
		if b == 10 {
			resolvTen()
		} else {
			resolv(b)
		}
	}
}

type solution []int
type allPossible []solution

func resolv(b int) {

	oneResult := solution(make([]int, b))

	testSize := 4

	//fmt.Fprintf(os.Stderr, "****** of foo: %d", len(result))
	first := true
	delta := 0
	index := 1
	nbRead := 0

	nbRetry := 0

	for index <= b/2 {
		//fmt.Fprintf(os.Stderr, "Read %d - %d\n\r", nbRead, nbRead % 10)

		if nbRead%10 == 0 && !first {

			testSize = 4

			newPart := make([]int, testSize*2)
			for n := 0; n < testSize; n++ {
				newPart[n] = readValue(n + 1 + delta)
				newPart[testSize*2-1-n] = readValue(b - n - delta)
			}
			nbRead += testSize * 2
			fmt.Fprintf(os.Stderr, "Result : %v\n\r", oneResult)
			fmt.Fprintf(os.Stderr, "Find changes : %v\n\r", newPart)
			newSolutions := findAndChange(&oneResult, &newPart, testSize, delta)

			var newSolutions2 []solution

			if len(newSolutions) > 1 {
				newPart2 := make([]int, 10)
				newPart2[0] = newPart[0]
				newPart2[1] = newPart[1]
				newPart2[2] = newPart[2]
				newPart2[3] = newPart[3]
				newPart2[4] = readValue(5 + delta)
				newPart2[5] = readValue(b - 4 - delta)
				newPart2[6] = newPart[4]
				newPart2[7] = newPart[5]
				newPart2[8] = newPart[6]
				newPart2[9] = newPart[7]

				fmt.Fprintf(os.Stderr, "Result**2 : %v\n\r", oneResult)
				fmt.Fprintf(os.Stderr, "Find changes**2 : %v\n\r", newPart2)
				newSolutions2 = findAndChange(&oneResult, &newPart2, 5, delta)

				nbRead += 2
			}

			if len(newSolutions2) > 0 {
				oneResult = newSolutions[0]
			} else if len(newSolutions) > 0 {
				oneResult = newSolutions[0]
			} else {
				//				panic(":(")
			}

		}

		if nbRead%10 != 0 || first {
			oneResult[index-1] = readValue(index)
			oneResult[b-index] = readValue(b - index + 1)
			index++
			nbRead += 2
		} else {
			nbRetry++
			delta++
		}
		first = false
	}

	if testSolution(&oneResult) {
		fmt.Fprintf(os.Stderr, "OK : \n\r")
		return
	}
	panic("No solution")
}

func testSolution(s *solution) bool {

	finalResult := make([]rune, len(*s))
	for i := 0; i < len(*s); i++ {
		if (*s)[i] == 1 {
			finalResult[i] = '1'
		} else {
			finalResult[i] = '0'
		}
	}
	fmt.Fprintf(os.Stderr, "**Test : %s", string(finalResult))
	var val string
	fmt.Printf("%s\n", string(finalResult))
	if _, err := fmt.Scan(&val); err != nil {
		fmt.Fprintf(os.Stderr, "****** of foo: %d", 1)
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "= --%s--\n\r", val)
	if val == "N" {
		fmt.Fprintf(os.Stderr, "Return false\n\r")
		return false
	}
	fmt.Fprintf(os.Stderr, "Return true\n\r")
	return true
}

func setData(all *allPossible, index, value int) {
	for i := 0; i < len(*all); i++ {
		(*all)[i][index] = value
	}
}

func findAndChange(allData *solution, part *[]int, testSize, delta int) []solution {
	size := len(*allData)

	var result []solution

	partData := make([]int, testSize*2)
	for n := 0; n < testSize; n++ {
		partData[n] = (*allData)[n+delta]
		partData[testSize*2-1-n] = (*allData)[size-1-delta-n]
	}

	if compareArray(&partData, part) {
		fmt.Fprintf(os.Stderr, "> NADA !\n\r")
		result = append(result, Same(*allData))
	}

	revertPar := Reverse(*part)
	if compareArray(&partData, &revertPar) {
		fmt.Fprintf(os.Stderr, "> Reverse !\n\r")

		oneSolution := Reverse(*allData)
		exit := false
		for _, s := range result {
			if compareArraySolution(&s, &oneSolution) {
				exit = true
				break
			}
		}
		if !exit {
			result = append(result, oneSolution)
		}
	}
	flipPart := BitFlip(*part)
	if compareArray(&partData, &flipPart) {
		fmt.Fprintf(os.Stderr, "> BitFlip !\n\r")

		oneSolution := BitFlip(*allData)

		exit := false
		for _, s := range result {
			if compareArraySolution(&s, &oneSolution) {
				exit = true
				break
			}
		}

		if !exit {
			result = append(result, oneSolution)
		}
	}
	revertFlipPart := BitFlip(revertPar)
	if compareArray(&partData, &revertFlipPart) {
		fmt.Fprintf(os.Stderr, "> Reverse + BitFlip !\n\r")

		oneSolution := Reverse(BitFlip(*allData))

		exit := false
		for _, s := range result {
			if compareArraySolution(&s, &oneSolution) {
				exit = true
				break
			}
		}
		if !exit {
			result = append(result, oneSolution)
		}
	}

	return result
}

func compareArray(a, b *[]int) bool {
	if len(*a) != len(*b) {
		return false
	}
	for i := 0; i < len(*a); i++ {
		if (*a)[i] != (*b)[i] {
			return false
		}
	}
	return true
}

func compareArraySolution(a *solution, b *[]int) bool {
	if len(*a) != len(*b) {
		return false
	}
	for i := 0; i < len(*a); i++ {
		if (*a)[i] != (*b)[i] {
			return false
		}
	}
	return true
}

func readValue(index int) (val int) {
	fmt.Printf("%d\n", index)
	if _, err := fmt.Scan(&val); err != nil {
		fmt.Fprintf(os.Stderr, "****** of foo: %d\n\r", index)
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "READVAL %d : %d\n\r", index, val)
	return
}

func Same(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	return result
}

func Reverse(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	for i := 0; i < len(result)/2; i++ {
		result[i], result[len(result)-i-1] = result[len(result)-i-1], result[i]
	}
	return result
}
func BitFlip(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	for i := 0; i < len(result); i++ {
		if result[i] == 1 {
			result[i] = 0
		} else {
			result[i] = 1
		}
	}
	return result
}

func resolvTen() {
	var val string
	var result []rune
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d\n", i)
		if _, err := fmt.Scan(&val); err != nil {
			fmt.Fprintf(os.Stderr, "****** of foo: %d", 1)
			panic(err)
		}
		if val == "1" {
			result = append(result, '1')
		} else {
			result = append(result, '0')
		}
	}

	fmt.Printf("%s\n", string(result))
	if _, err := fmt.Scan(&val); err != nil {
		fmt.Fprintf(os.Stderr, "****** of foo: %d", 1)
		panic(err)
	}
	fmt.Fprintf(os.Stderr, "Result ? : %s \n\r", val)
}

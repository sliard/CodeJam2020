package main

import "fmt"

func main() {

	var nbCase int64
	if _, err := fmt.Scanln(&nbCase); err != nil {
		panic(err)
	}
	for caseId := int64(1); caseId <= nbCase; caseId++ {
		var squareSize int64
		if _, err := fmt.Scanln(&squareSize); err != nil {
			panic(err)
		}

		k := int64(0)
		kIndex := int64(0)

		c := int64(0)
		r := int64(0)

		var allRaw []map[int64]rune
		for i := int64(0); i < squareSize; i++ {
			allRaw = append(allRaw, make(map[int64]rune))
		}

		for i := int64(0); i < squareSize; i++ {
			line := make(map[int64]rune)
			for j := int64(0); j < squareSize; j++ {
				var val int64
				if _, err := fmt.Scan(&val); err != nil {
					panic(err)
				}
				line[val] = ' '
				allRaw[j][val] = ' '
				if j == kIndex {
					k += val
				}
			}
			if int64(len(line)) != squareSize {
				r++
			}
			kIndex++
		}
		for i := int64(0); i < squareSize; i++ {
			if int64(len(allRaw[i])) != squareSize {
				c++
			}
		}
		fmt.Printf("Case #%d: %d %d %d \n", caseId, k, r, c)
	}
}

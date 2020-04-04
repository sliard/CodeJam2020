package main

import "fmt"

func main() {


	var nbCase int64
	if _, err := fmt.Scanln(&nbCase); err != nil {
		panic(err)
	}
	for caseId := int64(1); caseId <= nbCase; caseId++ {
		var inputString string
		if _, err := fmt.Scanln(&inputString); err != nil {
			panic(err)
		}
		var result []rune

		nbOpen := 0
		charBefore := '0'
		for _, char := range inputString {
			delta := charBefore - char
			if delta < 0 {
				for i:=int32(0); i<(delta*-1); i++ {
					nbOpen++
					result = append(result, '(')
				}
			} else if delta > 0 {
				for i:=int32(0); i<delta; i++ {
					nbOpen--
					result = append(result, ')')
				}
			}
			result = append(result, char)
			charBefore = char
		}

		if nbOpen > 0 {
			for i:=0; i<nbOpen; i++ {
				result = append(result, ')')
			}
		}

		fmt.Printf("Case #%d: %s\n",caseId,string(result))

	}
}
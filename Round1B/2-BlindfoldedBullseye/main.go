package main

import (
	"fmt"
	"os"
)

func main() {

	var nbCase int
	var a int
	var b int
	if _, err := fmt.Scanln(&nbCase, &a, &b); err != nil {
		panic(err)
	}
	for caseId := 1; caseId <= nbCase; caseId++ {
		resolv(a, b)
	}
}

var MIN = -1000000000
var MAX = 1000000000

type point struct {
	x int
	y int
}

func readString() string {
	var response string
	if _, err := fmt.Scanln(&response); err != nil {
		panic(err)
	}
	return response
}

func resolv(a, b int) {

	// find first point

	var firstPoint *point

	for x := MIN + a; x <= MAX && firstPoint == nil; x += a {
		for y := MIN + a; y <= MAX && firstPoint == nil; y += a {
			fmt.Printf("%d %d\n", x, y)

			response := readString()
			fmt.Fprintf(os.Stderr, "Test: %d %d = %s\n", x, y, response)

			if response == "CENTER" {
				return
			}
			if response == "HIT" {
				firstPoint = &point{x, y}
			}
		}
	}

	//fin y limit
	// for top
	minY := (*firstPoint).y
	maxY := (*firstPoint).y + 2*b
	found := false
	if maxY > MAX {
		maxY = MAX
		fmt.Printf("%d %d\n", (*firstPoint).x, maxY)
		response := readString()
		if response == "CENTER" {
			return
		}
		if response == "HIT" {
			found = true
			minY = MAX
		}
	}
	mid := 0
	for minY+1 != maxY && !found {
		mid = (maxY + minY) / 2
		fmt.Printf("%d %d\n", (*firstPoint).x, mid)
		response := readString()
		if response == "CENTER" {
			return
		}
		if response == "HIT" {
			minY = mid
		} else {
			maxY = mid
		}
	}
	topY := minY
	fmt.Fprintf(os.Stderr, "Limit TOP Y !!: %d \n", topY)

	found = false
	minY = (*firstPoint).y - 2*b
	maxY = (*firstPoint).y
	if minY < MIN {
		minY = MIN
		fmt.Printf("%d %d\n", (*firstPoint).x, minY)
		response := readString()
		if response == "CENTER" {
			return
		}
		if response == "HIT" {
			found = true
			maxY = MIN
		}
	}
	mid = 0
	for minY+1 != maxY && !found {
		mid = (maxY + minY) / 2
		fmt.Printf("%d %d\n", (*firstPoint).x, mid)
		response := readString()
		if response == "CENTER" {
			return
		}
		if response == "HIT" {
			maxY = mid
		} else {
			minY = mid
		}
	}
	butY := maxY
	fmt.Fprintf(os.Stderr, "Limit BUTTOM Y !!: %d \n", butY)

	found = false
	minX := (*firstPoint).x
	maxX := (*firstPoint).x + 2*b + 1
	if maxX > MAX {
		maxX = MAX
		fmt.Printf("%d %d\n", maxX, (*firstPoint).y)
		response := readString()
		if response == "CENTER" {
			return
		}
		if response == "HIT" {
			found = true
			minX = MAX
		}
	}
	mid = 0
	for minX+1 != maxX && !found {
		mid = (maxX + minX) / 2
		fmt.Printf("%d %d\n", mid, (*firstPoint).y)
		response := readString()
		if response == "CENTER" {
			return
		}
		if response == "HIT" {
			minX = mid
		} else {
			maxX = mid
		}
	}
	rightX := minX
	fmt.Fprintf(os.Stderr, "Limit Right X !!: %d \n", rightX)

	found = false
	minX = (*firstPoint).x - 2*b
	maxX = (*firstPoint).x
	if minX < MIN {
		minX = MIN
		fmt.Printf("%d %d\n", minX, (*firstPoint).y)
		response := readString()
		if response == "CENTER" {
			return
		}
		if response == "HIT" {
			found = true
			maxX = MIN
		}
	}
	mid = 0
	for minX+1 != maxX && !found {
		//		fmt.Fprintf(os.Stderr, "Loop: %d %d \n", minY, maxY)
		mid = (maxX + minX) / 2
		fmt.Printf("%d %d\n", mid, (*firstPoint).y)
		response := readString()
		if response == "CENTER" {
			return
		}
		if response == "HIT" {
			maxX = mid
		} else {
			minX = mid
		}
	}
	leftX := maxX
	fmt.Fprintf(os.Stderr, "Limit Left X !!: %d \n", leftX)

	fmt.Printf("%d %d\n", (rightX+leftX)/2, (butY+topY)/2)
	response := readString()
	fmt.Fprintf(os.Stderr, "response !!: %d %d = %s \n", (rightX+leftX)/2, (butY+topY)/2, response)

}

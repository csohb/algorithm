package main

import (
	"fmt"
	"strings"
)

const (
	BlockA  = "AAAA"
	BlockB  = "BB"
	BlockAB = "AAAABB"
)

func polyomino(board string) string {
	boardList := strings.Split(board, ".")
	var convertedString string
	for i, v := range boardList {
		if len(v)%6 == 0 {
			boardList[i] = BlockAB
		} else if len(v)%4 == 0 {
			boardList[i] = BlockA
		} else if len(v)%2 == 0 {
			boardList[i] = BlockB
		} else {
			return "Nope"
		}
	}
	fmt.Println("boardList:", boardList)
	for _, v := range boardList {
		convertedString += v
	}
	fmt.Println("convertedString:", convertedString)

	return convertedString
}

func main() {
	//board := "XXXXXX"
	//board := "XX.XX"
	board := "XX.XXXXXXXXXX..XXXXXXXX...XXXXXX"
	output := polyomino(board)
	fmt.Println("output:", output)
}

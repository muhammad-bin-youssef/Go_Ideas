package main

import (
	"fmt"
	"os"
)

func Loop(arr *[]byte, lineCounter *uint, start int, end int) {
	fmt.Printf("%08s", fmt.Sprintf("%b", *lineCounter))
	(*lineCounter)++
	for _, i := range (*arr)[start:end] {
		fmt.Printf("%4s ", fmt.Sprintf("%X", i))
	}
	fmt.Println()
	fmt.Print("        ")
	for _, i := range (*arr)[start:end] {
		fmt.Printf("%5s", fmt.Sprintf("%q", i))
	}
	fmt.Println()
}

func main() {
	file, err := os.ReadFile(os.Args[1])
	lineCounter := uint(0)
	lineLength := 16
	start := 0
	end := lineLength

	if err != nil {
		panic(err)
	}
	if len(file) < lineLength {
		Loop(&file, &lineCounter, int(0), int(len(file)))
		os.Exit(1)
	}
	for len(file) > 0 {
		if end > len(file) {
			Loop(&file, &lineCounter, start, len(file))
			os.Exit(1)
		}
		if len(file) > lineLength {
			Loop(&file, &lineCounter, start, end)
			start = end
			end = end + lineLength
		}
	}
}

package main

import (
	"fmt"
	"os"
)

func Loop(arr *[]byte, lineCounter *int64, start int64, end int64) {
	fmt.Printf("%08s", fmt.Sprintf("%b", *lineCounter))
	(*lineCounter) += 16
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
	if len(os.Args) < 2 {
		panic("no argument were provided. You need to pass one file to read.")
	}
	file, err := os.ReadFile(os.Args[1])
	lineCounter := int64(0)
	lineLength := int64(16)
	start := int64(0)
	end := lineLength
	filelen := int64(len(file))

	if err != nil {
		panic(err)
	}
	if filelen < lineLength {
		Loop(&file, &lineCounter, int64(0), filelen)
		os.Exit(1)
	}
	for filelen > 0 {
		if end > filelen {
			Loop(&file, &lineCounter, start, filelen)
			os.Exit(1)
		}
		if filelen > lineLength {
			Loop(&file, &lineCounter, start, end)
			start = end
			end = end + lineLength
		}
	}
}

package main

import (
	"fmt"
	"linear-stats/calculations"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please provide one txt file.")
		return
	}

	if !strings.HasSuffix(args[0], ".txt") {
		fmt.Println("Please provide txt files.")
		return
	}
	data := ReadFile(args[0])

	sliceX, sliceY := calculations.StingToSliceInt(data)
	a, b := calculations.CalculateLinearCoeff(sliceX, sliceY)
	c := calculations.CalculatePearsonCoeff(sliceX, sliceY)

	fmt.Printf(`Linear Regression Line: y = %.6fx + %.6f%s`, a, b, "\n")
	fmt.Printf(`Pearson Correlation Coefficient: %.10f%s`, c, "\n")
}

func ReadFile(sampleFile string) string {
	data, err := os.ReadFile(sampleFile)
	if err != nil {
		log.Print(err)
	}
	dataStr := string(data)
	return dataStr
}

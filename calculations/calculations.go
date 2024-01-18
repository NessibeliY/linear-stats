package calculations

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func StingToSliceInt(s string) (sliceX, sliceY []int) {
	lines := strings.Split(s, "\n")

	x := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		y, err := strconv.Atoi(line)
		if err != nil {
			log.Println(err)
			return nil, nil
		}
		sliceX = append(sliceX, x)
		sliceY = append(sliceY, y)
		x++
	}
	return sliceX, sliceY
}

func CalculateAuxSums(sliceX, sliceY []int) (n, sumX, sumY, sumXY, sumX2, sumY2 float64) {
	sumX = float64(Sum(sliceX))
	sumY = float64(Sum(sliceY))
	n = float64(len(sliceX))
	for i := 0; i < len(sliceX); i++ {
		sumX2 += float64(sliceX[i] * sliceX[i])
		sumY2 += float64(sliceY[i] * sliceY[i])
		sumXY += float64(sliceX[i] * sliceY[i])
	}

	return n, sumX, sumY, sumXY, sumX2, sumY2
}

func CalculateLinearCoeff(sliceX, sliceY []int) (a, b float64) {
	n, sumX, sumY, sumXY, sumX2, _ := CalculateAuxSums(sliceX, sliceY)
	a = (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	b = (sumY*sumX2 - sumX*sumXY) / (n*sumX2 - sumX*sumX)
	return a, b
}

func CalculatePearsonCoeff(sliceX, sliceY []int) (c float64) {
	n, sumX, sumY, sumXY, sumX2, sumY2 := CalculateAuxSums(sliceX, sliceY)
	c = (n*sumXY - sumX*sumY) / math.Sqrt((n*sumX2-sumX*sumX)*(n*sumY2-sumY*sumY))
	return c
}

func Sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func CalculateAverage(slice []int) float64 {
	var average float64
	for _, num := range slice {
		average += float64(num)
	}
	lengthOfSlice := float64(len(slice))
	average = math.Round(average / lengthOfSlice)
	return average
}

func CalculateVariance(slice []int) float64 {
	var variance float64
	average := CalculateAverage(slice)
	for _, num := range slice {
		variance += (float64(num) - average) * (float64(num) - average)
	}
	lengthOfSlice := float64(len(slice))
	variance = variance / (lengthOfSlice)
	return variance
}

func CalculateStdDev(slice []int) float64 {
	var stdDev float64
	variance := CalculateVariance(slice)
	stdDev = math.Sqrt(variance)
	return stdDev
}

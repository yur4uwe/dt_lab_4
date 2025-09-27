package randanalysis

import (
	"fmt"
	"math"
	"slices"
)

func Mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	sum := 0.0
	for _, v := range data {
		sum += v
	}

	return sum / float64(len(data))
}

func Mode(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	freq := make(map[float64]int)

	for _, v := range data {

		freq[v]++
	}

	mode := data[0]
	maxCount := 0
	for k, v := range freq {
		if v > maxCount {
			maxCount = v
			mode = k
		}
	}

	return mode
}

func Median(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	slices.Sort(data)

	mid := len(data) / 2

	if len(data)%2 == 0 {
		return (data[mid-1] + data[mid]) / 2
	}
	return data[mid]
}

func RMS(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	sum := 0.0
	for _, v := range data {
		sum += v * v
	}

	return math.Sqrt(sum / float64(len(data)))
}

func GeometricMean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	product := 1.0
	for _, v := range data {
		if v <= 0 {
			return 0
		}
		product *= v
	}

	fmt.Println("Product:", product)

	return math.Pow(product, 1.0/float64(len(data)))
}

func HarmonicMean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}

	sum := 0.0
	for _, v := range data {
		if v == 0 {
			return 0
		}
		sum += 1 / v
	}

	return float64(len(data)) / sum
}

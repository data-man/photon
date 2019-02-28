package photon

import (
	"math"
)

func min(params ...int) int {
	minVal := math.MaxInt32
	for _, val := range params {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func max(params ...int) int {
	maxVal := math.MinInt32
	for _, val := range params {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

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

func calculateSlope(p1, p2 Point) float32 {
	return float32(p2.Y-p1.Y) / float32(p2.X-p1.X)
}

func isPointOnLine(p, lineStart, lineEnd Point) bool {
	slope := calculateSlope(lineStart, lineEnd)
	b := float32(lineStart.Y) - (slope * float32(lineStart.X))
	return float32(p.Y) == (slope*float32(p.X))+b
}

func isPointInsideCircle(p, centerPoint Point, radius int) bool {
	return math.Pow(float64(p.X-centerPoint.X), 2)+math.Pow(float64(p.Y-centerPoint.Y), 2) < math.Pow(float64(radius), 2)
}

func distance(p1, p2 Point) float32 {
	return float32(math.Sqrt(float64(p1.X-p2.X) + float64(p1.Y-p2.Y)))
}

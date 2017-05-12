package ava

import "math"

// PooledVarianceTStat calculates the t statistic to be used in the comparison of two means.
func PooledVarianceTStat(groups [][]float64) float64 {
	x1 := average(groups[0])
	n1 := float64(len(groups[0]))
	x2 := average(groups[1])
	n2 := float64(len(groups[1]))
	s := pooledVariance(groups)
	denom := math.Sqrt(s * ((1 / n1) + (1 / n2)))
	return (x1 - x2) / denom
}

package ava

import (
	"math"

	"github.com/ThePaw/probab/dst"
)

// TResult is returned after completing a T test. Stat is the computed t-statistic, and PValue is the corresponding probability.
type TResult struct {
	Stat   float64
	PValue float64
}

// PooledVarianceTTest performs a pooled variance t-test on the provided data.
func PooledVarianceTTest(groups [][]float64) TResult {
	t := pooledVarianceTStat(groups)
	p := probabilityOfTStat(t, groups)
	return TResult{
		Stat:   t,
		PValue: p,
	}
}

// pooledVarianceTStat calculates the t statistic to be used in the comparison of two means.
func pooledVarianceTStat(groups [][]float64) float64 {
	x1 := average(groups[0])
	n1 := float64(len(groups[0]))
	x2 := average(groups[1])
	n2 := float64(len(groups[1]))
	s := pooledVariance(groups)
	denom := math.Sqrt(s * ((1 / n1) + (1 / n2)))
	return (x1 - x2) / denom
}

func probabilityOfTStat(t float64, groups [][]float64) float64 {
	df := float64(len(groups[0]) + len(groups[1]) - 2)
	return (1 - dst.StudentsTCDFAt(df, math.Abs(t))) * 2
}

package ava

import "github.com/ThePaw/probab/dst"

type AnovaResult struct {
	Stat   float64
	PValue float64
}

func OneWayAnova(groups [][]float64) AnovaResult {
	f := fStat(groups)
	p := probabilityOfFStat(f, groups)
	return AnovaResult{
		Stat:   f,
		PValue: p,
	}
}

func (a AnovaResult) Significant(alpha float64) bool {
	return a.PValue < alpha
}

func ssa(groups [][]float64) float64 {
	var ssa float64
	grandAvg := average(groups)
	for _, j := range groups {
		n := float64(len(j))
		xbar := average(j)
		ssa += n * ((xbar - grandAvg) * (xbar - grandAvg))
	}
	return ssa
}

func ssw(groups [][]float64) float64 {
	var ssw float64
	for _, j := range groups {
		xbar := average(j)
		for _, xi := range j {
			ssw += ((xi - xbar) * (xi - xbar))
		}
	}
	return ssw
}

func sst(groups [][]float64) float64 {
	var sst float64
	grandAvg := average(groups)
	for _, j := range groups {
		for _, xi := range j {
			sst += ((xi - grandAvg) * (xi - grandAvg))
		}
	}
	return sst
}

func msa(groups [][]float64) float64 {
	return ssa(groups) / float64(len(groups)-1)
}

func msw(groups [][]float64) float64 {
	return ssw(groups) / float64(nestedLen(groups)-len(groups))
}

func fStat(groups [][]float64) float64 {
	return msa(groups) / msw(groups)
}

func probabilityOfFStat(f float64, groups [][]float64) float64 {
	n := len(groups)
	c := nestedLen(groups) - n
	return 1 - dst.FCDFAt(int64(n-1), int64(c), f)
}

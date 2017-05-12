package ava

import "math"

// TukeyTest ...
type TukeyTest struct {
	MSW    float64
	Groups int
	Df     int
}

// TukeyHSD ...
func TukeyHSD(data [][]float64) TukeyTest {
	msw := MSW(data)
	return TukeyTest{
		MSW:    msw,
		Groups: len(data),
		Df:     nestedLen(data) - len(data),
	}
}

// Compare ...
func (t *TukeyTest) Compare(g1 []float64, g2 []float64) bool {
	q := lookupQValue(t.Groups, t.Df)
	criticalRange := tukeyCriticalRange(q, t.MSW, len(g1), len(g2))
	return math.Abs(average(g1)-average(g2)) > criticalRange
}

func tukeyCriticalRange(q float64, msw float64, n1 int, n2 int) float64 {
	return q * math.Sqrt((msw/2)*((1/float64(n1))+(1/float64(n2))))
}

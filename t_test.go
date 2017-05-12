package ava

import "testing"

func TestPooledVarianceTStat(t *testing.T) {
	data := [][]float64{
		[]float64{22, 34, 52, 62, 30, 40, 64, 84, 56, 59},
		[]float64{52, 71, 76, 54, 67, 83, 66, 90, 77, 84},
	}
	actual := PooledVarianceTStat(data)
	expected := -3.044550122546798

	if actual != expected {
		t.Errorf("PooledVarianceTStat: Expected %+v, got %+v", expected, actual)
	}
}

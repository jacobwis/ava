package ava

import "testing"

func TestPooledVarianceTTest(t *testing.T) {
	data := [][]float64{
		[]float64{22, 34, 52, 62, 30, 40, 64, 84, 56, 59},
		[]float64{52, 71, 76, 54, 67, 83, 66, 90, 77, 84},
	}
	actual := PooledVarianceTTest(data)
	expected := TResult{
		Stat:   -3.044550122546798,
		PValue: 0.006974856613970015,
	}

	if actual != expected {
		t.Errorf("PooledVarianceTTest: Expected %+v, got %+v", expected, actual)
	}
}

func TestPooledVarianceTStat(t *testing.T) {
	data := [][]float64{
		[]float64{22, 34, 52, 62, 30, 40, 64, 84, 56, 59},
		[]float64{52, 71, 76, 54, 67, 83, 66, 90, 77, 84},
	}
	actual := pooledVarianceTStat(data)
	expected := -3.044550122546798

	if actual != expected {
		t.Errorf("PooledVarianceTStat: Expected %+v, got %+v", expected, actual)
	}
}

func TestProbabilityofTStat(t *testing.T) {
	data := [][]float64{
		[]float64{22, 34, 52, 62, 30, 40, 64, 84, 56, 59},
		[]float64{52, 71, 76, 54, 67, 83, 66, 90, 77, 84},
	}
	stat := -3.044550122546798
	expected := probabilityOfTStat(stat, data)
	actual := 0.006974856613970015
	if actual != expected {
		t.Errorf("probabilityofTStat: Expected %+v, got %+v", expected, actual)
	}
}

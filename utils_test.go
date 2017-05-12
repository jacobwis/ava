package ava

import "testing"

func TestAverageProperlyHandlesSlice(t *testing.T) {
	data := []float64{18.5, 24.0, 17.2, 19.9, 18.0}
	actual := average(data)
	expected := 19.52

	if actual != expected {
		t.Errorf("Average Properly Handles Slice: Expected %v, got %v", expected, actual)
	}
}

func TestAverageProperlyHandlesNestedSlice(t *testing.T) {
	data := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := average(data)
	expected := 21.945

	if actual != expected {
		t.Errorf("Average Properly Handles Nested Slice: Expected %v, got %v", expected, actual)
	}
}

func TestAverageProperlyHandlesIncorrectType(t *testing.T) {
	actual := average("ok")
	expected := 0.0
	if actual != expected {
		t.Errorf("Average Properly Handles Nested Slice: Expected %v, got %v", expected, actual)
	}
}

func TestAverageOfFloats(t *testing.T) {
	ds := []float64{18.5, 24.0, 17.2, 19.9, 18.0}
	actualResult := averageOfFloats(ds)
	expectedResult := 19.52

	if actualResult != expectedResult {
		t.Errorf("Average of Floats: Expected %v, got %v", expectedResult, actualResult)
	}
}

func TestAverageOfNestedFloats(t *testing.T) {
	ds := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := averageofNestedFloats(ds)
	expected := 21.945

	if expected != actual {
		t.Errorf("averageOfNestedFloats: Expected %v, got %v", expected, actual)
	}
}

func TestSum(t *testing.T) {
	ds := []float64{26.3, 25.3, 24.0, 21.2, 24.5}
	actualResult := sum(ds)
	expectedResult := 121.3

	if actualResult != expectedResult {
		t.Errorf("Sum: Expected %v, got %v", expectedResult, actualResult)
	}
}

func TestNestedLen(t *testing.T) {
	ds := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actualResult := nestedLen(ds)
	expectedResult := 20
	if actualResult != expectedResult {
		t.Errorf("nestedLen: Expected %v, got %v", expectedResult, actualResult)
	}
}

func TestVariance(t *testing.T) {
	sample := []float64{39, 29, 43, 52, 39, 44, 40, 31, 44, 35}
	actual := variance(sample)
	expected := 45.82222222222223

	if expected != actual {
		t.Errorf("variance: Expected %v, got %v", expected, actual)
	}
}

func TestPooledVariance(t *testing.T) {
	data := [][]float64{
		[]float64{22, 34, 52, 62, 30, 40, 64, 84, 56, 59},
		[]float64{52, 71, 76, 54, 67, 83, 66, 90, 77, 84},
	}
	actual := pooledVariance(data)
	expected := 254.0055555555556

	if expected != actual {
		t.Errorf("pooledVariance: Expected %v, got %v", expected, actual)
	}
}

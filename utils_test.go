package ava

import (
	"testing"

	"github.com/wisniewskij26/xu"
)

func TestAverageProperlyHandlesSlice(t *testing.T) {
	data := []float64{18.5, 24.0, 17.2, 19.9, 18.0}
	actual := average(data)
	expected := 19.52

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
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

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAverageProperlyHandlesIncorrectType(t *testing.T) {
	actual := average("ok")
	expected := 0.0
	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAverageOfFloats(t *testing.T) {
	ds := []float64{18.5, 24.0, 17.2, 19.9, 18.0}
	actual := averageOfFloats(ds)
	expected := 19.52

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
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

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSum(t *testing.T) {
	ds := []float64{26.3, 25.3, 24.0, 21.2, 24.5}
	actual := sum(ds)
	expected := 121.3

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestNestedLen(t *testing.T) {
	ds := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := nestedLen(ds)
	expected := 20
	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestVariance(t *testing.T) {
	sample := []float64{39, 29, 43, 52, 39, 44, 40, 31, 44, 35}
	actual := variance(sample)
	expected := 45.82222222222223

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestPooledVariance(t *testing.T) {
	data := [][]float64{
		[]float64{22, 34, 52, 62, 30, 40, 64, 84, 56, 59},
		[]float64{52, 71, 76, 54, 67, 83, 66, 90, 77, 84},
	}
	actual := pooledVariance(data)
	expected := 254.0055555555556

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

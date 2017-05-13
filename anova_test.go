package ava

import (
	"testing"

	"github.com/wisniewskij26/xu"
)

func TestOneWayAnova(t *testing.T) {
	data := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := OneWayAnova(data)
	expected := AnovaResult{
		Stat:   3.4616289246253085,
		PValue: 0.041365599903377515,
	}
	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestAnovaSignificant(t *testing.T) {
	res := AnovaResult{
		Stat:   3.4616289246253085,
		PValue: 0.041365599903377515,
	}
	actual := res.Significant(0.05)
	expected := true

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSSA(t *testing.T) {
	data := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := ssa(data)
	expected := 63.2854999999999

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSSW(t *testing.T) {
	data := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := ssw(data)
	expected := 97.5040

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestSST(t *testing.T) {
	data := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := sst(data)
	expected := 160.7895

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestMSA(t *testing.T) {
	data := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := MSA(data)
	expected := 21.095166666666632

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestMSW(t *testing.T) {
	data := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := MSW(data)
	expected := 6.0940

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestFStat(t *testing.T) {
	data := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := fStat(data)
	expected := 3.4616289246253085

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestProbabilityOfFStat(t *testing.T) {
	data := [][]float64{
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	actual := probabilityOfFStat(3.4616289246253085, data)
	expected := 0.041365599903377515

	err := xu.AssertEqual(actual, expected)
	if err != nil {
		t.Errorf(err.Error())
	}
}

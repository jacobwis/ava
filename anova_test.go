package ava

import "testing"

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

	if actual != expected {
		t.Errorf("OneWayAnova: Expected %+v, got %+v", expected, actual)
	}
}

func TestAnovaSignificant(t *testing.T) {
	res := AnovaResult{
		Stat:   3.4616289246253085,
		PValue: 0.041365599903377515,
	}
	actual := res.Significant(0.05)
	expected := true
	if actual != expected {
		t.Errorf("AnovaSignificant: Expected %v, got %v", expected, actual)
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

	if actual != expected {
		t.Errorf("SSA: Expected %v, got %v", expected, actual)
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

	if actual != expected {
		t.Errorf("SSW: Expected %v, got %v", expected, actual)
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

	if actual != expected {
		t.Errorf("SST: Expected %v, got %v", expected, actual)
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

	if actual != expected {
		t.Errorf("MSA: Expected %v, got %v", expected, actual)
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

	if actual != expected {
		t.Errorf("MSW: Expected %v, got %v", expected, actual)
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

	if actual != expected {
		t.Errorf("fStat: Expected %v, got %v", expected, actual)
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

	if actual != expected {
		t.Errorf("probabilityOfFStat: Expected %v, got %v", expected, actual)
	}
}

package ava

import "testing"

func TestTukeyHSD(t *testing.T) {
	data := [][]float64{
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}

	actual := TukeyHSD(data)
	expected := TukeyTest{
		MSW:    6.094,
		Groups: 4,
		Df:     16,
	}

	if actual != expected {
		t.Errorf("TukeyHSD: Expected %v, got %v", expected, actual)
	}
}

func TestTukeyCriticalRange(t *testing.T) {
	actual := tukeyCriticalRange(4.05, 6.094, 5, 5)
	expected := 4.471170652077596

	if actual != expected {
		t.Errorf("tukeyCriticalRange: Expected %v, got %v", expected, actual)
	}
}

func TestTukeyComparison(t *testing.T) {
	data := [][]float64{
		[]float64{18.5, 24.0, 17.2, 19.9, 18.0},
		[]float64{26.3, 25.3, 24.0, 21.2, 24.5},
		[]float64{20.6, 25.2, 20.8, 24.7, 22.9},
		[]float64{25.4, 19.9, 22.6, 17.5, 20.4},
	}
	tukey := TukeyHSD(data)
	actual := tukey.Compare(data[0], data[1])
	expected := true
	if actual != expected {
		t.Errorf("tukeyComparison: Expected %v, got %v", expected, actual)
	}
	actual = tukey.Compare(data[0], data[2])
	expected = false
	if actual != expected {
		t.Errorf("tukeyComparison: Expected %v, got %v", expected, actual)
	}
}

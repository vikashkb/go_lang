package main

import "testing"

func TestSum1(t *testing.T) {
	expected := 5
	actual := MySum(2, 3)

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSum2(t *testing.T) {
	expected := 36
	actual := MySum(19, 17)

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSub1(t *testing.T) {
	expected := 1
	actual := MySub(3, 2)

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSub2(t *testing.T) {
	expected := -1
	actual := MySub(2, 3)

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestMul1(t *testing.T) {
	expected := 6
	actual := MyMul(2, 3)

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestMul2(t *testing.T) {
	expected := -6
	actual := MyMul(-2, 3)

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestDiv1(t *testing.T) {
	expected := -3
	actual, err := MyDiv(-6, 2)

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
		t.Error(err)
	}
}

func TestDiv2(t *testing.T) {
	expected := 3
	actual, err := MyDiv(-6, -2)

	if expected != actual {
		t.Errorf("Expected %d, got %d", expected, actual)
		t.Error(err)
	}
}

func TestDiv3(t *testing.T) {
	res, err := MyDiv(0, 0)

	if err == nil {
		t.Error(res, err)
	}
}

package main

import (
	"testing"
)

func TestAdd1(t *testing.T) {
	if Add(1, 3) != 4 {
		t.Errorf("Errors in case 1 : Expected 4, Actual %d", Add(1, 3))
	}
	if Add(1, 0) != 1 {
		t.Errorf("Errors in case 2 : Expected 1, Actual %d", Add(1, 0))
	}
	if Add(1, -1) != 0 {
		t.Errorf("Errors in case 3 : Expected 0, Actual %d", Add(1, -1))
	}
}

func TestAdd2(t *testing.T) {
	set := [][]int{{1, 1, 3, 4}, {2, 1, 0, 1}, {3, 1, -1, 0}}
	for _, v := range set {
		if Add(v[1], v[2]) != v[3] {
			t.Errorf("Errors in case %d : Expected %d, Actual %d", v[0], v[3], Add(v[1], v[2]))
		}
	}
}

func TestAdd3(t *testing.T) {
	set := [][]int{{1, 3, 4}, {1, 0, 1}, {1, -1, 0}}
	for i, v := range set {
		if Add(v[0], v[1]) != v[2] {
			t.Errorf("Errors in case %d : Expected %d, Actual %d", i+1, v[2], Add(v[0], v[1]))
		}
	}
}

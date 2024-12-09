package day8

import (
	"reflect"
	"testing"
)

func verticalTest(t *testing.T) {
	p1 := point{4, 4}
	p2 := point{4, 6}

	got := getAntis(p1, p2)
	want := [2]point{{4, 8}, {4, 2}}

	if !reflect.DeepEqual(got, want) {
		t.Fail()
	}
	t.Logf("\nVertical Test: Got: %v | Wanted: %v\n", got, want)
}

func horizontalTest(t *testing.T) {
	p1 := point{4, 5}
	p2 := point{6, 5}

	got := getAntis(p1, p2)
	want := [2]point{{8, 5}, {2, 5}}

	if !reflect.DeepEqual(got, want) {
		t.Fail()
	}
	t.Logf("\nHorizontal Test: Got: %v | Wanted: %v\n", got, want)
}

func diagonalTest(t *testing.T) {
	p1 := point{4, 4}
	p2 := point{6, 5}

	got := getAntis(p1, p2)
	want := [2]point{{6, 8}, {3, 2}}

	if !reflect.DeepEqual(got, want) {
		t.Fail()
	}
	t.Logf("\nDiagonal Test: Got: %v | Wanted: %v\n", got, want)
}

func TestAntis(t *testing.T) {
	horizontalTest(t)
	verticalTest(t)
	diagonalTest(t)

}

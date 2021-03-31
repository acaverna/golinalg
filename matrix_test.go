package main

import (
	"math/rand"
	"testing"
)

func TestCreate3by3Matrix(t *testing.T) {
	t.Log("Creating 3x3 matrix")

	var (
		rows     = 3
		cols     = 3
		elements = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
		testA    = 1.0
		testB    = 5.0
		testC    = 9.0
	)

	matrix := CreateMatrix(rows, cols, elements)

	if matrix.Rows != rows && matrix.Cols != cols {
		t.Errorf("Expected matrix dimension rows: %d and cols: %d but got rows: %d and cols: %d", rows, cols, matrix.Rows, matrix.Cols)
	}

	if matrix.Get(1, 1) != testA {
		t.Errorf("Expected the value %f on, but got %f on position i=1 and j=1", testA, matrix.Get(1, 1))
	}

	if matrix.Get(2, 2) != testB {
		t.Errorf("Expected the value %f on, but got %f on position i=1 and j=1", testB, matrix.Get(2, 2))
	}

	if matrix.Get(3, 3) != testC {
		t.Errorf("Expected the value %f on, but got %f on position i=1 and j=1", testC, matrix.Get(2, 2))
	}

}

func TestCreateNbyNMatrix(t *testing.T) {

	var (
		rows = rand.Intn(1000)
		cols = rand.Intn(1000)
	)

	var elements []float64
	base := float64(rand.Intn(1000))
	size := rows * cols
	for i := 0; i < size; i++ {
		elements = append(elements, rand.Float64()*base)
	}

	matrix := CreateMatrix(rows, cols, elements)

	if matrix.Rows != rows && matrix.Cols != cols {
		t.Errorf("Expected matrix dimension rows: %d and cols: %d but got rows: %d and cols: %d", rows, cols, matrix.Rows, matrix.Cols)
	}

	i := rand.Intn(rows)
	j := rand.Intn(cols)
	element := elements[(j-1)+(i-1)*matrix.Cols]

	if matrix.Get(i, j) != element {
		t.Errorf("Expected the value %f on, but got %f on position i: %d and j: %d", element, matrix.Get(1, 1), i, j)
	}

}

func TestCreateZeroMatrix(t *testing.T) {

	var (
		rows = rand.Intn(1000)
		cols = rand.Intn(1000)
	)

	matrix := CreateZeroMatrix(rows, cols)

	i := rand.Intn(rows)
	j := rand.Intn(cols)

	element := 0.0

	if matrix.Get(i, j) != element {
		t.Errorf("Expected the value %f on, but got %f on position i: %d and j: %d", element, matrix.Get(1, 1), i, j)
	}

}

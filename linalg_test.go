package main

import (
	"math/rand"
	"testing"
)

func TestMatrixTranspose(t *testing.T) {

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

	a := CreateMatrix(rows, cols, elements)

	la := CreateLinAlg()

	c := la.Transpose(a)

	if a.Rows != c.Cols || a.Cols != c.Rows {
		t.Errorf("The dimensions of A matrix and C matrix is wrong!")
	}

	i := rand.Intn(rows)
	j := rand.Intn(cols)

	if a.Get(i, j) != c.Get(j, i) {
		t.Errorf("The ij-element of A matrix is different of ji-elemento of C matrix")
	}

}

func TestMatrixAdd(t *testing.T) {

	var (
		rows = rand.Intn(1000)
		cols = rand.Intn(1000)
	)

	base := float64(rand.Intn(1000))
	size := rows * cols

	var elementsA []float64
	for i := 0; i < size; i++ {
		elementsA = append(elementsA, rand.Float64()*base)
	}
	a := CreateMatrix(rows, cols, elementsA)

	var elementsB []float64
	for i := 0; i < size; i++ {
		elementsB = append(elementsB, rand.Float64()*base)
	}
	b := CreateMatrix(rows, cols, elementsB)

	la := CreateLinAlg()
	c := la.Add(a, b)

	if c.Rows != a.Rows || c.Cols != a.Cols {
		t.Errorf("The dimensions of A matrix and C matrix is wrong!")
	}

	i := rand.Intn(rows)
	j := rand.Intn(cols)

	value := a.Get(i, j) + b.Get(i, j)
	if c.Get(i, j) != value {
		t.Errorf("Wrong value on C(i,j): Got %f but expected %f!", c.Get(i, j), value)
	}

}

func TestMatrixHadamardProduct(t *testing.T) {

	var (
		rows = rand.Intn(1000)
		cols = rand.Intn(1000)
	)

	base := float64(rand.Intn(1000))
	size := rows * cols

	var elementsA []float64
	for i := 0; i < size; i++ {
		elementsA = append(elementsA, rand.Float64()*base)
	}
	a := CreateMatrix(rows, cols, elementsA)

	var elementsB []float64
	for i := 0; i < size; i++ {
		elementsB = append(elementsB, rand.Float64()*base)
	}
	b := CreateMatrix(rows, cols, elementsB)

	la := CreateLinAlg()
	c := la.Times(a, b)

	if c.Rows != a.Rows || c.Cols != a.Cols {
		t.Errorf("The dimensions of A matrix and C matrix is wrong!")
	}

	i := rand.Intn(rows)
	j := rand.Intn(cols)

	value := a.Get(i, j) * b.Get(i, j)
	if c.Get(i, j) != value {
		t.Errorf("Wrong value on C(i,j): Got %f but expected %f!", c.Get(i, j), value)
	}

}

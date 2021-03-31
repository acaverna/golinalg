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

func TestScalarTimesMatrix(t *testing.T) {

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

	s := rand.Float64() * base

	la := CreateLinAlg()
	c := la.STimes(s, a)

	if c.Rows != a.Rows || c.Cols != a.Cols {
		t.Errorf("The dimensions of A matrix and C matrix is wrong!")
	}

	i := rand.Intn(rows)
	j := rand.Intn(cols)

	value := s * a.Get(i, j)
	if c.Get(i, j) != value {
		t.Errorf("Wrong value on C(i,j): Got %f but expected %f!", c.Get(i, j), value)
	}

}

func TestMatrixDot(t *testing.T) {

	var (
		rows = rand.Intn(1000)
		cols = rand.Intn(1000)
		k    = rand.Intn(1000)
	)

	base := float64(rand.Intn(1000))
	sizeA := rows * k

	var elementsA []float64
	for i := 0; i < sizeA; i++ {
		elementsA = append(elementsA, rand.Float64()*base)
	}
	a := CreateMatrix(rows, k, elementsA)

	sizeB := k * cols
	var elementsB []float64
	for i := 0; i < sizeB; i++ {
		elementsB = append(elementsB, rand.Float64()*base)
	}
	b := CreateMatrix(k, cols, elementsB)

	la := CreateLinAlg()
	c := la.Dot(a, b)

	if c.Rows != a.Rows || c.Cols != b.Cols {
		t.Errorf("The dimensions of C matrix is wrong!: Got C(%d,%d), but expected C(%d,%d)", c.Rows, c.Cols, a.Rows, b.Cols)
	}

	i := rand.Intn(rows)
	j := rand.Intn(cols)

	var value float64
	for k := 1; k <= a.Cols; k++ {
		value = value + a.Get(i, k)*b.Get(k, j)
	}

	if c.Get(i, j) != value {
		t.Errorf("Wrong value on C(i,j): Got %f but expected %f!", c.Get(i, j), value)
	}

}

func TestMatrixGauss(t *testing.T) {

	a := CreateMatrix(3, 3, []float64{1, 2, 3, 2, 5, 3, 1, 0, 8})
	la := CreateLinAlg()
	c := la.Gauss(a)

	result := []float64{2, 5, 3, 0, -0.5, 1.5, 0, 0, -1}
	for i := 0; i < len(c.Elements); i++ {
		if c.Elements[i] != result[i] {
			t.Errorf("Expected %f but got %f", result[i], c.Elements[i])
		}
	}
}

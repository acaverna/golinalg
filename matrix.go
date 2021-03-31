package main

import "log"

type Matrix struct {
	Rows     int
	Cols     int
	elements []float64
}

func CreateMatrix(rows int, cols int, elements []float64) *Matrix {

	if rows < 0 || cols < 0 {
		log.Panic("The rows and cols must be greater or equal to 1")
	}

	if len(elements) != rows*cols {
		log.Panic("The size of elements array is incompatible with matrix dimensions")
	}

	return &Matrix{
		Rows:     rows,
		Cols:     cols,
		elements: elements,
	}
}

func CreateZeroMatrix(rows int, cols int) *Matrix {

	var elements []float64
	size := rows * cols
	for i := 0; i < size; i++ {
		elements = append(elements, 0)
	}

	return &Matrix{
		Rows:     rows,
		Cols:     cols,
		elements: elements,
	}
}

func (m *Matrix) Get(i int, j int) float64 {
	return m.elements[m.getIndex(i, j)]
}

func (m *Matrix) Set(i int, j int, value float64) {
	m.elements[m.getIndex(i, j)] = value
}

func (m *Matrix) getIndex(i int, j int) int {
	return (j - 1) + (i-1)*m.Cols
}

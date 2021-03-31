package main

import (
	"log"
	"math"
)

type linalg struct{}

func CreateLinAlg() *linalg {
	return &linalg{}
}

func (la *linalg) Transpose(a *Matrix) *Matrix {

	c := CreateZeroMatrix(a.Cols, a.Rows)

	for i := 1; i <= c.Rows; i++ {
		for j := 1; j <= c.Cols; j++ {
			c.Set(i, j, a.Get(j, i))
		}
	}

	return c

}

func (la *linalg) Add(a *Matrix, b *Matrix) *Matrix {

	if a.Rows != b.Rows || a.Cols != b.Cols {
		log.Panic("The matrices dimension are incompatibles")
	}

	c := CreateZeroMatrix(a.Rows, a.Cols)

	for i := 1; i <= c.Rows; i++ {
		for j := 1; j <= c.Cols; j++ {
			c.Set(i, j, a.Get(i, j)+b.Get(i, j))
		}
	}

	return c

}

func (la *linalg) Times(a *Matrix, b *Matrix) *Matrix {

	if a.Rows != b.Rows || a.Cols != b.Cols {
		log.Panic("The matrices dimension are incompatibles")
	}

	c := CreateZeroMatrix(a.Rows, a.Cols)

	for i := 1; i <= c.Rows; i++ {
		for j := 1; j <= c.Cols; j++ {
			c.Set(i, j, a.Get(i, j)*b.Get(i, j))
		}
	}

	return c

}

func (la *linalg) STimes(s float64, a *Matrix) *Matrix {

	c := CreateZeroMatrix(a.Rows, a.Cols)

	for i := 1; i <= c.Rows; i++ {
		for j := 1; j <= c.Cols; j++ {
			c.Set(i, j, s*a.Get(i, j))
		}
	}

	return c

}

func (la *linalg) Dot(a *Matrix, b *Matrix) *Matrix {

	if a.Cols != b.Rows {
		log.Panic("The matrices dimension are incompatibles")
	}

	c := CreateZeroMatrix(a.Rows, b.Cols)

	for i := 1; i <= a.Rows; i++ {
		for j := 1; j <= b.Cols; j++ {
			for k := 1; k <= a.Cols; k++ {
				c.Set(i, j, c.Get(i, j)+a.Get(i, k)*b.Get(k, j))
			}
		}
	}

	return c

}

func (la *linalg) Gauss(a *Matrix) *Matrix {

	c := CreateMatrix(a.Rows, a.Cols, a.Elements)

	for i := 1; i <= c.Rows; i++ {

		max := math.Inf(-1)
		maxLineIndex := 0
		for l := i; l <= c.Rows; l++ {
			if c.Get(l, i) > max {
				max = c.Get(l, i)
				maxLineIndex = l
			}
		}
		la.changeLinePosition(i, maxLineIndex, c)
		if c.Get(i, i) == 0 {
			log.Panicf("There is a zero value on main diagonal on line %d!", i)
		}

		for l := i + 1; l <= c.Rows; l++ {
			k := -1.0 * c.Get(l, i) / c.Get(i, i)
			la.multiplyLineByScalarAddOtherLine(i, l, k, c)
		}

	}

	return c
}

func (la *linalg) changeLinePosition(lineA int, lineB int, a *Matrix) {
	aux := 0.0
	for j := 1; j <= a.Cols; j++ {
		aux = a.Get(lineA, j)
		a.Set(lineA, j, a.Get(lineB, j))
		a.Set(lineB, j, aux)
	}
}

func (la *linalg) multiplyLineByScalar(line int, scalar float64, a *Matrix) {
	for j := 1; j <= a.Cols; j++ {
		a.Set(line, j, scalar*a.Get(line, j))
	}
}

func (la *linalg) multiplyLineByScalarAddOtherLine(lineA int, lineB int, scalar float64, a *Matrix) {
	for j := 1; j <= a.Cols; j++ {
		a.Set(lineB, j, scalar*a.Get(lineA, j)+a.Get(lineB, j))
	}
}

package main

import "log"

type linagl struct{}

func CreateLinAlg() *linagl {
	return &linagl{}
}

func (la *linagl) Transpose(a *Matrix) *Matrix {

	c := CreateZeroMatrix(a.Cols, a.Rows)

	for i := 1; i <= c.Rows; i++ {
		for j := 1; j <= c.Cols; j++ {
			c.Set(i, j, a.Get(j, i))
		}
	}

	return c

}

func (la *linagl) Add(a *Matrix, b *Matrix) *Matrix {

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

func (la *linagl) Times(a *Matrix, b *Matrix) *Matrix {

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

func (la *linagl) STimes(s float64, a *Matrix) *Matrix {

	c := CreateZeroMatrix(a.Rows, a.Cols)

	for i := 1; i <= c.Rows; i++ {
		for j := 1; j <= c.Cols; j++ {
			c.Set(i, j, s*a.Get(i, j))
		}
	}

	return c

}

func (la *linagl) Dot(a *Matrix, b *Matrix) *Matrix {

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

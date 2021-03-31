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

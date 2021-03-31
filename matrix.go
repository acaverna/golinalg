package golinalg

type Matrix struct {
	Rows     uint
	Cols     uint
	Elements []float64
}

func CreateMatrix(rows uint, cols uint, elements []float64) *Matrix {
	return &Matrix{
		Rows:     rows,
		Cols:     cols,
		Elements: elements,
	}
}

func (m *Matrix) Get(i uint, j uint) float64 {
	return m.Elements[m.getIndex(i, j)]
}

func (m *Matrix) Set(i uint, j uint, value float64) {
	m.Elements[m.getIndex(i, j)] = value
}

func (m *Matrix) getIndex(i uint, j uint) uint {
	return (j - 1) + (i-1)*m.Cols
}

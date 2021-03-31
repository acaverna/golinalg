package golinalg

type Vector struct {
	Dim      uint
	Elements []float64
}

func CreateVector(dim uint, elements []float64) *Vector {
	return &Vector{
		Dim:      dim,
		Elements: elements,
	}
}

func (v *Vector) Get(i uint) float64 {
	return v.Elements[i]
}

func (v *Vector) Set(i uint, value float64) {
	v.Elements[i] = value
}

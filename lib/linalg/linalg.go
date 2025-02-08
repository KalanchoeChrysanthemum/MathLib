package linalg

import (
	"math"
	"fmt"
)

type Vector struct {
	Components []float64
}

func VecFrom(components ...float64) Vector {
	return Vector{Components: components}
}

func (v Vector) Dim() int {
	return len(v.Components)
}

func (v Vector) Mag() float64 {
	sum := 0.0
	for _, c := range v.Components {
		sum += c * c
	}
	return math.Sqrt(sum)
}

func (v Vector) Norm() Vector {
	mag := v.Mag()
	if mag == 0 {
		return VecFrom(make([]float64, v.Dim())...)
	}

	norm := make([]float64, v.Dim())
	for i, c := range v.Components {
		norm[i] = c / mag
	}
	return VecFrom(norm...)
}

func (v *Vector) NormInPlace() {
	mag := v.Mag()
	if mag == 0 {
		return
	}
	for i := range v.Components {
		v.Components[i] /= mag
	}
}

func (v Vector) Dot(o Vector) (float64, error) {
	if v.Dim() != o.Dim() {
		return 0, fmt.Errorf("Supplied vector dimension does not match expected. Supplied: %d -- Expected: %d", v.Dim(), o.Dim())
	}
	
	res := 0.0
	for i, c := range v.Components {
		res += c * o.Components[i]
	}
	return res, nil
}

func (v Vector) Add(o Vector) (Vector, error) {
	if v.Dim() != o.Dim() {
		return Vector{}, fmt.Errorf("Supplied vector dimension does not match expected. Supplied: %d -- Expected: %d", v.Dim(), o.Dim())
	}

	sum := make([]float64, v.Dim())
	for i := range v.Components {
		sum[i] = v.Components[i] + o.Components[i]
	}
	return VecFrom(sum...), nil
}

func (v Vector) Subtract(o Vector) (Vector, error) {
	if v.Dim() != o.Dim() {
		return Vector{}, fmt.Errorf("Supplied vector dimension does not match expected. Supplied: %d -- Expected: %d", v.Dim(), o.Dim())
	}

	sub := make([]float64, v.Dim())
	for i := range v.Components {
		sub[i] = v.Components[i] - o.Components[i]
	}
	return VecFrom(sub...), nil
}


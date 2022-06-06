package golang_united_school_homework

import (
	"math"
)

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func NewTriangle(side float64) Triangle {
	return Triangle{Side: side}
}

func (t Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

func (t Triangle) CalcArea() float64 {
	return (math.Sqrt(3) / 4) * t.Side * t.Side
}
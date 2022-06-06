package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

func NewRectangle(height, weight float64) Rectangle {
	return Rectangle{Height: height, Weight: weight}
}

func (r Rectangle) CalcPerimeter() float64 {
	return 2 * r.Height + 2 * r.Weight
}

func (r Rectangle) CalcArea() float64 {
	return r.Height * r.Weight
}
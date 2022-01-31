package simplefactory

type ShapeType int

const (
	RectangleShape ShapeType = 1 << iota
	CircleShape
)

func NewShape(t ShapeType) Shape {
	switch t {
	case RectangleShape:
		return NewRectangle()
	case CircleShape:
		return NewCircle()
	default:
		return nil
	}
}

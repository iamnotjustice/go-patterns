package simplefactory

import "fmt"

type Shape interface {
	Draw() string
}

type TextRenderer struct {
	Shapes []Shape
}

func (tr *TextRenderer) Render() string {
	buf := ""

	for _, v := range tr.Shapes {
		buf = fmt.Sprintf("%s|%s", buf, v.Draw())
	}

	return buf + "|"
}

type Rect struct {
	posX int
	posY int

	width  int
	height int
}

func NewRectangle() Shape {
	return &Rect{}
}

func (r *Rect) Draw() string {
	return fmt.Sprintf("I'm a rectangle ([x:%d y:%d] [w:%d h:%d])", r.posX, r.posY, r.width, r.height)
}

type Circle struct {
	cenX int
	cenY int
	rad  int
}

func NewCircle() Shape {
	return &Circle{}
}

func (r *Circle) Draw() string {
	return fmt.Sprintf("I'm a circle ([cen_x:%d cen_y:%d] [rad:%d])", r.cenX, r.cenY, r.rad)
}

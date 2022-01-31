package factorymethod_test

import (
	"testing"

	"github.com/iamnotjustice/go-patterns/creational/factorymethod/creationmethod"
	"github.com/iamnotjustice/go-patterns/creational/factorymethod/simplefactory"

	"github.com/stretchr/testify/assert"
)

func Test_CreationMethod_ExactType(t *testing.T) {
	person := creationmethod.NewPerson(10, "Chief")

	assert.Equal(t, 10, person.Age)
	assert.Equal(t, "Chief", person.Name)
}

func Test_SimpleFactory_NoFactory(t *testing.T) {
	r := simplefactory.TextRenderer{Shapes: []simplefactory.Shape{
		simplefactory.NewCircle(), simplefactory.NewRectangle(), simplefactory.NewCircle(),
	}}

	assert.Equal(
		t,
		"|I'm a circle ([cen_x:0 cen_y:0] [rad:0])|I'm a rectangle ([x:0 y:0] [w:0 h:0])|I'm a circle ([cen_x:0 cen_y:0] [rad:0])|",
		r.Render(),
	)
}

func Test_SimpleFactory_WithFactory(t *testing.T) {
	r := simplefactory.TextRenderer{Shapes: []simplefactory.Shape{
		simplefactory.NewShape(simplefactory.CircleShape), simplefactory.NewShape(simplefactory.RectangleShape), simplefactory.NewShape(simplefactory.CircleShape),
	}}

	assert.Equal(
		t,
		"|I'm a circle ([cen_x:0 cen_y:0] [rad:0])|I'm a rectangle ([x:0 y:0] [w:0 h:0])|I'm a circle ([cen_x:0 cen_y:0] [rad:0])|",
		r.Render(),
	)
}

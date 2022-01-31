package builder

import "fmt"

type ComponentType string
type HouseComponents map[ComponentType]*HouseComponent

const (
	Roof    ComponentType = "ROOF"
	Walls   ComponentType = "WALLS"
	Windows ComponentType = "WINDOWS"
	Floor   ComponentType = "FLOOR"
)

type HouseComponent struct {
	Name string
	Type ComponentType
}

func (hc *HouseComponent) String() string {
	return fmt.Sprintf("[Component Name:%s, Component Type: %s]", hc.Name, hc.Type)
}

type House struct {
	components HouseComponents
}

// Has prints all house components to stdout.
func (h *House) Has(t ComponentType) *HouseComponent {
	return h.components[t]
}

type HouseBuilder interface {
	WithRoof() HouseBuilder
	WithWalls() HouseBuilder
	WithWindows() HouseBuilder
	WithFloor() HouseBuilder
	Build() *House
}

type SimpleHouseBuilder struct {
	components HouseComponents
}

func NewSimpleHouseBuilder() HouseBuilder {
	return &SimpleHouseBuilder{
		components: make(HouseComponents),
	}
}

func (b *SimpleHouseBuilder) WithRoof() HouseBuilder {
	b.components[Roof] = &HouseComponent{"Simple roof", Roof}

	return b
}

func (b *SimpleHouseBuilder) WithWalls() HouseBuilder {
	b.components[Walls] = &HouseComponent{"Simple walls", Walls}

	return b
}

func (b *SimpleHouseBuilder) WithWindows() HouseBuilder {
	b.components[Windows] = &HouseComponent{"Simple windows", Windows}

	return b
}

func (b *SimpleHouseBuilder) WithFloor() HouseBuilder {
	b.components[Floor] = &HouseComponent{"Simple floor", Floor}

	return b
}

func (b *SimpleHouseBuilder) Build() *House {
	return &House{
		components: b.components,
	}
}

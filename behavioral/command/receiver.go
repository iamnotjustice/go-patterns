package command

type Receiver struct{}

func (r *Receiver) MoveUp() string {
	return "moved one tile up"
}

func (r *Receiver) MoveDown() string {
	return "moved one tile down"
}

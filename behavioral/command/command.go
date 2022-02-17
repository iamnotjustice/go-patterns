package command

type Command interface {
	Exec() string
}

type MoveUpCommand struct {
	Receiver *Receiver
}

func (c *MoveUpCommand) Exec() string {
	return c.Receiver.MoveUp()
}

type MoveDownCommand struct {
	Receiver *Receiver
}

func (c *MoveDownCommand) Exec() string {
	return c.Receiver.MoveDown()
}

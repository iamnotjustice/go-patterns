package command

type Invoker struct {
	commands []Command
}

func (i *Invoker) SaveCommand(c Command) {
	i.commands = append(i.commands, c)
}

func (i *Invoker) RemoveCommand() {
	if len(i.commands) != 0 {
		// remove last
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func (i *Invoker) ExecuteAll() string {
	res := ""
	for _, v := range i.commands {
		res += v.Exec() + "\n"
	}

	return res
}

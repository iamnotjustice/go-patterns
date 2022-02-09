package adapter

import "fmt"

type BigEndianBytes struct {
	Bytes string
}

func (b *BigEndianBytes) String() string {
	return fmt.Sprintf(b.Bytes)
}

type PowerPCCommand struct {
	Command  string
	OperandA BigEndianBytes
	OperandB BigEndianBytes
}

type PowerPCProcessor interface {
	ExecuteCommand(command *PowerPCCommand) string
}

type PPC struct{}

func NewPPC() PowerPCProcessor {
	return &PPC{}
}

func (p *PPC) ExecuteCommand(command *PowerPCCommand) string {
	return fmt.Sprintf("PowerPC exec: %s\t%s %s", command.Command, command.OperandA.Bytes, command.OperandB.Bytes)
}

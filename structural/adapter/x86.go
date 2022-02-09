package adapter

import "fmt"

type LittleEndianBytes struct {
	Bytes string
}

func (b *LittleEndianBytes) String() string {
	return fmt.Sprintf(b.Bytes)
}

type X86 struct{}

func (p *X86) Execute(command string, operandA, operandB LittleEndianBytes) string {
	return fmt.Sprintf("exec: %s\t%s %s", command, operandA.Bytes, operandB.Bytes)
}

// Our "translation" adapter structure, which makes it possible to execute PPC commands on x86.
type X86toPPCAdapter struct {
	p *X86
}

func (adapter *X86toPPCAdapter) ExecuteCommand(command *PowerPCCommand) string {
	return adapter.p.Execute(
		adapter.ConvertCommand(command.Command),
		adapter.ConvertBytes(command.OperandA),
		adapter.ConvertBytes(command.OperandB),
	)
}

func (adapter *X86toPPCAdapter) ConvertBytes(input BigEndianBytes) LittleEndianBytes {
	r := []rune(input.Bytes)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+2, j-2 {
		// we reverse the order of the bytes (we need to work on pairs here)
		t := r[i]
		r[i] = r[j-1]
		r[j-1] = t
		t = r[j]
		r[j] = r[i+1]
		r[i+1] = t
	}
	return LittleEndianBytes{string(r)}
}

func (adapter *X86toPPCAdapter) ConvertCommand(input string) string {
	// 						*** 					//
	// let's imagine there's something more complex //
	// 						*** 					//
	return input
}

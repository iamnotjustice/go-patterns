package adapter_test

import (
	"testing"

	"github.com/iamnotjustice/go-patterns/structural/adapter"

	"github.com/stretchr/testify/assert"
)

func Test_NoAdapter(t *testing.T) {
	mPC := adapter.NewPPC()

	c := &adapter.PowerPCCommand{
		Command:  "MOV",
		OperandA: adapter.BigEndianBytes{"78BA68FF"},
		OperandB: adapter.BigEndianBytes{"110022FF"},
	}

	assert.Equal(t, "PowerPC exec: MOV	78BA68FF 110022FF", mPC.ExecuteCommand(c))
}

func Test_Adapter_ToX86(t *testing.T) {
	x86PC := adapter.X86toPPCAdapter{}

	c := &adapter.PowerPCCommand{
		Command:  "MOV",
		OperandA: adapter.BigEndianBytes{"78BA68FF"},
		OperandB: adapter.BigEndianBytes{"110022FF"},
	}

	assert.Equal(t, "exec: MOV	FF68BA78 FF220011", x86PC.ExecuteCommand(c))
}

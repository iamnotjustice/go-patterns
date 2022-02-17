package command_test

import (
	"testing"

	"github.com/iamnotjustice/go-patterns/behavioral/command"
	"github.com/stretchr/testify/assert"
)

func Test_Command(t *testing.T) {
	expected := `moved one tile up
moved one tile down
`
	receiver := &command.Receiver{}

	invoker := &command.Invoker{}

	invoker.SaveCommand(
		&command.MoveUpCommand{receiver},
	)
	invoker.SaveCommand(
		&command.MoveDownCommand{receiver},
	)

	result := invoker.ExecuteAll()

	assert.Equal(t, expected, result)
}

func Test_Command_WithRemove(t *testing.T) {
	invoker := &command.Invoker{}
	receiver := &command.Receiver{}

	invoker.SaveCommand(&command.MoveUpCommand{receiver})
	invoker.SaveCommand(&command.MoveDownCommand{receiver})

	invoker.RemoveCommand()

	result := invoker.ExecuteAll()

	assert.Equal(t, "moved one tile up\n", result)
}

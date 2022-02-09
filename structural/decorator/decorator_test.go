package decorator_test

import (
	"testing"

	"github.com/iamnotjustice/go-patterns/structural/decorator/composition"
	"github.com/iamnotjustice/go-patterns/structural/decorator/fn"

	"github.com/stretchr/testify/assert"
)

func Test_Decorator_Fn(t *testing.T) {
	originalResult := fn.Process([]string{"a", "b", "c"})

	assert.Equal(t, "|a|b|c|", originalResult)

	wrappedProccessFunc := fn.ProcessWithColon(fn.Process)

	modifiedResult := wrappedProccessFunc([]string{"a", "b", "c"})

	assert.Equal(t, ":a:b:c:", modifiedResult)
}

func Test_Decorator_Composition(t *testing.T) {
	originalProc := composition.OriginalProcessor{}
	originalResult := originalProc.Process([]string{"a", "b", "c"})

	assert.Equal(t, "|a|b|c|", originalResult)

	wrappedProc := composition.NewDecoratedProcessor(originalProc)

	modifiedResult := wrappedProc.Process([]string{"a", "b", "c"})

	assert.Equal(t, ":a:b:c:", modifiedResult)
}

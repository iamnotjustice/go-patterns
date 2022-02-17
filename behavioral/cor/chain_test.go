package cor_test

import (
	"testing"

	"github.com/iamnotjustice/go-patterns/behavioral/cor"

	"github.com/stretchr/testify/assert"
)

func Test_Chain_ShouldFallback(t *testing.T) {
	chain := cor.NewChain()

	result := chain.Handle("img.png")

	assert.Equal(t, "no file handler found!", result)
}

func Test_Chain_JsonHandler(t *testing.T) {
	chain := cor.NewChain()

	result := chain.Handle("settings.json")

	assert.Equal(t, "changed JSON file", result)
}

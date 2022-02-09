package proxy_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/iamnotjustice/go-patterns/structural/proxy"
	"github.com/stretchr/testify/assert"
)

var (
	sampleText = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium.",
		"At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis.",
	}
	expectedResult = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium.
At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis.
`
)

func Test_Proxy(t *testing.T) {
	var str bytes.Buffer

	// we redirect our stdout to bytes buffer
	// so we can easily test our log output
	log.SetOutput(&str)

	compilerProxy := proxy.NewProxyCompiler(&proxy.RealCompiler{})

	assert.Equal(t, expectedResult, compilerProxy.Compile(sampleText))

	// check that we indeed logged something
	assert.NotEmpty(t, str.String())
}

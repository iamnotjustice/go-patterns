package singleton_test

import (
	"testing"

	"github.com/iamnotjustice/go-patterns/creational/singleton/init_singleton"
	"github.com/iamnotjustice/go-patterns/creational/singleton/mutex_singleton"
	"github.com/iamnotjustice/go-patterns/creational/singleton/once_singleton"

	"github.com/stretchr/testify/assert"
)

func Test_Singleton_WithOnce(t *testing.T) {
	bo := once_singleton.GetOrder()

	assert.Equal(t, "big-endian", bo)
}

func Test_Singleton_WithInit(t *testing.T) {
	arch := init_singleton.GetArch()

	assert.Equal(t, "PowerPC", arch)
}

func Test_Singleton_WithMutex(t *testing.T) {
	bytes := mutex_singleton.GetAddressableBytes()

	assert.Equal(t, 32, bytes)
}

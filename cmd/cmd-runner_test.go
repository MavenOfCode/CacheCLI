package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockKeyValueCache (t *testing.T){
	t.Run("it creates a mock cache", func(t *testing.T) {
		mockCache := NewMockSimpleKVCache()
		assert.NotNil(t, mockCache)
	})
}

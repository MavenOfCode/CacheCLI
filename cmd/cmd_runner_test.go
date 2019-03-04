package cmd

import (
	"CacheCLI/kvcache"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockKeyValueCache (t *testing.T){
	t.Run("it creates a mock cache", func(t *testing.T) {
		mockCache := kvcache.NewMockSimpleKVCache()
		assert.NotNil(t, mockCache)
	})
}

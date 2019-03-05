package cmd

import (
	"CacheCLI/kvcache"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommandRunner_DeleteCmd(t *testing.T) {
	//set up for all tests
	var RootCmd = &cobra.Command{Use:"kvc"}
	mockCache := NewMockSimpleKVCache()
	require.NotNil(t, mockCache)

	t.Run("it deletes")
}
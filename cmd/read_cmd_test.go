package cmd

import (
	"CacheCLI/kvcache"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommandRunner_ReadCmd(t *testing.T) {
	var RootCmd = &cobra.Command{Use:"kvc"}
	mockCache := kvcache.NewMockSimpleKVCache()
	require.NotNil(t, mockCache)


	t.Run("it reads", func(t *testing.T) {
		ReturnString := "hi"

		args := []string{ReturnString}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.ReadCmd(RootCmd, args)

		assert.ObjectsAreEqualValues(err, ReturnString)
	})

	t.Run("read command returns error when key is invalid", func(t *testing.T) {

		args := []string{"false"}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.ReadCmd(RootCmd, args)
		fmt.Println(err)

		assert.Error(t, err)
	})

	t.Run("read command returns error when cache is nil ", func(t *testing.T) {

		args := []string{"false"}

		commandRun := CommandRunner{cache:nil}
		err := commandRun.ReadCmd(RootCmd, args)
		fmt.Println(err)

		assert.ObjectsAreEqualValues(err, "read failed: cache empty - read failed")
	})

}

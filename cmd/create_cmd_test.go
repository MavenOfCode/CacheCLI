package cmd

import (
	"CacheCLI/kvcache"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCmdRunner_CreateCmd(t *testing.T) {
	var RootCmd = &cobra.Command{Use:"cli"}
	mockCache := kvcache.NewMockSimpleKVCache()
	require.NotNil(t, mockCache)

	//testing only create command not create method or kvc
	t.Run("it creates", func(t *testing.T) {
		key := "testString"
		value := "testValueString"

		args := []string{key,value}

		//object literals struct literals
		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.CreateCmd(RootCmd, args)

		assert.NoError(t,err)

		b, _ := mockCache.Read(key)
		assert.Equal(t, b, "")
	})

	t.Run("create command returns error when cache is nil", func(t *testing.T) {

		key := "keyTest"
		value := "testValueString"

		args := []string{key,value}

		commandRun := CommandRunner{cache:nil}
		err := commandRun.CreateCmd(RootCmd, args)

		assert.ObjectsAreEqualValues(err, "cache not initialized - create failed: ")

	})

	t.Run("create command returns error when one of 2 args missing", func(t *testing.T) {

		key := "testKey"
		args := []string{key}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.CreateCmd(RootCmd, args)
		assert.ObjectsAreEqualValues(err, "create failed: insufficient arguments provided")

	})
}


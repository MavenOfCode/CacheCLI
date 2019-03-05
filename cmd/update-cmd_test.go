package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCommandRunner_UpdateCmd(t *testing.T) {
	var RootCmd = &cobra.Command{Use:"kvc"}
	mockCache := NewMockSimpleKVCache()
	require.NotNil(t, mockCache)

	t.Run("it updates", func(t *testing.T) {
		ReturnString := "ReturnString"
		value :="hi"
		args := []string{ReturnString,value}

		commandCache := CommandRunner{cache:mockCache}
		commandCache.CreateCmd(RootCmd,args)

		args2 :=[]string{ReturnString,"bye"}
		err := commandCache.UpdateCmd(RootCmd,args2)

		//mockUpdate returns nil
		assert.Nil(t, err)

	})
	
	t.Run("update returns error when invalid key provided", func(t *testing.T) {

		key := "key"
		value := "value"
		args := []string{key,value}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.UpdateCmd(RootCmd, args)
		assert.Error(t, err)
	})

	t.Run("update returns error when cache is nil", func(t *testing.T) {
		key := "key"
		value := "value"
		args := []string{key,value}

		commandRun := CommandRunner{cache:nil}
		err := commandRun.UpdateCmd(RootCmd, args)
		assert.Error(t, err)
	})

	t.Run("update returns error when min args aren't provided", func(t *testing.T) {
		key := "key"
		args := []string{key}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.UpdateCmd(RootCmd, args)
		assert.Error(t, err)
	})
}

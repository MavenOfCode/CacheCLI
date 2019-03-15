package cmd

import (
	"CacheCLI/kvcache"
	
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	
	"testing"
)

func TestMockSimpleKeyValueCache(t *testing.T) {
	t.Run("it creates a mock cache", func(t *testing.T) {
		mockCache := kvcache.NewMockSimpleKVCache(true, "")
		assert.NotNil(t, mockCache)
	})
}

/*Command Tests using Command Runner with MockSKVC*/
func TestCommandRunner_CreateCmd(t *testing.T) {
	var RootCmd = &cobra.Command{Use: "kvc"}
	mockCache := kvcache.NewMockSimpleKVCache(false, "")
	require.NotNil(t, mockCache)

	t.Run("it creates", func(t *testing.T) {
		key := "true"
		value := "testValueString"

		args := []string{key, value}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.CreateCmd(RootCmd, args)

		assert.Nil(t, err)
		assert.NotNil(t, res)

		b, _ := mockCache.Read(key)
		assert.Equal(t, b, value)
	})

	t.Run("create command returns error when cache is nil", func(t *testing.T) {

		key := "true"
		value := "testValueString"

		args := []string{key, value}

		commandRun := CommandRunner{cache: nil}
		res, err := commandRun.CreateCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})

	t.Run("create command returns error when one of 2 args missing", func(t *testing.T) {

		key := "testKey"
		args := []string{key}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.CreateCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})

	t.Run("create command returns error when key arg is missing", func(t *testing.T) {
		value := "return"
		args := []string{value}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.CreateCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})

	t.Run("create command returns error when value arg is missing", func(t *testing.T) {

		key := "true"
		args := []string{key}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.CreateCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})
}

func TestCommandRunner_ReadCmd(t *testing.T) {
	var RootCmd = &cobra.Command{Use: "kvc"}
	mockCache := kvcache.NewMockSimpleKVCache(false, "")
	require.NotNil(t, mockCache)

	t.Run("it reads", func(t *testing.T) {
		ReturnString := "ReturnString"
		Success := "true"
		args := []string{Success, ReturnString}

		commandRun := CommandRunner{cache: mockCache}
		_, _ = commandRun.CreateCmd(RootCmd, args)

		args2 := []string{Success}

		res, err := commandRun.ReadCmd(RootCmd, args2)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("read command returns error when key is invalid", func(t *testing.T) {

		args := []string{"false"}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.ReadCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})

	t.Run("read command returns error when args are insufficient", func(t *testing.T) {

		args := []string{}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.ReadCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})

	t.Run("read command returns error when cache is nil ", func(t *testing.T) {

		args := []string{"false"}

		commandRun := CommandRunner{}
		res, err := commandRun.ReadCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})
}

func TestCommandRunner_UpdateCmd(t *testing.T) {
	var rootCmd = &cobra.Command{Use: "kvc"}
	mockCache := kvcache.NewMockSimpleKVCache(false, "")
	require.NotNil(t, mockCache)

	t.Run("it updates", func(t *testing.T) {
		ReturnString := "ReturnString"
		Success := "true"
		args := []string{Success, ReturnString}

		commandCache := CommandRunner{cache: mockCache}
		_, _ = commandCache.CreateCmd(rootCmd, args)

		args2 := []string{Success, "bye"}
		res, err := commandCache.UpdateCmd(rootCmd, args2)

		//mockUpdate returns nil
		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("update returns error when invalid key provided", func(t *testing.T) {
		Success := "false"
		ReturnString := "ReturnString"

		args := []string{Success, ReturnString}

		commandCache := CommandRunner{cache: mockCache}
		_, _ = commandCache.CreateCmd(rootCmd, args)

		Success = "key"
		ReturnString = "value"
		args2 := []string{Success, ReturnString}
		res, err := commandCache.UpdateCmd(rootCmd, args2)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})

	t.Run("update returns error when cache is nil", func(t *testing.T) {
		key := "true"
		value := "value"
		args := []string{key, value}

		commandRun := CommandRunner{cache: nil}
		res, err := commandRun.UpdateCmd(rootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})

	t.Run("update returns error when required args aren't provided", func(t *testing.T) {
		key := "key"
		args := []string{key}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.UpdateCmd(rootCmd, args)
		assert.Error(t, err)
		assert.Equal(t, res, "")

	})

	t.Run("update returns error when key is empty string", func(t *testing.T) {
		key := ""
		value := "notEmpty"
		args := []string{key, value}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.UpdateCmd(rootCmd, args)
		assert.Error(t, err)
		assert.Equal(t, res, "")
	})
}

func TestCommandRunner_DeleteCmd(t *testing.T) {
	var RootCmd = &cobra.Command{Use: "kvc"}
	mockCache := kvcache.NewMockSimpleKVCache(false, "returnString")
	require.NotNil(t, mockCache)

	t.Run("it deletes", func(t *testing.T) {
		Success := "true"

		args := []string{Success}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.DeleteCmd(RootCmd, args)

		assert.Nil(t, err)
		assert.NotNil(t, res)
	})

	t.Run("delete command returns error when key is does not exist in cache", func(t *testing.T) {
		args := []string{"betty"}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.DeleteCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})

	t.Run("delete command returns error when args do not equal one", func(t *testing.T) {
		args := []string{}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.DeleteCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})

	t.Run("delete command returns error when key is empty string", func(t *testing.T) {
		args := []string{""}

		commandRun := CommandRunner{cache: mockCache}
		res, err := commandRun.DeleteCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})

	t.Run("delete command returns error when cache is nil ", func(t *testing.T) {

		args := []string{"false"}

		commandRun := CommandRunner{cache: nil}
		res, err := commandRun.DeleteCmd(RootCmd, args)

		assert.Error(t, err)
		assert.Equal(t, res, "")
	})
}

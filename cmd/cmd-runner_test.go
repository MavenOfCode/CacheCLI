package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMockSimpleKeyValueCache (t *testing.T){
	t.Run("it creates a mock cache", func(t *testing.T) {
		mockCache := NewMockSimpleKVCache()
		assert.NotNil(t, mockCache)
	})
}

/*Command Tests using Command Runner with MockSKVC*/

func TestCommandRunner_CreateCmd(t *testing.T) {
	var RootCmd = &cobra.Command{Use:"kvc"}
	mockCache := NewMockSimpleKVCache()
	require.NotNil(t, mockCache)

	t.Run("it creates", func(t *testing.T) {
		key := "testString"
		value := "testValueString"

		args := []string{key,value}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.CreateCmd(RootCmd, args)

		assert.Nil (t,err, "create works")

		assert.NoError(t,err,"no error generated")

		b, _ := mockCache.Read(key)
		assert.Equal(t, b, "")
	})

	t.Run("create command returns error when cache is nil", func(t *testing.T) {

		key := "keyTest"
		value := "testValueString"

		args := []string{key,value}

		commandRun := CommandRunner{cache:nil}
		err := commandRun.CreateCmd(RootCmd, args)

		assert.Equal(t, err.Error(), "create failed: cache not initialized")
	})

	t.Run("create command returns error when one of 2 args missing", func(t *testing.T) {

		key := "testKey"
		args := []string{key}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.CreateCmd(RootCmd, args)
		assert.ObjectsAreEqualValues(err, "create failed: insufficient arguments provided")
	})

	t.Run("create command returns error when both args missing", func(t *testing.T) {

		args := []string{}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.CreateCmd(RootCmd, args)
		assert.ObjectsAreEqualValues(err, "create failed: insufficient arguments provided")
	})
}

func TestCommandRunner_ReadCmd(t *testing.T) {
	var RootCmd = &cobra.Command{Use:"kvc"}
	mockCache := NewMockSimpleKVCache()
	require.NotNil(t, mockCache)


	t.Run("it reads", func(t *testing.T) {
		ReturnString := "ReturnString"
		Success := "true"
		args := []string{Success, ReturnString}

		commandRun := CommandRunner{cache:mockCache}
		commandRun.CreateCmd(RootCmd, args)

		args2 := []string{Success}

		err := commandRun.ReadCmd(RootCmd, args2)
		assert.Nil(t,err)
	})

	t.Run("read command returns error when key is invalid", func(t *testing.T) {

		args := []string{"false"}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.ReadCmd(RootCmd, args)

		assert.Error(t, err)
	})

	t.Run("read command returns error when args are insufficient", func(t *testing.T) {

		args := []string{}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.ReadCmd(RootCmd, args)

		assert.Error(t, err)
	})

	t.Run("read command returns error when cache is nil ", func(t *testing.T) {

		args := []string{"false"}

		commandRun := CommandRunner{}
		err := commandRun.ReadCmd(RootCmd, args)
		fmt.Println(commandRun)

		assert.Error(t, err)
	})
}

func TestCommandRunner_UpdateCmd(t *testing.T) {
	var RootCmd = &cobra.Command{Use:"kvc"}
	mockCache := NewMockSimpleKVCache()
	require.NotNil(t, mockCache)

	t.Run("it updates", func(t *testing.T) {
		ReturnString := "ReturnString"
		Success := "true"
		args := []string{Success, ReturnString}

		commandCache := CommandRunner{cache:mockCache}
		commandCache.CreateCmd(RootCmd,args)

		args2 :=[]string{Success,"bye"}
		err := commandCache.UpdateCmd(RootCmd,args2)
		fmt.Println(err)

		//mockUpdate returns nil
		assert.Nil(t, err)

	})

	t.Run("update returns error when invalid key provided", func(t *testing.T) {
		Success :="false"
		ReturnString := "ReturnString"

		args := []string{Success,ReturnString}

		commandCache := CommandRunner{cache:mockCache}
		commandCache.CreateCmd(RootCmd,args)

		key := "key"
		value2 := "value"
		args2 := []string{key,value2}
		err := commandCache.UpdateCmd(RootCmd, args2)

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

	t.Run("update returns error when key is empty string", func(t *testing.T) {
		key := ""
		value:="notEmpty"
		args := []string{key,value}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.UpdateCmd(RootCmd, args)
		fmt.Println(err)
		assert.Error(t, err)
	})
}

func TestCommandRunner_DeleteCmd(t *testing.T) {
	var RootCmd = &cobra.Command{Use:"kvc"}
	mockCache := NewMockSimpleKVCache()
	require.NotNil(t, mockCache)

	t.Run("it deletes", func(t *testing.T) {
		Success := "true"

		args := []string{Success}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.DeleteCmd(RootCmd, args)

		assert.Nil(t,err)
	})

	t.Run("delete command returns error when key is invalid", func(t *testing.T) {
		args := []string{"betty"}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.DeleteCmd(RootCmd, args)

		assert.Error(t, err)
	})

	t.Run("delete command returns error when args are insufficient", func(t *testing.T) {
		args := []string{}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.DeleteCmd(RootCmd, args)

		assert.Error(t, err)
	})

	t.Run("delete command returns error when key is empty string", func(t *testing.T) {
		args := []string{""}

		commandRun := CommandRunner{cache:mockCache}
		err := commandRun.DeleteCmd(RootCmd, args)

		assert.Error(t, err)
	})

	t.Run("delete command returns error when cache is nil ", func(t *testing.T) {

		args := []string{"false"}

		commandRun := CommandRunner{cache:nil}
		err := commandRun.DeleteCmd(RootCmd, args)

		assert.Error(t,err)
	})
}

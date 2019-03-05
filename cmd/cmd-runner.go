package cmd

import (
	"CacheCLI/kvcache"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

type CommandRunner struct {
	cache kvcache.KeyValueCache
}

/* MockCache struct and implementation of KVC interface for testing of KVC CLI commands */
type MockKeyValueCache struct{
	WillFail bool
	ReturnString string
}

func (m *MockKeyValueCache) Create(key,value string) error{
	return nil
}

func (m *MockKeyValueCache) Read(key string) (string,error){
	if m.WillFail {
		return m.ReturnString, nil
	}
	return "", fmt.Errorf("read error")
}

func (m *MockKeyValueCache) Update(key, value string) error  {
	return nil
}

func (m *MockKeyValueCache) Delete(key string) error{
	return nil
}

//constructor function for generating test MockCache
func NewMockSimpleKVCache() *MockKeyValueCache{
	return &MockKeyValueCache{}
}

/*Commands for CLI using CommandRunner*/
func (c *CommandRunner) CreateCmd (cmd *cobra.Command, args []string) error {

	if len(args) < 2{
		return errors.New("create failed: insufficient arguments provided")
	}

	if  c != nil {
		createResult := c.cache.Create(args[0],args[1])
		if createResult == nil {
			fmt.Printf("create success:  cache '%v' ", c.cache)
			fmt.Println()
			return nil
		}
	}
	return errors.New("create failed: cache not initialized")
}

func (c *CommandRunner) ReadCmd(cmd *cobra.Command, args []string) error  {

	if len(args) < 1{
		return errors.New("read failed: at least one argument required")
	}

	fmt.Println(c)
	if c.cache == nil {
		fmt.Println("is nil")
		return errors.New("read failed: cache empty - read failed")
	}

	fmt.Println("is not nil")

	readResult, err := c.cache.Read(args[0])
	if err !=nil {
		return err
	}
	fmt.Println(">> value for key is: ", readResult)
	return nil
}

func (c *CommandRunner) UpdateCmd(cmd *cobra.Command, args []string) error {

	if len(args) < 2{
		return errors.New("update failed: insufficient arguments provided")
	}
	if c.cache == nil {
		return errors.New("update failed: cache not initialized ")
	}

	updateResult := c.cache.Update(args[0],args[1])
	if updateResult == nil {
		fmt.Printf("update success:  cache '%v' ", c.cache)
		fmt.Println()
		return nil
	}
	fmt.Println(updateResult)
	return errors.New("")
}

func (c *CommandRunner) DeleteCmd(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		return errors.New("delete failed: at least one argument required ")
	}

	if c.cache == nil {
		return errors.New("delete failed: cache not initialized - delete failed")
	}

	deleteResult := c.cache.Delete(args[0])
	if deleteResult == nil {
		fmt.Printf("delete success: cache '%v' ", c.cache)
		fmt.Println()
		return nil
	}
	fmt.Println(deleteResult)
	return errors.New("")
}
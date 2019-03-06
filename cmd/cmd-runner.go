package cmd

import (
	"CacheCLI/kvcache"
	"fmt"
	"github.com/spf13/cobra"
)

type CommandRunner struct {
	cache kvcache.KeyValueCache
}

/*Commands for CLI using CommandRunner*/
func (c *CommandRunner) CreateCmd(cmd *cobra.Command, args []string) (string, error) {
	if len(args) != 2{
		return "", fmt.Errorf("create failed: insufficient arguments provided")
	}

	if  c.cache != nil {
		err := c.cache.Create(args[0],args[1])
		if err == nil {
			fmt.Printf("create success: cache '%v'\n ", c.cache)
			return "TEST", nil
		}
		return "", fmt.Errorf("create failed: '%v' ", err)
	}
	return "", fmt.Errorf("create failed: cache not initialized")
}

func (c *CommandRunner) ReadCmd(cmd *cobra.Command, args []string) error  {
	if len(args) != 1{
		return fmt.Errorf("read failed: only one argument required")
	}

	if c.cache == nil {
		return fmt.Errorf("read failed: cache nil")
	}

	readResult, err := c.cache.Read(args[0])
	if err !=nil {
		return err
	}
	fmt.Printf(">> value for key '%v' is '%v'", args[0], readResult)
	return nil
}

func (c *CommandRunner) UpdateCmd(cmd *cobra.Command, args []string) (string, error) {
	if len(args) != 2{
		return "", fmt.Errorf("update failed: two arguments required")
	}

	if c.cache == nil {
		return "", fmt.Errorf("update failed: cache is nil ")
	}

	err := c.cache.Update(args[0],args[1])
	if err == nil {
		//fmt.Printf("update success: cache '%v' \n", c.cache)
		res := fmt.Sprintf("update success: cache '%v' \n", c.cache)
		return res, nil
	}
	return "", fmt.Errorf("update failed: '%v'", err)
}

func (c *CommandRunner) DeleteCmd(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("delete failed: one argument required ")
	}

	if c.cache == nil {
		return fmt.Errorf("delete failed: cache is nil")
	}

	err := c.cache.Delete(args[0])
	if err == nil {
		fmt.Printf("delete success: cache '%v' \n", c.cache)

		return nil
	}
	return fmt.Errorf("delete failed: '%v'", err)
}

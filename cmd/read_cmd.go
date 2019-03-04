package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func (c *CommandRunner) ReadCmd(cmd *cobra.Command, args []string) error  {
	//pre-seeding cache for read command for now since cache won't persist until CLI/Cache connection built
	//cache.Create("name", "harley")
	//cache.Create("animal", "horse")
	//cache.Create("kitten", "Bene")

	if len(args) < 1{
		return errors.New("read failed: at least one argument required")

	}

	if cache == nil {
		return errors.New("cache empty - read failed: ")
	}

	readResult, err := cache.Read(args[0])
	if err !=nil {
		return err
	}
	fmt.Println(">> value for key is: ", readResult)
	return nil
}

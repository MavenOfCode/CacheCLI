package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func (c *CommandRunner) UpdateCmd(cmd *cobra.Command, args []string) error {
	////pre-seeding cache for read command for now since cache won't persist until CLI/Cache connection built
	cache.Create("name", "harley")
	cache.Create("animal", "horse")
	cache.Create("kitten", "Bene")
	if len(args) < 2{
		return errors.New("update failed: insufficient arguments provided")
	}
	if cache == nil {
		return errors.New("update failed: cache not initialized ")
	}

	updateResult := cache.Update(args[0],args[1])
	if updateResult == nil {
		fmt.Printf("update success:  cache '%v' ", cache)
		fmt.Println()
		return nil
	}
	fmt.Println(updateResult)
	return errors.New("")
}
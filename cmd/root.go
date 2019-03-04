package cmd

import (
	"KVCache/kvcache"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

//use constructor from kvcache package for global access to struct and it's private fields
var cache = kvcache.NewSimpleKVCache()

//use struct CommandRunner to enable running of either Mock or Real commands with Mock or Simple KVCache
var CommandRun = CommandRunner{cache:cache}

//make root command not executable without subcommand by not providing a 'Run' for the 'rootCmd'
var RootCmd = &cobra.Command{Use:"cli"}
var createCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.MinimumNArgs(2),
	Short: "create key-value pair",
	Long:  "create key value strings into the key-value cache",
	RunE:   CommandRun.CreateCmd,
}

//trying use of minimum args in command to avoid writing RunE function with error to test for args length
var readCmd = &cobra.Command{
	Use:  "read",
	Short: "read given key and return value",
	Args: cobra.MinimumNArgs(1),
	Long: "read value string out to command line from key-value cache given key string input from command line",
	//RunE:
}

var updateCmd = &cobra.Command{
	Use:  "update",
	Args: cobra.MinimumNArgs(2),
	Short: "update key-value pair",
	Long:  "update key value strings into the key-value cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cache == nil {
			return errors.New("cache not initialized - update failed: ")
		}
		//pre-seeding cache for read command for now since cache won't persist until CLI/Cache connection built
		cache.Create("name", "harley")
		cache.Create("animal", "horse")
		cache.Create("kitten", "Bene")
		updateResult := cache.Update(args[0],args[1])
		if updateResult == nil {
			fmt.Printf("update success:  cache '%v' ", cache)
			fmt.Println()
			return nil
		}
		fmt.Println(updateResult)
		return errors.New("")
	},
}

var deleteCmd = &cobra.Command{
	Use:  "delete",
	Args: cobra.MinimumNArgs(1),
	Short: "delete key-value pair",
	Long:  "delete key value strings into the key-value cache",
	RunE: func(cmd *cobra.Command, args []string) error {
		if cache == nil {
			return errors.New("cache not initialized - delete failed: ")
		}
		//pre-seeding cache for read command for now since cache won't persist until CLI/Cache connection built
		cache.Create("name", "harley")
		cache.Create("animal", "horse")
		cache.Create("kitten", "Bene")
		deleteResult := cache.Delete(args[0])
		if deleteResult == nil {
			fmt.Printf("delete success: cache '%v' ", cache)
			fmt.Println()
			return nil
		}
		fmt.Println(deleteResult)
		return errors.New("")
	},
}


// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//attach subcommands to rootcommand
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(readCmd)
	RootCmd.AddCommand(updateCmd)
	RootCmd.AddCommand(deleteCmd)
	RootCmd.Execute()
}


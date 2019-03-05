package cmd

import (
	"CacheCLI/kvcache"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

//use struct CommandRunner to enable running of either Mock or Real commands with Mock or Simple KVCache
var CommandRun = CommandRunner{
	cache: kvcache.NewSimpleKVCache(),
}

//make root command not executable without subcommand by not providing a 'Run' for the 'rootCmd'
var RootCmd = &cobra.Command{Use:"kvc"}
var createCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.MinimumNArgs(2),
	Short: "create key-value pair",
	Long:  "create key value strings into the key-value cache",
	RunE:   CommandRun.CreateCmd,
}

var readCmd = &cobra.Command{
	Use:  "read",
	Short: "read given key and return value",
	Args: cobra.MinimumNArgs(1),
	Long: "read value string out to command line from key-value cache given key string input from command line",
	RunE:  CommandRun.ReadCmd,
}

var updateCmd = &cobra.Command{
	Use:  "update",
	Args: cobra.MinimumNArgs(2),
	Short: "update key-value pair",
	Long:  "update key value strings into the key-value cache",
	RunE:  CommandRun.UpdateCmd,
}

var deleteCmd = &cobra.Command{
	Use:  "delete",
	Args: cobra.MinimumNArgs(1),
	Short: "delete key-value pair",
	Long:  "delete key value strings into the key-value cache",
	RunE: CommandRun.DeleteCmd,
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


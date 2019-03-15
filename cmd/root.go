package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	"CacheCLI/client"
)

//create CommandRunner to use ClientCache implementation of the KeyValueCache interface
var CommandRun = CommandRunner{
	cache: client.NewCacheClient(),
}

//require sub commands in CLI by not providing 'Run/RunE' in the definition of 'RootCmd'
var RootCmd = &cobra.Command{Use: "kvc"}

var createCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.ExactArgs(2),
	Short: "create key-value pair",
	Long:  "create key value strings into the key-value cache",
	RunE:  Create,
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read given key and return value",
	Args:  cobra.ExactArgs(1),
	Long:  "read value string out to command line from key-value cache given key string input from command line",
	RunE:  Read,
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Args:  cobra.ExactArgs(2),
	Short: "update key-value pair",
	Long:  "update key value strings into the key-value cache",
	RunE:  Update,
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Args:  cobra.ExactArgs(1),
	Short: "delete key-value pair",
	Long:  "delete key value strings into the key-value cache",
	RunE:  Delete,
}

// Execute appends sub commands to the root command and sets flags appropriately.
// This is called by main.main().
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(readCmd)
	RootCmd.AddCommand(updateCmd)
	RootCmd.AddCommand(deleteCmd)
	_ = RootCmd.Execute()
}

func Create(cmd *cobra.Command, args []string) error {
	res, err := CommandRun.CreateCmd(cmd, args)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func Update(cmd *cobra.Command, args []string) error {
	res, err := CommandRun.UpdateCmd(cmd, args)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func Read(cmd *cobra.Command, args []string) error {
	res, err := CommandRun.ReadCmd(cmd, args)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func Delete(cmd *cobra.Command, args []string) error {
	res, err := CommandRun.DeleteCmd(cmd, args)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

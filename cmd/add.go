/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	data = "data.json"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Password",
	Long:  `The command adds a password`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Adding Password")
		insertPassword()
	},
}

func insertPassword() {

	details := &Details{}

	if err := details.Load(data); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	data1 := userDetail{
		Username: "test",
		Email:    "1@2.com",
		Password: "asdf",
		Note:     "note",
	}
	details.add(data1)
	err := details.Store(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

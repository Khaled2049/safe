/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listPasswordsCmd represents the listPasswords command
var listPasswordsCmd = &cobra.Command{
	Use:   "listPasswords",
	Short: "List all store Passwords",
	Long:  `The command allows you to view Passwords that you have saved in the past`,
	Run: func(cmd *cobra.Command, args []string) {
		showList()
	},
}

func showList() {
	fmt.Println("Showing Passwords")
}

func init() {
	rootCmd.AddCommand(listPasswordsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listPasswordsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listPasswordsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

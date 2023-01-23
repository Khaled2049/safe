/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

const (
	userData = "user.json"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register as a user",
	Long:  `This command allows you to register as a user and use the password manager`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")

		users := &Users{}
		if err := users.Load(data); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		if username != "" && password != "" {

			user := userModel{
				Username: username,
				Password: hashAndSalt([]byte(password)),
			}

			users.add(user)
			e := users.Store(userData)
			if e != nil {
				fmt.Fprintln(os.Stderr, e.Error())
			}

		} else {
			fmt.Println("Let's get you registerd fam!")
			register()

		}
	},
}

func register() {
	users := &Users{}

	if err := users.Load(userData); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Print("Username: ")
	reader := bufio.NewReader(os.Stdin)

	username, err := reader.ReadString('\n')
	username = strings.TrimRight(username, "\r\n")
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	fmt.Print("Password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	user := userModel{
		Username: username,
		Password: hashAndSalt(password),
	}

	users.add(user)
	e := users.Store(userData)
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}

}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "Enter your username")
	registerCmd.Flags().StringP("password", "p", "", "Enter your password")
}

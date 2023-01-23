/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"
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
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		note, _ := cmd.Flags().GetString("note")

		details := &Details{}

		if err := details.Load(data); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}

		if username != "" && password != "" && email != "" && note != "" {
			key := []byte(getSecret())
			keyStr := hex.EncodeToString(key)
			hash := encrypt(keyStr, []byte(password))
			detail := userDetail{
				Username: username,
				Email:    email,
				Password: hash,
				Note:     note,
			}
			details.add(detail)
			e := details.Store(data)
			if e != nil {
				fmt.Fprintln(os.Stderr, e.Error())
			}
		} else {
			insertDetails()
		}
	},
}

func insertDetails() {

	details := &Details{}

	if err := details.Load(data); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	fmt.Print("Username: ")
	reader := bufio.NewReader(os.Stdin)
	// ReadString will block until the delimiter is entered
	username, err := reader.ReadString('\n')
	username = strings.TrimRight(username, "\r\n")
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	fmt.Print("Email: ")
	email, err := reader.ReadString('\n')
	email = strings.TrimRight(email, "\r\n")
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

	key := []byte(getSecret())
	keyStr := hex.EncodeToString(key)
	hash := encrypt(keyStr, password)

	fmt.Print("\nNote: ")
	note, err := reader.ReadString('\n')
	note = strings.TrimRight(note, "\r\n")
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	detail := userDetail{
		Username: username,
		Email:    email,
		Password: hash,
		Note:     note,
	}
	details.add(detail)
	e := details.Store(data)
	if e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
	}
}

func getSecret() string {
	config, err := LoadConfig(".")
	if err != nil {
		log.Fatalf("Can't load environment app.env: %v", err)
	}

	return config.SecretKey
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("username", "u", "", "Add the username")
	addCmd.Flags().StringP("password", "p", "", "Add the password")
	addCmd.Flags().StringP("email", "e", "", "Add the email")
	addCmd.Flags().StringP("note", "n", "", "Add the note")

}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var LOGGED_IN bool = false

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List passwords after logging in",
	Long: `Use this command to view the stored passwords
	Example: safe list -login=username
	`,
	Run: func(cmd *cobra.Command, args []string) {
		user, _ := cmd.Flags().GetString("login")
		if user != "" {
			if authenticate(user) {
				fmt.Println(`
				
				`)
				list()
			}
		} else {
			fmt.Println("Please Login to view passwords")
		}
	},
}

func list() {
	details := &Details{}
	if err := details.Load(data); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: gray("#")},
			{Align: simpletable.AlignCenter, Text: green("Username")},
			{Align: simpletable.AlignCenter, Text: green("Email")},
			{Align: simpletable.AlignCenter, Text: green("Password")},
			{Align: simpletable.AlignCenter, Text: green("Note")},
		},
	}

	var cells [][]*simpletable.Cell
	key := []byte(getSecret())
	keyStr := hex.EncodeToString(key)
	for i, data := range *details {
		i++

		pwd := decrypt(keyStr, data.Password)

		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i)},
			{Text: blue(data.Username)},
			{Text: blue(data.Email)},
			{Text: red(pwd)},
			{Text: blue(data.Note)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: "🔒 Password Manager 🚀"},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	fmt.Println(table.String())
}

func authenticate(user string) bool {
	users := &Users{}
	if err := users.Load(userData); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("Username: ", user)
	fmt.Print("Password: ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
	}

	for i, data := range *users {
		i++
		if data.Username == user {
			if comparePasswords(data.Password, password) {
				return true
			}
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(listCmd)
	// listCmd.PersistentFlags().String("login", "", "Authenticate to see passwords")
	listCmd.Flags().StringP("login", "l", "", "Login to view passwords")
	listCmd.MarkFlagRequired("login")
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

const (
	userData = "user.json"
)

type userModel struct {
	Username string
	Password string
}
type Users []userModel

func (u *Users) add(info userModel) {
	*u = append(*u, info)
	fmt.Println("Added successfully!")
}

func (u *Users) Delete(index int) error {
	ls := *u
	if index <= 0 || index > len(ls) {
		return errors.New("Invalid index")
	}

	*u = append(ls[:index-1], ls[index:]...)
	fmt.Println("Delete successful")
	return nil
}

func (u *Users) Load(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) Store(filename string) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register as a user",
	Long:  `This command allows you to register as a user to use the CLI password manager`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Registering user")
		register()
	},
}

func register() {
	users := &Users{}
	if err := users.Load(data); err != nil {
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
		fmt.Fprintln(os.Stderr, err.Error())
	}

}

func init() {
	rootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"fmt"
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
		fmt.Println("Enter yo details bruv")
		insertDetails()
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

	key := []byte("theultimatesupersecretpasswordis")
	// key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		panic(err.Error())
	}
	keyStr := hex.EncodeToString(key) //convert to string for saving
	fmt.Println("Encrypting.....")
	// encrypt value to base64
	hash := encrypt(keyStr, password)
	// hash := hashAndSalt(password)
	fmt.Println("hash: ", hash)

	// fmt.Println("Decrypting.....")
	// // encrypt base64 crypto to original value
	// text := decrypt(keyStr, hash)
	// fmt.Println(text)

	// test := comparePasswords(hash, password)
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

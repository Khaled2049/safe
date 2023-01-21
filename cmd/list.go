/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/alexeyco/simpletable"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List")
		list()
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
	for i, data := range *details {
		i++
		cells = append(cells, *&[]*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i)},
			{Text: blue(data.Username)},
			{Text: blue(data.Email)},
			{Text: red(data.Password)},
			{Text: blue(data.Note)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: "🔒"},
	}}

	table.SetStyle(simpletable.StyleMarkdown)
	fmt.Println(table.String())
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

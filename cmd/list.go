/*
Copyright © 2023 u-yas <yuya3.21998@gmail.com>
*/
package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/u-yas/ennbu/list"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all key-value pairs in .env file",
	Example: heredoc.Doc(`
	$ ennbu list
	$ ennbu list -e .env.development
	$ ennbu list --json
`),
	RunE: list.Run,
}

func init() {

	listCmd.Flags().BoolP(list.FlagJson, "j", false, "Output as json")
	rootCmd.AddCommand(listCmd)
}

/*
Copyright © 2023 u-yas <yuya3.21998@gmail.com>
*/
package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/u-yas/ennbu/flags"
	"github.com/u-yas/ennbu/get"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value of a specified key in .env file",
	Example: heredoc.Doc(`
$ ennbu get KEY
$ ennbu get -e .env.development
`),
	RunE: get.Run,
}

func init() {
	_ = getCmd.MarkFlagRequired(flags.FlagKey)

	rootCmd.AddCommand(getCmd)
}

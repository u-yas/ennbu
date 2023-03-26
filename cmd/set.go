/*
Copyright © 2023 u-yas <yuya3.21998@gmail.com>
*/
package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/u-yas/ennbu/flags"
	"github.com/u-yas/ennbu/set"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a value for a key in .env file",
	Example: heredoc.Doc(`
$ ennbu set -k KEY VALUE
$ ennbu set -e .env.development -k KEY VALUE
`),
	RunE: set.Run,
}

func init() {

	_ = setCmd.MarkFlagRequired(flags.FlagKey)
	rootCmd.AddCommand(setCmd)
}

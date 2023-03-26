/*
Copyright © 2023 u-yas <yuya3.21998@gmail.com>
*/
package cmd

import (
	"github.com/MakeNowJust/heredoc/v2"
	"github.com/spf13/cobra"
	"github.com/u-yas/ennbu/flags"
	"github.com/u-yas/ennbu/replace"
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replace a value of a key in .env file",
	Example: heredoc.Doc(`
$ ennbu replace -k KEY VALUE
$ ennbu replace -e .env.development -k KEY VALUE
`),
	RunE: replace.Run,
}

func init() {
	_ = replaceCmd.MarkFlagRequired(flags.FlagKey)

	rootCmd.AddCommand(replaceCmd)

}

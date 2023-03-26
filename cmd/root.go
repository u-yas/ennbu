/*
Copyright © 2023 u-yas:yuya3.21998@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/u-yas/ennbu/flags"
)

var rootCmd = &cobra.Command{
	Use:   "ennbu",
	Short: "ennbu is a CLI tool to manage .env files",
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP(flags.FlagEnvFile, "e", ".env", "Path to the .env file")
	rootCmd.PersistentFlags().StringP(flags.FlagKey, "k", "", "Key to the .env file")
}

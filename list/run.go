package list

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
	"github.com/u-yas/ennbu/env"
	"github.com/u-yas/ennbu/flags"
)

func Run(cmd *cobra.Command, args []string) error {
	envFilePath, _ := cmd.Flags().GetString(flags.FlagEnvFile)
	jsonOutput, _ := cmd.Flags().GetBool(FlagJson)

	e, err := env.ReadEnv(envFilePath)
	if err != nil {
		return fmt.Errorf("failed to read env file: %w", err)
	}

	if jsonOutput {
		var buffer bytes.Buffer
		enc := json.NewEncoder(&buffer)
		enc.SetEscapeHTML(false)
		enc.SetIndent("", "  ")

		if err := enc.Encode(e); err != nil {
			return fmt.Errorf("error generating JSON output: %w", err)
		}

		jVal := buffer.Bytes()
		if err != nil {
			return fmt.Errorf("error generating JSON output: %w", err)
		}

		fmt.Println(string(pretty.Color(jVal, nil)))
	} else {
		rawEnv, err := godotenv.Marshal(e)
		if err != nil {
			return fmt.Errorf("error generating raw output: %w", err)
		}
		fmt.Println(rawEnv)
	}

	return nil

}

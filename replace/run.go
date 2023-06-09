package replace

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/u-yas/ennbu/env"
	"github.com/u-yas/ennbu/flags"
)

func Run(cmd *cobra.Command, args []string) error {
	key, _ := cmd.Flags().GetString(flags.FlagKey)
	envFilePath, _ := cmd.Flags().GetString(flags.FlagEnvFile)
	escapeFlag, _ := cmd.Flags().GetBool(FlagEscape)
	importFileFlag, _ := cmd.Flags().GetBool(FlagImportFile)

	value := args[0]
	if importFileFlag {
		file, err := os.ReadFile(value)
		if err != nil {
			return fmt.Errorf("failed to read value from file: %w", err)
		}
		value = string(file)
	}
	if escapeFlag {
		value = env.EscapeSpecialChars(value)
	}

	e, err := env.ReadEnv(envFilePath)

	if err != nil {
		return fmt.Errorf("failed to read env file: %w", err)
	}

	// replace cmd should not add new key
	if _, ok := e[key]; !ok {
		return fmt.Errorf("key not found: %s", key)
	}

	e[key] = value

	err = env.WriteEnv(e, envFilePath)
	if err != nil {
		return fmt.Errorf("failed to write env file: %w", err)
	}
	return nil
}

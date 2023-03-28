package set

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/u-yas/ennbu/env"
	"github.com/u-yas/ennbu/flags"
)

func Run(cmd *cobra.Command, args []string) error {
	key, _ := cmd.Flags().GetString(flags.FlagKey)
	envFilePath, _ := cmd.Flags().GetString(flags.FlagEnvFile)
	escapeFlag, _ := cmd.Flags().GetBool(FlagEscape)

	value := args[0]
	if escapeFlag {
		value = env.EscapeSpecialChars(value)
	}

	e, err := env.ReadEnv(envFilePath)

	if err != nil {
		return fmt.Errorf("failed to read env file: %w", err)
	}

	e[key] = value

	err = env.WriteEnv(e, envFilePath)
	if err != nil {
		return fmt.Errorf("failed to write env file: %w", err)
	}

	return nil
}

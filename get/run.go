package get

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/u-yas/ennbu/env"
	"github.com/u-yas/ennbu/flags"
)

func Run(cmd *cobra.Command, args []string) error {
	key := args[0]
	envFilePath, _ := cmd.Flags().GetString(flags.FlagEnvFile)
	unescapeFlag, _ := cmd.Flags().GetBool(FlagUnescape)
	e, err := env.ReadEnv(envFilePath)
	if err != nil {
		return fmt.Errorf("failed to read env file: %w", err)
	}

	if _, ok := e[key]; !ok {
		return fmt.Errorf("key not found: %s", key)
	}

	value := e[key]
	if !unescapeFlag {
		value = strconv.Quote(value)
	}

	if strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"") {
		value = strings.TrimSuffix(strings.TrimPrefix(value, "\""), "\"")
	}

	fmt.Println(value)
	return nil
}

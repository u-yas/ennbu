package set

import (
	"fmt"
	"os"
	"strings"

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
	// Check if the value contains newline characters and wrap it with double quotes if necessary
	if strings.Contains(value, "\n") {
		value = fmt.Sprintf("\"%s\"", value)
	}

	envContent, err := os.ReadFile(envFilePath)
	if err != nil {
		return fmt.Errorf("error reading %s: %w", envFilePath, err)
	}

	lines := strings.Split(string(envContent), "\n")
	found := false

	for i, line := range lines {
		if strings.HasPrefix(line, key+"=") {
			lines[i] = key + "=" + value
			found = true
			break
		}
	}

	if !found {
		lines = append(lines, key+"="+value)
	}

	newContent := strings.Join(lines, "\n")
	err = os.WriteFile(envFilePath, []byte(newContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing %s: %w", envFilePath, err)
	}

	return nil
}

package set

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) error {
	key, _ := cmd.Flags().GetString("key")
	envFilePath, _ := cmd.Flags().GetString("envFile")
	value := args[0]

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

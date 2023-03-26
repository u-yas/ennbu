package get

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/u-yas/ennbu/flags"
)

func Run(cmd *cobra.Command, args []string) error {
	key := args[0]
	envFilePath, _ := cmd.Flags().GetString(flags.FlagEnvFile)

	envFile, err := os.Open(envFilePath)
	envFileName := filepath.Base(envFilePath)
	if err != nil {
		return fmt.Errorf("error opening %s file: %w", envFileName, err)
	}
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, key+"=") {
			value := strings.TrimPrefix(line, key+"=")
			fmt.Println(value)
			found = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading %s file: %w", envFileName, err)
	}

	if !found {
		return fmt.Errorf("key not found in %s file", envFileName)
	}

	return nil
}

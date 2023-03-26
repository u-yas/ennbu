package list

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tidwall/pretty"
	"github.com/u-yas/ennbu/flags"
)

func Run(cmd *cobra.Command, args []string) error {
	envFilePath, _ := cmd.Flags().GetString(flags.FlagEnvFile)
	jsonOutput, _ := cmd.Flags().GetBool(FlagJson)

	envfileName := filepath.Base(envFilePath)
	envContent, err := os.ReadFile(envFilePath)
	if err != nil {
		return fmt.Errorf("error reading %s file: %w", envfileName, err)
	}

	lines := strings.Split(string(envContent), "\n")
	envMap := make(map[string]string)

	for _, line := range lines {
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		kv := strings.SplitN(line, "=", 2)
		if len(kv) == 2 {
			envMap[kv[0]] = kv[1]
		}
	}

	if jsonOutput {
		jsonData, err := json.MarshalIndent(envMap, "", "  ")
		if err != nil {
			return fmt.Errorf("error generating JSON output: %w", err)
		}
		fmt.Println(string(pretty.Color(jsonData, nil)))
	} else {
		for key, value := range envMap {
			fmt.Printf("%s=%s\n", key, value)
		}
	}

	return nil

}

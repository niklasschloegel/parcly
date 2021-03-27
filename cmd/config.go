/*
Copyright © 2021 Niklas Schlögel <niklasschloegel@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Config command prints all available configuration variables.
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Information about config",
	Long: `Lists all available configuration variables for parcly
found in a config file or via environment variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		for key, val := range viper.AllSettings() {
			fmt.Printf("%s=%s\n", key, val)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}

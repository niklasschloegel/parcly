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
	"os"

	"github.com/niklasschloegel/parcly/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setApiKey string

// Sets a configuration variable and saves it to config file.
// Config file location is specified in root.go/initConfig
var setCmd = &cobra.Command{
	Use:   "set [-key=value | -key value]*",
	Short: "Sets a config variable",
	Long: `Sets a configuration variable with specified flag
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if setApiKey != "" {
			viper.Set(config.TrackTryConfigKey, setApiKey)
			viper.WriteConfig()
		} else {
			cmd.Help()
			os.Exit(0)
		}
	},
}

func init() {
	configCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().StringVar(&setApiKey, "tracktry", "", "Sets Tracktry API Key")
}

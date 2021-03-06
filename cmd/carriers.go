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
	"errors"

	"github.com/niklasschloegel/parcly/config"
	"github.com/spf13/cobra"
)

// Carrier noun which does not provide own functionality.
var carriersCmd = &cobra.Command{
	Use:   "carriers",
	Short: "Get informations about carriers",
	Long:  `Retreive information about available carriers.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if config.TracktryApiKey == "" {
			return errors.New(config.ErrorMsg)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(carriersCmd)
}

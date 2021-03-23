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
	"fmt"

	"github.com/niklasschloegel/parcly/config"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "parcly",
	Short: "Simply tracks parcels",
	Long:  `Parcly is a tool for tracking parcels.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if config.TracktryApiKey == "" {
			errorMsg := `API Key is missing.
You need to provide an API Key from Tracktry.
Sign up at https://www.tracktry.com/signup-en.html,
and copy the API key under 'Settings'.

The API Key can be provided in three ways:

1) As a flag:
	parcly <noun> <command> --tracktrykey <key>

2) As an environment variable:
	export PARCLY_TRACKTRYKEY=<key>

3) In a config file:
	default config file is $HOME/.parcly.yaml and should contain:
	tracktrykey: <key>

	When you want to use another location, you can
	specify the location with another flag:
	parcly ... --config <filepath>
			`
			return errors.New(errorMsg)
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.parcly.yaml)")
	rootCmd.PersistentFlags().StringVar(&config.TracktryApiKey, "tracktrykey", "", "Tracktry API Key")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".parcly" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".parcly")
	}

	viper.SetEnvPrefix("parcly")
	viper.AutomaticEnv() // read in environment variables that match
	config.TracktryApiKey = viper.GetString("tracktrykey")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Reading env file")
		config.TracktryApiKey = viper.GetString("tracktrykey")
	}
}

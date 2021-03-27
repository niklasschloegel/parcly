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
	"log"
	"strings"

	"github.com/niklasschloegel/parcly/tracking"
	"github.com/spf13/cobra"
)

// Searches for carriers.
var carrierSearchCmd = &cobra.Command{
	Use:   "search [substr1 ... substrN]",
	Short: "Searches for carriers",
	Long:  `Searches through all available carriers matching substr`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			filteredCarriers := []tracking.Carrier{}
			allCarriers := tracking.GetCarriers()
			for _, c := range allCarriers {
				for _, substr := range args {
					if strings.Contains(
						strings.ToLower(c.Name),
						strings.ToLower(substr),
					) {
						filteredCarriers = append(filteredCarriers, c)
						continue
					}
				}
			}

			for _, c := range filteredCarriers {
				fmt.Println(c.Info())
			}

		} else {
			log.Fatal("No search arguments given")
		}
	},
}

func init() {
	carriersCmd.AddCommand(carrierSearchCmd)
}

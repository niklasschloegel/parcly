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

	"github.com/niklasschloegel/parcly/tracking"
	"github.com/spf13/cobra"
)

// Tries to detect carrier from specified tracking number.
var detectCmd = &cobra.Command{
	Use:   "detect [trackingNr]",
	Short: "Detects carrier",
	Long:  `Tries to detect carrier from tracking number`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			detectedCarriers := tracking.Detect(args[0])
			if len(detectedCarriers) > 0 {
				fmt.Println("Following carriers are a possible match:")
				for _, c := range detectedCarriers {
					fmt.Println(c.Info())
				}
			} else {
				fmt.Println("No matching carrier found for tracking number:", args[0])
			}
		} else {
			log.Fatal("No or too many tracking number(s) provided")
		}
	},
}

func init() {
	carriersCmd.AddCommand(detectCmd)
}

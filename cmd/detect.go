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

// detectCmd represents the detect command
var detectCmd = &cobra.Command{
	Use:   "detect [trackingNr]",
	Short: "Detects carrier",
	Long:  `Tries to detect carrier from tracking number`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			detectedCarriers := tracking.Detect(args[0])
			fmt.Println("Following carriers are a possible match:")
			for _, c := range detectedCarriers {
				fmt.Println(c.Info())
			}
		} else {
			log.Fatal("No or too many tracking number(s) provided")
		}
	},
}

func init() {
	carriersCmd.AddCommand(detectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// detectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// detectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

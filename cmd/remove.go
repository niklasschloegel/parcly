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

	"github.com/niklasschloegel/parcly/tracking"
	"github.com/spf13/cobra"
)

var carrier string

var removeCmd = &cobra.Command{
	Use:     "remove <trackingNr>",
	Short:   "Removes tracking",
	Long:    `Removes tracking item from specified tracking nr`,
	Aliases: []string{"rm"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			trackingNr := args[0]
			if carrier == "" {
				allTrackings := tracking.GetTrackings()
				for _, t := range allTrackings {
					if t.TrackingNumber == trackingNr {
						carrier = t.CarrierCode
						break
					}
				}
				if carrier == "" {
					fmt.Printf("No tracking with tracking number '%s' found.\n", trackingNr)
					return
				}
			}
			tracking.DeleteTracking(carrier, trackingNr)
			fmt.Println("Successfully deleted tracking item.")
		} else {
			fmt.Println("No or too many tracking numbers provided.")
		}
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
	removeCmd.PersistentFlags().StringVarP(&carrierCode, "carrier", "c", "", "Defines carrier of parcel")
}

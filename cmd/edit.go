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

var editCarrierCode, editTitle string

var editCmd = &cobra.Command{
	Use:   "edit <trackingNr>",
	Short: "Edits tracking item",
	Long:  `Edits tracking item with specified tracking number`,
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
			updateRequest := tracking.UpdateTracking{}
			if editTitle != "" {
				updateRequest.Title = editTitle
			}

			if updateRequest != (tracking.UpdateTracking{}) {
				// if struct not empty
				tracking.EditTrackings(carrier, trackingNr, updateRequest)
				fmt.Println("Successfully edited tracking item.")
			} else {
				fmt.Println("Nothing to change.")
			}
		} else {
			fmt.Println("No or too many tracking numbers provided.")
		}
	},
}

func init() {
	trackingCmd.AddCommand(editCmd)
	editCmd.PersistentFlags().StringVarP(&editCarrierCode, "carrier", "c", "", "Defines carrier of parcel to edit")
	editCmd.PersistentFlags().StringVarP(&editTitle, "title", "t", "", "Sets new title")
}

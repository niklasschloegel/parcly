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

var carriersListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists available carriers",
	Long:  `Lists all available carriers.`,
	Run: func(cmd *cobra.Command, args []string) {

		carriers := tracking.GetCarriers()
		for _, c := range carriers {
			fmt.Println(c.Info())
		}
	},
}

var trackingDetail bool

var trackingListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all trackings",
	Long:  `Lists all registered tracking items.`,
	Run: func(cmd *cobra.Command, args []string) {
		allTrackings := tracking.GetTrackings()
		if len(allTrackings) > 0 {
			for _, t := range allTrackings {
				if trackingDetail {
					fmt.Println(t.Info())
				} else {
					fmt.Println(t.ShortInfo())
				}
			}
		} else {
			fmt.Println("No parcels tracked.")
		}
	},
}

func init() {
	carriersCmd.AddCommand(carriersListCmd)
	trackingCmd.AddCommand(trackingListCmd)

	trackingListCmd.PersistentFlags().BoolVarP(&trackingDetail, "detail", "d", false,
		"Specifies if previous tracking status should get shown.")
}

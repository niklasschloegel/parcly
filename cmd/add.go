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

var carrierCode, shipDate, postalCode, specNumberDest, orderID, title string

// Adds a tracking item to Tracktry.
// Prints short tracking information after successful creatin.
var addCmd = &cobra.Command{
	Use:   "add <trackingNr>",
	Short: "Tracks a parcel",
	Long: `Adds a parcel to tracking items.

If no carrier defined (see flags), parcly tries 
to detect the carrier and lets you choose the
best matching carrier.
To find out which carrier code applies to your parcel 
separately you can use the 'detect' command. 
	-> parcly help carriers detect
	
If you know the name of the carrier you can search 
for the specific carrier code
	-> parcly help carriers search
	`,
	Aliases: []string{"track"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			trackingNr := args[0]

			trackingCreation := tracking.TrackingCreation{TrackingNumber: trackingNr}
			addOptions(&trackingCreation)

			if carrierCode != "" {
				trackingCreation.CarrierCode = carrierCode
			} else {
				possibleCarriers := tracking.Detect(trackingNr)
				if len(possibleCarriers) > 0 {
					fmt.Println("Following carriers are a possible match:")
					for _, c := range possibleCarriers {
						fmt.Println(c.Info())
					}
				} else {
					fmt.Println("No matching carrier found for tracking number", trackingNr)
					return
				}
				fmt.Print("Please enter the applicable carrier code: ")
				var code string
				_, err := fmt.Scan(&code)
				if err != nil {
					log.Fatal("Error reading carrier code")
					return
				}
				trackingCreation.CarrierCode = code
			}
			resp := tracking.CreateTracking(trackingCreation)
			fmt.Println(resp.Info())
		} else {
			log.Fatal("No or too many tracking number(s) provided.")
		}
	},
}

func init() {
	trackingCmd.AddCommand(addCmd)

	addCmd.PersistentFlags().StringVarP(&carrierCode, "carrier", "c", "", "Defines carrier of parcel")
	addCmd.PersistentFlags().StringVar(&shipDate, "shipdate", "", "Shipping date in YYYYMMDD format. Required by some couriers, such as deutsche-post")
	addCmd.PersistentFlags().StringVar(&postalCode, "postalcode", "", "Postal code of receiver's address. Required by some couriers, such as postnl-3s")
	addCmd.PersistentFlags().StringVar(&specNumberDest, "special", "", "Destination Country of the shipment for a specific couriers, such as postnl-3s")
	addCmd.PersistentFlags().StringVarP(&orderID, "orderid", "o", "", "Order ID of parcel")
	addCmd.PersistentFlags().StringVarP(&title, "title", "t", "", "Title for parcel (e.g. content of parcel)")

}

func addOptions(t *tracking.TrackingCreation) {
	if shipDate != "" {
		t.TrackingShipDate = shipDate
	}
	if postalCode != "" {
		t.TrackingPostalCode = postalCode
	}
	if specNumberDest != "" {
		t.SpecialNumberDestination = specNumberDest
	}
	if orderID != "" {
		t.Order = orderID
	}
	if title != "" {
		t.Comment = title
	}
}

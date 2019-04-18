// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/spf13/cobra"
	gcalendar "github.com/spf13/my-next-meeting/lib"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := ioutil.ReadFile("calendar.config")
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		// If modifying these scopes, delete your previously saved token.json.
		config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope, "https://www.googleapis.com/auth/userinfo.email")
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}

		client := gcalendar.GetClient(config)
		user := gcalendar.GetLoggedInUser(client)

		srv, err := calendar.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve Calendar client: %v", err)
		}

		t := time.Now().Format(time.RFC3339)
		events, err := srv.Events.List("primary").ShowDeleted(false).
			SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
		if err != nil {
			log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
		}
		if len(events.Items) == 0 {
			fmt.Println("No upcoming events found to cancel.")
		} else {
			item := events.Items[0]
			att := findAttendee(item.Attendees, user.Email)
			att.ResponseStatus = "declined"

			_, err := srv.Events.Patch("primary", item.Id, item).Do()
			if err != nil {
				log.Fatalf("Unable to cancel event: %v", err)
			}
		}
	},
}

func findAttendee(attendees []*calendar.EventAttendee, email string) (ret *calendar.EventAttendee) {
	for _, att := range attendees {
		if att.Email == email {
			return att
		}
	}
	return
}

func init() {
	rootCmd.AddCommand(cancelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cancelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cancelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

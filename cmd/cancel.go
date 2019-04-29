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

	"github.com/spf13/cobra"
	gcalendar "github.com/spf13/my-next-meeting/lib"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel your next Google Calendar meeting",
	Long:  `Using the Google Calendar API cancel your next meeting optionally providing a related comment.`,
	Run: func(cmd *cobra.Command, args []string) {
		const userInfoScope = "https://www.googleapis.com/auth/userinfo.email"

		b, err := ioutil.ReadFile("calendar.config")
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		// get config for google client
		config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope, userInfoScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}

		client := gcalendar.GetClient(config)

		user := gcalendar.GetLoggedInUser(client)

		srv, err := calendar.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve Calendar client: %v", err)
		}

		event := gcalendar.GetNextEvent(srv)
		if event == nil {
			fmt.Println("No upcoming events found to cancel.")
		} else {
			att := findAttendee(event.Attendees, user.Email)
			att.ResponseStatus = "declined"

			comment, _ := cmd.Flags().GetString("comment")
			if comment != "" {
				att.Comment = comment
			}
			_, err := srv.Events.Patch("primary", event.Id, event).Do()
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
	cancelCmd.Flags().StringP("comment", "c", "", "Add a comment in your response")
}

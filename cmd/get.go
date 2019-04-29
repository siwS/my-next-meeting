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
	calendar "google.golang.org/api/calendar/v3"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get your next Google Calendar meeting",
	Long:  `Using the Google Calendar API get your next meeting and all its details: Summary, Details, Date, Attendees, Location, Hangout Link`,
	Run: func(cmd *cobra.Command, args []string) {
		b, err := ioutil.ReadFile("calendar.config")
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		config, err := google.ConfigFromJSON(b, calendar.CalendarEventsScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}
		client := gcalendar.GetClient(config)

		srv, err := calendar.New(client)
		if err != nil {
			log.Fatalf("Unable to retrieve Calendar client: %v", err)
		}

		event := gcalendar.GetNextEvent(srv)
		if event == nil {
			fmt.Println("No upcoming events found.")
		} else {
			printEventDetails(event)
		}
	},
}

func printEventDetails(event *calendar.Event) {
	fmt.Println("Your next meeting is:")

	date := event.Start.DateTime
	if date == "" {
		date = event.Start.Date
	}

	fmt.Printf("'%v' and is scheduled for %v\n", event.Summary, parseAndFormatDate(date))

	if event.Description != "" {
		fmt.Printf("%v \n", event.Description)
	}

	fmt.Printf("Organised by: %v\n", event.Organizer.Email)

	fmt.Printf("Attendees:")
	for _, attendee := range event.Attendees {
		fmt.Printf("%v \n", attendee.Email)
	}

	if event.HangoutLink != "" {
		fmt.Printf("You can dial in here: %v\n", event.HangoutLink)
	}

	if event.Location != "" {
		fmt.Printf("Event location is: %v\n", event.Location)
	}
}

func parseAndFormatDate(date string) string {
	layout := "2006-01-02T15:04:05+01:00"
	// https://flaviocopes.com/go-date-time-format/
	printLayout := "2006-01-02 15:04"

	dateParsed, err := time.Parse(layout, date)
	if err != nil {
		log.Fatalf("Unable to parse date correctly: %v", err)
	}

	dateFormatted := dateParsed.Format(printLayout)
	return dateFormatted
}

func init() {
	rootCmd.AddCommand(getCmd)
}

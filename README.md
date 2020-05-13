# my-next-meeting

A simple CLI written in Go to get your next meeting from Google calendar.

The CLI was built using [Cobra](https://github.com/spf13/cobra) and the [Google Calendar API](https://godoc.org/google.golang.org/api/calendar/v3)


## Installation

To install locally: 

1. Clone this repo in your $GOPATH

2. Install the Cobra library

`➜  go get -u github.com/spf13/cobra/cobra`

3. Install the Google Calendar API

`➜  go get google.golang.org/api/calendar/v3`

4. Navigate to the project directory

`cd my-next-meeting`

5. Create a token.json file

You need to create a `token.json` file in the project root directory with your Google Access and Refresh Token.\
Those are needed for the Google OAuth flow.\
For information on how to get the keys, check the [Google Calendar Quickstart Guide](https://developers.google.com/calendar/quickstart/go).\
*Note: do not share, or commit those on Github repos, keep them safe.*

6. Build the project

`go install main.go`

7. Run the commands to get or cancel your next meeting:

`my-next-meeting get`

`my-next-meeting cancel --comment "Sorry, I can't make it."`

## Usage

Usage: `my-next-meeting [command]`


Available Commands:
-  cancel      Cancel your next Google Calendar meeting
-  get         Get your next Google Calendar meeting
-  help        Help about any command

### TODO:

- [ ] Allow getting or cancelling more than the next meeting passing a flag. 
- [ ] Write tests.

## Contributing

Bug reports and pull requests are welcome. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [Contributor Covenant](http://contributor-covenant.org) code of conduct.

## License

The gem is available as open source under the terms of the [MIT License](https://opensource.org/licenses/MIT).





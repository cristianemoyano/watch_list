package constants

import (
	"github.com/common-nighthawk/go-figure"
)

var AddUsage = `
NAME:
   add - Add content to the watchlist by Name.

USAGE:
	add Name tipo
`

var SelectUsage = `
NAME:
   select - Add content to the watchlist by Id.

USAGE:
	select id
`

var ChangeUsage = `
NAME:
   change - Change watchlist's content

USAGE:
   change id Name tipo
`

var RemoveUsage = `
NAME:
   remove - Remove content by ID from the watchlist

USAGE:
   remove id
`

var ListUsage = `
NAME:
   list - List watchlist's content

USAGE:
   list
`

var HomeUsage = `
NAME:
   home - List movies, series and documentaries.

USAGE:
   home
`

var SearchUsage = `
NAME:
   search - Advanced search content

USAGE:
   search <content's name / category >
`

var LoginUsage = `
NAME:
   login - Login into the platform.

USAGE:
   login username
`
var ExitMsg = "Goodbye"

var DefaultMsg = "Invalid command:  '%v' - type> help \n"

var WelcomeMsg = "Welcome to: "

var WelcomeFigure = figure.NewFigure("Watchlist CLI", "ogre", true)

var ListAvailableCmds = `
COMMANDS:

   login       Login into the platform.                 | login username
   add         Add content to the watchlist by Name.    | add Name tipo
   search      Advanced search content                  | search <content's name / category >
   select      Add content to the watchlist by ID.      | select id
   change      Change watchlist's content.              | change id Name tipo
   remove      Remove content by ID from the watchlist. | remove id
   list        List watchlist's content.                | list
   home        List movies, series and documentaries.   | home

   help        List available commands.                 | help
   exit        Finish program.                          | exit
`

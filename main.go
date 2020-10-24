package main

import (
	"bufio"
	"os"
	"strings"

	commands "github.com/cristianemoyano/watch_list/commands"
	constants "github.com/cristianemoyano/watch_list/constants"
	watchlist "github.com/cristianemoyano/watch_list/watch_list"
	"github.com/fatih/color"
)

// Instantiate ContentLists
var watch = &watchlist.ContentList{}
var movies = &watchlist.ContentList{}

// User
var user = &watchlist.User{
	ID:       0,
	Username: "",
}

// Loggers
var red = color.New(color.FgRed)
var msg = color.New(color.FgCyan).Add(color.Bold)

func set_content(movies **watchlist.ContentList) {
	monterInc := watchlist.Content{
		ID:    commands.GenerateID(),
		Title: "Monster Inc.",
		Tipo:  watchlist.Pelicula,
	}
	friends := watchlist.Content{
		ID:    commands.GenerateID(),
		Title: "Friends",
		Tipo:  watchlist.Serie,
	}
	asado := watchlist.Content{
		ID:    commands.GenerateID(),
		Title: "Todo sobre el asado",
		Tipo:  watchlist.Documental,
	}
	(**movies).Add(monterInc)
	(**movies).Add(friends)
	(**movies).Add(asado)
}

func init() {
	// Welcome msg
	color.Cyan(constants.WelcomeMsg)
	constants.WelcomeFigure.Print()
	color.Cyan(constants.ListAvailableCmds)
	// Set movies
	set_content(&movies)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		c := color.New(color.FgCyan)
		c.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err := recover(); err != nil {
			red.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			red.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "login":
		defer func() {
			if err := recover(); err != nil {
				red.Println("Error:", err)
				msg.Println(constants.LoginUsage)
			}
		}()
		commands.Login(&user, args)
	case "add":
		defer func() {
			if err := recover(); err != nil {
				red.Println("Error:", err)
				msg.Println(constants.AddUsage)
			}
		}()
		commands.Add(&user, &watch, args)
	case "select":
		defer func() {
			if err := recover(); err != nil {
				red.Println("Error:", err)
				msg.Println(constants.SelectUsage)
			}
		}()
		commands.Select(&user, &watch, &movies, args)
	case "search":
		defer func() {
			if err := recover(); err != nil {
				red.Println("Error:", err)
				msg.Println(constants.SearchUsage)
			}
		}()
		commands.Search(&user, &movies, args)
	case "change":
		defer func() {
			if err := recover(); err != nil {
				red.Println("Error:", err)
				msg.Println(constants.ChangeUsage)
			}
		}()
		commands.Change(&user, &watch, args)
	case "remove":
		defer func() {
			if err := recover(); err != nil {
				red.Println("Error:", err)
				msg.Println(constants.RemoveUsage)
			}
		}()
		commands.Remove(&user, &watch, args)
	case "list":
		defer func() {
			if err := recover(); err != nil {
				red.Println("Error:", err)
				msg.Println(constants.ListUsage)
			}
		}()
		commands.List(&user, &watch)
	case "movies":
		defer func() {
			if err := recover(); err != nil {
				red.Println("Error:", err)
				msg.Println(constants.ListUsage)
			}
		}()
		commands.List(&user, &movies)
	case "help":
		msg.Println(constants.ListAvailableCmds)
	case "exit":
		os.Exit(0)
	default:
		red.Printf("Invalid command:  '%v' - type> help \n", args[0])
	}
	return nil
}

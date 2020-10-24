package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	commands "github.com/cristianemoyano/watch_list/commands"
	constants "github.com/cristianemoyano/watch_list/constants"
	watchlist "github.com/cristianemoyano/watch_list/watch_list"
)

// Instantiate ContentList
var watch = &watchlist.ContentList{}
var movies = &watchlist.ContentList{}

func set_default_list(watch **watchlist.ContentList) {
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
	(**watch).Add(monterInc)
	(**watch).Add(friends)
	(**watch).Add(asado)
}

func init() {
	// Welcome msg
	fmt.Println(constants.WelcomeMsg)
	constants.WelcomeFigure.Print()
	fmt.Println(constants.ListAvailableCmds)
	// Set a default content list
	set_default_list(&movies)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
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
	case "add":
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error:", err)
				log.Println(constants.AddUsage)
			}
		}()
		commands.Add(&watch, args)
	case "select":
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error:", err)
				log.Println(constants.SelectUsage)
			}
		}()
		commands.Select(&watch, &movies, args)
	case "change":
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error:", err)
				log.Println(constants.ChangeUsage)
			}
		}()
		commands.Change(&watch, args)
	case "remove":
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error:", err)
				log.Println(constants.RemoveUsage)
			}
		}()
		commands.Remove(&watch, args)
	case "list":
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error:", err)
				log.Println(constants.ListUsage)
			}
		}()
		commands.List(&watch)
	case "movies":
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error:", err)
				log.Println(constants.ListUsage)
			}
		}()
		commands.List(&movies)
	case "help":
		fmt.Println(constants.ListAvailableCmds)
	case "exit":
		os.Exit(0)
	default:
		fmt.Printf("Invalid command:  '%v' - type> help \n", args[0])
	}
	return nil
}

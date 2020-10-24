package commands

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cristianemoyano/watch_list/constants"
	watchlist "github.com/cristianemoyano/watch_list/watch_list"
	"github.com/fatih/color"
	"github.com/sony/sonyflake"
)

var flake = sonyflake.NewSonyflake(sonyflake.Settings{})

func GenerateID() uint64 {

	id, err := flake.NextID()
	if err != nil {
		log.Fatalf("flake.NextID() failed with %s\n", err)
	}
	return id
}
func is_authenticated(user **watchlist.User) {
	if (**user).Username == "" {
		panic(errors.New("You must to be logged in."))
	}
}

func Add(user **watchlist.User, watch **watchlist.ContentList, arguments []string) {
	is_authenticated(user)

	name := arguments[1]
	categorySelected := arguments[2]

	newContent := watchlist.Content{
		ID:    GenerateID(),
		Title: name,
		Tipo:  watchlist.GetCategory(categorySelected),
	}
	(**watch).Add(newContent)
	(**watch).Print()
	fmt.Printf("---Added by Name---\n")
}

func Login(user **watchlist.User, arguments []string) {
	username := arguments[1]
	id := GenerateID()
	(**user).Login(username, id)
	(**user).Print()
	fmt.Printf("---Welcome %v !---\n", username)
}

func Select(user **watchlist.User, watch **watchlist.ContentList, movies **watchlist.ContentList, arguments []string) {
	is_authenticated(user)
	id, _ := strconv.ParseUint(arguments[1], 10, 64)

	content := (**movies).Search(id)

	newContent := watchlist.Content{
		ID:    content.ID,
		Title: content.Title,
		Tipo:  content.Tipo,
	}
	(**watch).Add(newContent)
	(**watch).Print()
	fmt.Printf("---Added by ID---\n")
}

func Change(user **watchlist.User, watch **watchlist.ContentList, arguments []string) {
	is_authenticated(user)

	id, _ := strconv.ParseUint(arguments[1], 10, 64)
	name := arguments[2]
	tipoSelected := arguments[3]

	tipo := watchlist.Pelicula
	switch tipoSelected {
	case "pelicula":
		tipo = watchlist.Pelicula
	case "serie":
		tipo = watchlist.Serie
	case "documental":
		tipo = watchlist.Documental
	}
	newContent := watchlist.Content{
		ID:    GenerateID(),
		Title: name,
		Tipo:  tipo,
	}
	(**watch).Change(id, newContent)
	(**watch).Print()
	fmt.Printf("---Changed---\n")
}

func Remove(user **watchlist.User, watch **watchlist.ContentList, arguments []string) {
	is_authenticated(user)

	id, _ := strconv.ParseUint(arguments[1], 10, 64)
	fmt.Printf("ID to remove: %v \n", id)
	(**watch).Remove(id)
	(**watch).Print()
	fmt.Printf("---Removed---\n")
}

func List(user **watchlist.User, watch **watchlist.ContentList) {
	is_authenticated(user)
	(**watch).Print()
}

func Home(logger *color.Color, user **watchlist.User, movies **watchlist.ContentList) {
	is_authenticated(user)

	logger.Println("\n Movies")
	(**movies).SearchAdvanced("pelicula")

	logger.Println("\n Series")
	(**movies).SearchAdvanced("serie")

	logger.Println("\n Documentaries")
	(**movies).SearchAdvanced("documental")
}

func Search(user **watchlist.User, movies **watchlist.ContentList, arguments []string) {
	is_authenticated(user)
	search := arguments[1]
	(**movies).SearchAdvanced(search)
}

func Help(logger *color.Color) {
	// No authenticated command
	logger.Println(constants.ListAvailableCmds)
}

func Exit(logger *color.Color, user **watchlist.User) {
	// No authenticated command
	logger.Println(constants.ExitMsg, (**user).Username)
	os.Exit(0)
}

func Default(logger *color.Color, args []string) {
	// No authenticated command
	logger.Printf(constants.DefaultMsg, args[0])
}

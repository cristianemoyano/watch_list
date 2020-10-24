package commands

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	watchlist "github.com/cristianemoyano/watch_list/watch_list"
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

func Add(user **watchlist.User, watch **watchlist.ContentList, arguments []string) {
	if (**user).Username == "" {
		panic(errors.New("You must to be logged in."))
	}

	name := arguments[1]
	tipoSelected := arguments[2]

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
	if (**user).Username == "" {
		panic(errors.New("You must to be logged in."))
	}

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
	if (**user).Username == "" {
		panic(errors.New("You must to be logged in."))
	}

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
	if (**user).Username == "" {
		panic(errors.New("You must to be logged in."))
	}

	id, _ := strconv.ParseUint(arguments[1], 10, 64)
	fmt.Printf("ID to remove: %v \n", id)
	(**watch).Remove(id)
	(**watch).Print()
	fmt.Printf("---Removed---\n")
}

func List(user **watchlist.User, watch **watchlist.ContentList) {
	if (**user).Username == "" {
		panic(errors.New("You must to be logged in."))
	}
	(**watch).Print()
}

func Search(user **watchlist.User, movies **watchlist.ContentList, arguments []string) {
	if (**user).Username == "" {
		panic(errors.New("You must to be logged in."))
	}

	name := arguments[1]
	content := (**movies).SearchName(name)

	content.Print()
}

package commands

import (
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

func Add(watch **watchlist.ContentList, arguments []string) {
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

func Select(watch **watchlist.ContentList, movies **watchlist.ContentList, arguments []string) {
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

func Change(watch **watchlist.ContentList, arguments []string) {
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

func Remove(watch **watchlist.ContentList, arguments []string) {
	id, _ := strconv.ParseUint(arguments[1], 10, 64)
	fmt.Printf("ID to remove: %v \n", id)
	(**watch).Remove(id)
	(**watch).Print()
	fmt.Printf("---Removed---\n")
}

func List(watch **watchlist.ContentList) {
	(**watch).Print()
}

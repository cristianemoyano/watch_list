package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"

	watchlist "github.com/cristianemoyano/watch_list/watch_list"
)

var app = cli.NewApp()
var pizza = []string{"Enjoy your pizza with some delicious"}
var watch = &watchlist.ContentList{}

func prelisted() {
	monterInc := watchlist.Content{
		ID:    1,
		Title: "Monster Inc.",
		Tipo:  watchlist.Pelicula,
	}
	friends := watchlist.Content{
		ID:    2,
		Title: "Friends",
		Tipo:  watchlist.Serie,
	}
	asado := watchlist.Content{
		ID:    3,
		Title: "Todo sobre el asado",
		Tipo:  watchlist.Documental,
	}
	watch.Add(monterInc)
	watch.Add(friends)
	watch.Add(asado)
}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add content to the watchlist: ./main add 'Name' 'tipo' ",
			Action: func(c *cli.Context) error {
				name := c.Args().Get(0)
				tipoSelected := c.Args().Get(1)
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
					ID:    4,
					Title: name,
					Tipo:  tipo,
				}
				watch.Add(newContent)
				watch.Print()
				fmt.Printf("---Added---\n")
				return nil
			},
		},
		{
			Name:    "change",
			Aliases: []string{"c"},
			Usage:   "Change watchlist's content: ./main change id 'Name' 'tipo'",
			Action: func(c *cli.Context) error {
				id, _ := strconv.Atoi(c.Args().Get(0))
				name := c.Args().Get(1)
				tipoSelected := c.Args().Get(2)
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
					ID:    id,
					Title: name,
					Tipo:  tipo,
				}
				watch.Change(id, newContent)
				watch.Print()
				fmt.Printf("---Changed---\n")
				return nil
			},
		},
		{
			Name:    "remove",
			Aliases: []string{"r"},
			Usage:   "Remove content by ID from the watchlist: ./main remove id",
			Action: func(c *cli.Context) error {
				id, _ := strconv.Atoi(c.Args().Get(0))
				watch.Remove(id)
				watch.Print()
				fmt.Printf("Remove ID %v \n", id)
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List watchlist's content: ./main list",
			Action: func(c *cli.Context) error {
				watch.Print()
				return nil
			},
		},
	}
}

func main() {

	app.Name = "Watchlist CLI"
	app.Usage = "An example CLI for watchlist"
	app.Version = "1.0.0"

	commands()
	prelisted()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
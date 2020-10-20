/*
Exercise:

Develop a program to manage content that you want to watch like in Netflix.
The program should works as the "Watch list" that the platform has.

Each content (node) must save the id, the name and the type (series, film or documentary).

Functions to build: Add, Remove and Change.
The program must have a menu to perform these functions.
*/
package watchlist

import (
	"errors"
	"fmt"
)

// ContentType Enum
type ContentType string

const (
	// Documental type
	Documental ContentType = "documental"
	// Pelicula type
	Pelicula ContentType = "pelicula"
	// Serie type
	Serie ContentType = "serie"
)

// Content struct
type Content struct {
	ID    int
	Title string
	Tipo  ContentType
	next  *Content
}

// ContentList struct
type ContentList struct {
	Length int
	start  *Content
}

// Add method
func (n *ContentList) Add(newContent Content) {
	if n.Length == 0 {
		n.start = &newContent
	} else {
		currentContent := n.start
		for currentContent.next != nil {
			currentContent = currentContent.next
		}
		currentContent.next = &newContent
	}
	n.Length++
}

// Remove method
func (n *ContentList) Remove(id int) {
	if n.Length == 0 {
		panic(errors.New("Content list is empty"))
	}

	var previousContent *Content
	currentContent := n.start

	for currentContent.ID != id {
		if currentContent.next == nil {
			panic(errors.New("Content id not found"))
		}

		previousContent = currentContent
		currentContent = currentContent.next
	}
	previousContent.next = currentContent.next

	n.Length--
}

// Change method
func (n *ContentList) Change(id int, newContent Content) {
	if n.Length == 0 {
		panic(errors.New("Content list is empty"))
	}
	n.Add(newContent)
	n.Remove(id)
}

// Print method
func (n *ContentList) Print() {
	if n.Length == 0 {
		panic(errors.New("Content list is empty"))
	}
	currentContent := n.start
	for {
		fmt.Printf("#%v - %v - type: %v\n", currentContent.ID, currentContent.Title, currentContent.Tipo)
		if currentContent.next == nil {
			break
		}
		currentContent = currentContent.next
	}
}

// Search method
func (n *ContentList) Search(id int) *Content {
	if n.Length == 0 {
		panic(errors.New("Content list is empty"))
	}
	currentContent := n.start
	for {
		if currentContent.ID == id {
			fmt.Printf("#%v - %v - type: %v\n", currentContent.ID, currentContent.Title, currentContent.Tipo)
			return currentContent
		}
		if currentContent.next == nil {
			break
		}
		currentContent = currentContent.next
	}
	return currentContent
}

func demo() {
	n := &ContentList{}
	fmt.Printf("---add---\n")
	monterInc := Content{
		ID:    1,
		Title: "Monster Inc.",
		Tipo:  Pelicula,
	}
	friends := Content{
		ID:    2,
		Title: "Friends",
		Tipo:  Serie,
	}
	asado := Content{
		ID:    3,
		Title: "Todo sobre el asado",
		Tipo:  Documental,
	}
	n.Add(monterInc)
	n.Add(friends)
	n.Add(asado)
	n.Print()
	fmt.Printf("---remove: 3---\n")
	n.Remove(3)
	n.Print()
	fmt.Printf("---search: 2---\n")
	n.Search(2)
	fmt.Printf("---change: 2---\n")
	bigBang := Content{
		ID:    4,
		Title: "The Big Bang Theory",
		Tipo:  Serie,
	}
	n.Change(2, bigBang)
	n.Print()
}

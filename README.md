# Watchlist CLI

A watchlist using linked lists in Go.


# Get started

```
git clone https://github.com/cristianemoyano/watch_list.git

cd watch_list

go build main.go

./main [command options]

or

go run main.go [command options]
```

# Commands

```
go run main.go help
NAME:
   Watchlist CLI - An example CLI for watchlist

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   add, a     Add content to the watchlist: ./main add 'Name' 'tipo' 
   change, c  Change watchlist's content: ./main change id 'Name' 'tipo'
   remove, r  Remove content by ID from the watchlist: ./main remove id
   list, l    List watchlist's content: ./main list
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Add

```
NAME:
   main add - Add content to the watchlist: ./main add 'Name' 'tipo' 

USAGE:
   main add [arguments...]
```

Example:

```
go run main.go add 'Rapido y Furioso' 'pelicula' 

#1 - Monster Inc. - type: pelicula
#2 - Friends - type: serie
#3 - Todo sobre el asado - type: documental
#4 - Rapido y Furioso - type: pelicula
---Added---
```

## Change

```
NAME:
   main change - Change watchlist's content: ./main change id 'Name' 'tipo'

USAGE:
   main change [arguments...]
```


Example:

```
go run main.go change 2 'Mr. Robot' 'serie' 

#1 - Monster Inc. - type: pelicula
#3 - Todo sobre el asado - type: documental
#2 - Mr. Robot - type: serie
---Changed---
```

## Remove

```
NAME:
   main remove - Remove content by ID from the watchlist: ./main remove id

USAGE:
   main remove [arguments...]
```

Example:

```
go run main.go remove 2

#1 - Monster Inc. - type: pelicula
#3 - Todo sobre el asado - type: documental
Remove ID 2
---Changed---
```

## List

```
NAME:
   main list - List watchlist's content: ./main list

USAGE:
   main list [arguments...]
```

Example:

```
go run main.go list

#1 - Monster Inc. - type: pelicula
#2 - Friends - type: serie
#3 - Todo sobre el asado - type: documental
```


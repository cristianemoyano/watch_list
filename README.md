# Watchlist CLI

A watchlist using linked lists in Go.

##### Table of Contents  

- [Get started](#get-started)
- [Commands](#commands)
  * [Add](#add)
  * [Select](#select)
  * [Change](#change)
  * [Remove](#remove)
  * [List](#list)
  * [Movies](#movies)
  * [Help](#help)
  * [Exit](#exit)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>


# Get started

```
git clone https://github.com/cristianemoyano/watch_list.git

cd watch_list

go build main.go

./main

or

go run main.go
```

# Commands

```
go run main.go

Welcome to:
 __    __         _          _      _  _       _
/ / /\ \ \  __ _ | |_   ___ | |__  | |(_) ___ | |_
\ \/  \/ / / _` || __| / __|| '_ \ | || |/ __|| __|
 \  /\  / | (_| || |_ | (__ | | | || || |\__ \| |_
  \/  \/   \__,_| \__| \___||_| |_||_||_||___/ \__|

COMMANDS:

   login       Login into the platform.                 | login username
   add         Add content to the watchlist by Name.    | add Name tipo
   search      Search content                           | search Name
   select      Add content to the watchlist by ID.      | select id
   change      Change watchlist's content.              | change id Name tipo
   remove      Remove content by ID from the watchlist. | remove id
   list        List watchlist's content.                | list
   movies      List movies's content.                   | movies

   help        List available commands.                 | help
   exit        Finish program.                          | exit

>
```


## Login

```
NAME:
   login       Login into the platform.                 | login username

USAGE:
   > login Username
```

Example:

```
> login Joe

#325515328916291879 - Joe

---Welcome Joe !---
```

## Add

```
NAME:
   add         Add content to the watchlist by Name.    | add Name tipo

USAGE:
   > add name tipo
```

Example:

```
> add Live pelicula

#325436075143069991 - Live - type: pelicula

---Added by Name---
```

## Search

```
NAME:
   search      Search content                           | search Name

USAGE:
   > search name
```

Example:

```
> search Friends

#325515209210921255 - Friends - type: serie
```

## Select

```
NAME:
   select      Add content to the watchlist by ID.      | select id

USAGE:
   > select id
```

Example:

```
> movies

#325436061184426279 - Monster Inc. - type: pelicula
#325436061184491815 - Friends - type: serie
#325436061184557351 - Todo sobre el asado - type: documental

> select 325436061184491815

#325436075143069991 - Live - type: pelicula
#325436061184491815 - Friends - type: serie

---Added by ID---
```

## Change

```
NAME:
   change      Change watchlist's content.              | change id Name tipo

USAGE:
   change id Name tipo
```


Example:

```
> list

#325436075143069991 - Live - type: pelicula
#325436061184491815 - Friends - type: serie
#325436330458743079 - Fight - type: pelicula

> change 325436330458743079 Pedro serie

#325436075143069991 - Live - type: pelicula
#325436061184491815 - Friends - type: serie
#325436481906671911 - Pedro - type: serie
```

## Remove

```
NAME:
   remove      Remove content by ID from the watchlist. | remove id

USAGE:
   remove id
```

Example:

```
> remove 325436481906671911

ID to remove: 325436481906671911
#325436075143069991 - Live - type: pelicula
#325436061184491815 - Friends - type: serie
---Removed---
```

## List

```
NAME:
   list        List watchlist's content.                | list

USAGE:
   list
```

Example:

```
> list

#325436075143069991 - Live - type: pelicula
#325436061184491815 - Friends - type: serie
```

## Movies

```
NAME:
   movies      List movies's content.                   | movies

USAGE:
   movies
```

Example:

```
> movies

#325437058388590887 - Monster Inc. - type: pelicula
#325437058388656423 - Friends - type: serie
#325437058388721959 - Todo sobre el asado - type: documental
```

## Help

```
NAME:
   help        List available commands.                 | help

USAGE:
   help
```

Example:

```
> help

COMMANDS:

   add         Add content to the watchlist by Name.    | add Name tipo
   select      Add content to the watchlist by ID.      | select id
   change      Change watchlist's content.              | change id Name tipo
   remove      Remove content by ID from the watchlist. | remove id
   list        List watchlist's content.                | list
   movies      List movies's content.                   | movies

   help        List available commands.                 | help
   exit        Finish program.                          | exit
```

## Exit

```
NAME:
   exit        Finish program.                          | exit

USAGE:
   exit
```

Example:

```
> exit
```



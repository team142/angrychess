# Angry Chess (Chess for Four)
<img src="https://europe-west1-captains-badges.cloudfunctions.net/function-clone-badge-pc?project=team142/angrychess" />&nbsp;
<a href="https://goreportcard.com/report/github.com/team142/chessfor4"><img src="https://goreportcard.com/badge/github.com/team142/chessfor4" /></a>&nbsp;<a href="https://trello.com/b/czSy3gLz/battle-royale-chess"><img src="https://img.shields.io/badge/Project-Trello-brightgreen.svg" /></a>

## Background

Four player chess is a game I played growing up. I would love to relive the fun of playing "chess doubles" from 18 years ago and so this project was born.

- Two people play on a board as in normal chess.
- There are two teams consisting of two players.
- Each team has a one person playing as black and one person playing as white.
- When a player takes an enemy piece they give it to their teammate who can then place it on the board as a move.
- Pieces cannot be placed on the last two lines or such that they put the other player in check or mate.
- A timer can be introduced limiting the time a player has to make a move. If a move is not made within the time limit, then the other player may take a piece off the board from their opponent.
- I hope to add battle royale elements and abilities / points to increase the possibility space (and fun).

[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/3P7odVjrMiM/0.jpg)](https://www.youtube.com/watch?v=3P7odVjrMiM)


## Project structure

- Client is being written in JS
- Server has being written in Golang

## Running the project

### Server
- Use go version 1.11 or newer (Go mod support).
- Git clone the directory or your own fork of the project.
- `go run application.go`

### Client

- cd to client directory.
- `npm i`
- `npm run serve`

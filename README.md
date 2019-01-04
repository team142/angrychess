# Chess For 4 (Chess Doubles)
<a href="https://goreportcard.com/report/github.com/team142/chessfor4"><img src="https://goreportcard.com/badge/github.com/team142/chessfor4" /></a><img src="docs/trello.png" />

## Background

Four player chess is a game I played growing up. I would love to relive the fun of playing "chess doubles" from 18 years ago and so this project was born.

- Two people play on a board as in normal chess.
- There are two teams consisting of two players.
- Each team has a one person playing as black and one person playing as white.
- When a player takes an enemy piece they give it to their teammate who can then place it on the board as a move.
- Pieces cannot be placed on the last two lines or such that they put the other player in check or mate.
- A timer can be introduced limiting the time a player has to make a move. If a move is not made within the time limit, then the other player may take a piece off the board from their opponent.



## Project structure

- Client is being written in Vue JS
- Server has being written in Golang

## Running the project

### Server
- Use go version 1.11 or newer (Go mod support).
- Git clone the directory or your own fork of the project.
- `go run application.go`



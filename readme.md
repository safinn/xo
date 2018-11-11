# XO

DEMO: https://xo-go.now.sh

> A Go server using websockets to allow users to remotely play Noughts & Crosses through a Vue.js page and check the game status through a REST endpoint.

## How it works

At the root (`/`), a static html file using Vue.js is served connecting to the server via websockets that allows users to play the game. Clicking **New Game** will create and join a new game. Clicking **Join Game** will an join existing game with the specified ID.

An endpoint listening for requests at `/games/:id` where `id` is the game ID (int), returns the status of the specific game including the status and state. 

`status` property:  
**0** In progress  
**1** X won  
**2** O won  
**3** Draw  

The websocket handler listens for 3 actions from the client and returns a game state to the client including the results of the action. After a **MOVE** action the state of the game is sent to both players.

**NEW GAME** - Creates a Game using the struct and is persisted globaly inside a map.

**JOIN GAME** - Adds the client to an existing game with the provided ID if space is available.

**MOVE** - Checks if the client is a player, and the move is valid. The state of the game is then checked.

## Installation

```
dep ensure
```

## Run

```
go build && ./xo
```
http://localhost:8000

## Tests

```
go test
```

### Improvments

- File structure using packages
- Use external storage like Redis for example to persit data.
- Manage players stored in a game by removing them when a game is over.
- Allow clients other than the players to join a game for viewing only.
- Change status value on api from number to string

export let B = {}
export class NetworkManager {

    static state = {
        conn: {},
        game: {}

    }

    static connect(game, bIn) {
        B = bIn
        NetworkManager.state.game = game
        NetworkManager.state.conn = new WebSocket(NetworkManager.state.game.url);

        NetworkManager.state.conn.onopen = () => {
            NetworkManager.state.game.mutableViewServer = false;
            NetworkManager.state.game.mutableViewGames = true;
            NetworkManager.state.game.mutableViewGame = false;

            NetworkManager.sendNick(NetworkManager.state.game.nickname);
            NetworkManager.searchAgain();
        };

        NetworkManager.state.conn.onclose = () => {
            alert("Server is down!");
            location.reload();
        };
        NetworkManager.state.conn.onmessage = event => {
            try {
                const json = event.data;
                const o = JSON.parse(json);
                const msg = o.msg;
                if (msg === "secret") {
                    NetworkManager.state.game.secret = o.secret;
                    NetworkManager.state.game.id = o.id;
                } else if (msg === "list-games") {
                    NetworkManager.state.game.listOfGames = o.games.games;
                } else if (msg === "share-state") {
                    NetworkManager.state.game.gameState = o.game;
                } else if (msg === "view") {
                    if (o.view == "view-board") {
                        NetworkManager.state.game.mutableViewServer = false;
                        NetworkManager.state.game.mutableViewGames = false;
                        NetworkManager.state.game.mutableViewGame = true;

                        setTimeout(function () {
                            B.startup()
                        }, 500);


                    }
                } else {
                    alert(json);
                }
            } catch (e) {
                alert(e);
                alert(event.data);
            }
        };

    }

    static sendNick(nick) {
        NetworkManager.state.conn.send(
            JSON.stringify({
                msg: "nick",
                nick: nick
            })
        );
    }

    static searchAgain() {
        NetworkManager.state.conn.send(
            JSON.stringify({
                msg: "list-games"
            })
        );
    }

    static createGame() {
        NetworkManager.state.conn.send(
            JSON.stringify({
                msg: "create-game"
            })
        );
    }

    static joinGame(id) {
        NetworkManager.state.conn.send(
            JSON.stringify({
                msg: "join-game",
                id: id
            })
        );
    }

    static startGame() {
        NetworkManager.state.conn.send(
            JSON.stringify({
                msg: "start-game"
            })
        );
    }


}
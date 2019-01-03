export class NetworkManager {

    static state = {
        conn: {}
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
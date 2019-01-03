export class NetworkManager {

    static state = {
        conn: {},
        game: {}

    }

    static connect(game) {
        NetworkManager.state.game = game
        NetworkManager.state.conn = new WebSocket(NetworkManager.state.game.url);
  
        NetworkManager.state.conn.onopen = () => {
            NetworkManager.state.game.mutableViewServer = false;
            NetworkManager.state.game.mutableViewGames = true;
            NetworkManager.state.game.mutableViewGame = false;
    
            NetworkManager.sendNick(this.state.nickname);
            NetworkManager.searchAgain();
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
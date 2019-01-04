<template>
  <div class="hello">
    <div class="container-fluid">
      <div>
        <div v-if="state.mutableViewServer" class="row">
          <div class="col"></div>
          <div class="col-2">
            <h2>Chess for 4</h2>
            <img width="200px" alt="Logo" src="../assets/logo.png">

            <h3>Server</h3>
            <p>Get started by connecting to a server.</p>

            <input
              type="text"
              class="form-control-plaintext"
              v-model="state.nickname"
              placeholder="Nickname"
            >
            <br>
            <input type="text" class="form-control-plaintext" v-model="state.url" placeholder>
            <br>
            <button class="btn btn-success" v-on:click="connect">Connect</button>
          </div>
          <div class="col"></div>
        </div>
        <!-- /row -->
        <div v-if="state.mutableViewGames" class="row">
          <div class="col"></div>
          <div class="col-3">
            <h2>Chess for 4</h2>
            <img alt="Logo" src="../assets/logo.png">

            <h3>List of games</h3>
            <p>Pick a game or start your own.</p>

            <button class="btn btn-primary" v-on:click="createGame">New game</button>&nbsp;
            <button class="btn btn-primary" v-on:click="searchAgain">Refresh</button>

            <div class="md-layout">
              <div class="md-layout-item"></div>
              <div class="md-layout-item">
                <br>
                <table class="table table-hover">
                  <tr>
                    <th md-numeric>Name</th>
                    <th>Players</th>
                    <th>Actions</th>
                  </tr>

                  <tr v-for="game in state.listOfGames" :key="game['id']">
                    <td>{{ game["title"] }}</td>
                    <td>{{ game["players"] }}</td>
                    <td>
                      <button class="btn btn-primary btn-sm" v-on:click="joinGame(game['id'])">Join</button>
                    </td>
                  </tr>
                </table>
              </div>
              <div class="md-layout-item"></div>
            </div>
          </div>
          <div class="col"></div>
        </div>

        <div v-if="state.mutableViewGame">
          <canvas id="renderCanvas"></canvas>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
const networkR = require("./network.js");
const stateR = require("./state.js");
const B = require("./engine.js").B;

export default {
  name: "HelloWorld",
  props: {
    msg: String
  },
  data: function() {
    return {
      NetworkManager: networkR.NetworkManager,
      state: new stateR.State()
    };
  },
  methods: {
    connect() {
      if (!this.state.nickname) {
        alert("You need a nickname");
        return;
      }
      this.NetworkManager.connect(
        this.state,
        B
      );
    },

    searchAgain() {
      this.NetworkManager.searchAgain();
    },

    createGame() {
      this.NetworkManager.createGame();
      this.state.admin = true;
    },

    joinGame(id) {
      this.NetworkManager.joinGame(id);
      this.state.admin = false;
    },

    startGame() {
      this.NetworkManager.startGame();
    },

    getBoards() {
      let boards = [];
      let i = 1;
      for (i = 1; i <= this.state.gameState.boards; i++) {
        let board = {};
        board.id = i;
        boards.push(board);
      }

      return boards;
    },

    myTurn() {
      const id = this.state.id;
      // console.log(this.state.gameState.players)
      for (let seat in this.state.gameState.players) {
        if (this.state.gameState.players[seat].profile.id == id) {
          return this.state.gameState.players[seat].myTurn;
        }
      }

      return true;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
body {
  background: #000;
}
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
#renderCanvas {
  width: 100%;
  height: 100%;
  touch-action: none;
}
</style>



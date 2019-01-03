<template>
  <div class="hello">
    <div>
      <div v-if="state.mutableViewServer">
        <h2>Chess for 4</h2>
        <img alt="Logo" src="../assets/logo.png">

        <h3>Server</h3>
        <p>Get started by connecting to a server.</p>

        <input v-model="state.nickname" placeholder="Nickname">
        <br>
        <input v-model="state.url" placeholder>
        <br>
        <md-button class="md-raised md-primary" v-on:click="connect">Connect</md-button>
      </div>

      <div v-if="state.mutableViewGames">
        <h2>Chess for 4</h2>
        <img alt="Logo" src="../assets/logo.png">

        <h3>List of games</h3>
        <p>Pick a game or start your own.</p>

        <div class="md-layout">
          <div class="md-layout-item"></div>
          <div class="md-layout-item">
            <md-table>
              <md-table-row>
                <md-table-head md-numeric>Name</md-table-head>
                <md-table-head>Players</md-table-head>
                <md-table-head>Actions</md-table-head>
              </md-table-row>

              <md-table-row>
                <md-table-cell md-numeric>New game?</md-table-cell>
                <md-table-cell></md-table-cell>
                <md-table-cell>
                  <md-button class="md-raised md-primary" v-on:click="createGame">Create</md-button>
                </md-table-cell>
              </md-table-row>

              <md-table-row>
                <md-table-cell md-numeric>Search again?</md-table-cell>
                <md-table-cell></md-table-cell>
                <md-table-cell>
                  <md-button class="md-raised md-primary" v-on:click="searchAgain">Refresh</md-button>
                </md-table-cell>
              </md-table-row>

              <md-table-row v-for="game in state.listOfGames" :key="game['id']">
                <md-table-cell>{{ game["title"] }}</md-table-cell>
                <md-table-cell>{{ game["players"] }}</md-table-cell>
                <md-table-cell>
                  <md-button class="md-raised md-primary" v-on:click="joinGame(game['id'])">Join</md-button>
                </md-table-cell>
              </md-table-row>
            </md-table>
          </div>
          <div class="md-layout-item"></div>
        </div>
      </div>

      <div v-if="state.mutableViewGame">
        <h2>Chess for 4</h2>
        <md-button
          class="md-raised md-primary"
          v-on:click="startGame"
          v-if="state.admin && !state.gameState.started"
        >Start game</md-button>
        <br>
        {{ JSON.stringify(state.gameState)}}
      </div>
    </div>
  </div>
</template>

<script>
const network = require("./network.js");
const stateR = require("./state.js");
// const NetworkManager = network.NetworkManager;

export default {
  name: "HelloWorld",
  props: {
    msg: String
  },
  data: function() {
    return {
      NetworkManager: network.NetworkManager,
      state: new stateR.State(),

      // mutableViewServer: true,
      // mutableViewGames: false,
      // mutableViewGame: false,
      // nickname: "Swag",
      // url: "ws://localhost:8000/ws",
      // conn: {},
      // secret: "",
      // gameState: {},
      // listOfGames: [],
      // admin: false,
      // id: {}
    };
  },
  methods: {
    connect() {
      if (!this.state.nickname) {
        alert("You need a nickname");
        return;
      }
      this.state.conn = new WebSocket(this.state.url);
      this.NetworkManager.state.conn = this.state.conn;

      this.state.conn.onopen = () => {
        this.state.mutableViewServer = false;
        this.state.mutableViewGames = true;
        this.state.mutableViewGame = false;
        this.state.conn.send(
          JSON.stringify({
            msg: "nick",
            nick: this.state.nickname
          })
        );
        this.searchAgain();
      };

      this.state.conn.onclose = () => {
        alert("closed ws");
      };
      this.state.conn.onmessage = event => {
        try {
          const json = event.data;
          const o = JSON.parse(json);
          const msg = o.msg;
          if (msg === "secret") {
            this.state.secret = o.secret;
            this.state.id = o.id;
          } else if (msg === "list-games") {
            this.state.listOfGames = o.games.games;
          } else if (msg === "share-state") {
            this.state.gameState = o.game;
          } else if (msg === "view") {
            if (o.view == "view-board") {
              this.state.mutableViewServer = false;
              this.state.mutableViewGames = false;
              this.state.mutableViewGame = true;
            }
          } else {
            alert(json);
          }
        } catch (e) {
          alert(e);
          alert(event.data);
        }
      };
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
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
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
</style>



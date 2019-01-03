<template>
  <div class="hello">
    <div>
      <div v-if="mutableViewServer">
        <h2>Chess for 4</h2>
        <img alt="Logo" src="../assets/logo.png">

        <h3>Server</h3>
        <p>Get started by connecting to a server.</p>

        <input v-model="nickname" placeholder="Nickname">
        <br>
        <input v-model="url" placeholder>
        <br>
        <md-button class="md-raised md-primary" v-on:click="connect">Connect</md-button>
      </div>

      <div v-if="mutableViewGames">
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

              <md-table-row v-for="game in listOfGames" :key="game['id']">
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

      <div v-if="mutableViewGame">
        <h2>Chess for 4</h2>
        <md-button
          class="md-raised md-primary"
          v-on:click="startGame"
          v-if="admin && !gameState.started"
        >Start game</md-button>
        <br>
        {{ JSON.stringify(gameState)}}
      </div>
    </div>
  </div>
</template>

<script>
const network = require("./network.js");
const NetworkManager = network.NetworkManager;

export default {
  name: "HelloWorld",
  props: {
    msg: String
  },
  data: function() {
    return {
      mutableViewServer: true,
      mutableViewGames: false,
      mutableViewGame: false,
      nickname: "Swag",
      url: "ws://localhost:8000/ws",
      conn: {},
      secret: "",
      gameState: {},
      listOfGames: [],
      admin: false,
      id: {}
    };
  },
  methods: {
    connect() {
      if (!this.nickname) {
        alert("You need a nickname");
        return;
      }
      this.conn = new WebSocket(this.url);
      NetworkManager.state.conn = this.conn;
      this.conn.onopen = () => {
        this.mutableViewServer = false;
        this.mutableViewGames = true;
        this.mutableViewGame = false;
        this.conn.send(
          JSON.stringify({
            msg: "nick",
            nick: this.nickname
          })
        );
        this.searchAgain();
      };
      this.conn.onclose = () => {
        alert("closed ws");
      };
      this.conn.onmessage = event => {
        try {
          const json = event.data;
          const o = JSON.parse(json);
          const msg = o.msg;
          if (msg === "secret") {
            this.secret = o.secret;
            this.id = o.id;
          } else if (msg === "list-games") {
            this.listOfGames = o.games.games;
          } else if (msg === "share-state") {
            this.gameState = o.game;
          } else if (msg === "view") {
            if (o.view == "view-board") {
              this.mutableViewServer = false;
              this.mutableViewGames = false;
              this.mutableViewGame = true;
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
      NetworkManager.searchAgain();
    },

    createGame() {
      NetworkManager.createGame();
      this.admin = true;
    },

    joinGame(id) {
      NetworkManager.joinGame(id);
      this.admin = false;
    },

    startGame() {
      NetworkManager.startGame();
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



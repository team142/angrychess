<template>
  <div class="hello">
    <p>Get started by connecting to a server.</p>

    <div>
      <div v-if="mutableViewServer">
        <h3>Server</h3>
        <input type="text" id="nick" name="nick" value="Nickname" /><br />
        <input type="text" id="url" name="url" value="ws://localhost:8000/ws" /><br />
        <md-button class="md-raised md-primary" v-on:click="connect">Connect</md-button>
      </div>
      <div v-if="mutableViewGames">
        <h4>List of games</h4>

        <div class="md-layout">
          <div class="md-layout-item"></div>
          <div class="md-layout-item">
            <md-table>
              <md-table-row>
                <md-table-head md-numeric>Name</md-table-head>
                <md-table-head>Players</md-table-head>
                <md-table-head>Join</md-table-head>
              </md-table-row>

              <md-table-row>
                <md-table-cell md-numeric>Captains game</md-table-cell>
                <md-table-cell>1/4</md-table-cell>
                <md-table-cell>
                  <md-button class="md-raised md-primary" v-on:click="connect">Join</md-button>
                </md-table-cell>
              </md-table-row>
            </md-table>
          </div>
          <div class="md-layout-item"></div>
        </div>
      </div>
      <div v-if="mutableViewGame"></div>
    </div>
  </div>
</template>

<script>
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
      conn: {}
    };
  },
  methods: {
    connect() {
      const url = "ws://localhost:8000/ws";
      this.conn = new WebSocket(url);
      this.conn.onclose = function(evt) {
        console.log("Closed ws");
      };
      this.conn.onmessage = function(evt) {
        var messages = evt.data.split("\n");
        for (var i = 0; i < messages.length; i++) {
          console.log("Recieved: " + messages[i]);
        }
      };

      //TODO: connect to the server
      this.mutableViewGames = true;
      this.mutableViewServer = false;
      this.mutableViewGame = false;
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



<template>
  <div id="app">
    <info :msg="v" area="version"/>
    <player-control
      :status="status"
      :command="command"
      :position="position"
      :queue-length="queueLength"
      area="playerControl"
    />
    <view-control
      area="view"
      areaControl="viewControl"
      :queue="queue"
      :search="search"
      :browse="browse"
      :playlist="playlist"
      :command="command"
    />
    <active-song area="activeSong" :song="activeSong" class="bordered"/>
  </div>
</template>

<style>
/* FIXME: move the next 2 somewhere global? */
:root {
  /* copied from https://www.w3schools.com/lib/w3-theme-red.css */
  --background: #efefef;
  --foreground: #000000;
  --c1: #8d8c8c;
  --c2: #686868;
  --c4: #8d8c8c;
  --c5: #aaa8a8;
}

body {
  /* FIXME: better font selection */
  font-family: "Righteous", cursive;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

  font-size: 24px;
  margin: 4px;
  padding: 0px;
  color: var(--foreground);
  background: var(--background);
}

button {
  font-family: "Righteous", cursive;
  font-size: 24px;
}

input[type="text"],
input[type="number"],
input[type="password"] {
  width: 100%;
  border: 1px solid var(--foreground);
  border-radius: 4px;
  box-sizing: border-box;
  -webkit-box-sizing: border-box;
  -moz-box-sizing: border-box;
  font-size: 1em;
}

#app {
  display: grid;
  grid-template-columns: 4px 1fr 4px;
  grid-gap: 4px;
  grid-template-areas:
    ". viewControl ."
    ". activeSong ."
    ". playerControl ."
    ". status ."
    ". view ."
    ". version .";
}

@media (min-width: 768px) {
  #app {
    display: grid;
    grid-template-columns: 4px auto 1fr auto 4px;
    grid-template-areas:
      ".  viewControl activeSong status ."
      ". . playerControl . ."
      ". view view view ."
      ". version version version .";
  }
}

.bordered {
  /* add a round border & set the background */
  border: 2px solid var(--foreground);
  border-radius: 4px;
  margin: 1px;
  /*padding-bottom: 4px;*/

  background: var(--c5);
}
</style>

<script>
import Info from "./components/Info.vue";
import PlayerControl from "./components/PlayerControl.vue";
import ViewControl from "./components/views/Control.vue";
import ActiveSong from "./components/ActiveSong.vue";

export default {
  name: "app",
  data: function() {
    return {
      websocket: null,
      version: "unknown",
      status: "",
      activeSong: null,
      queue: null,
      search: null,
      browse: null,
      playlist: null
    };
  },
  components: {
    Info,
    PlayerControl,
    ViewControl,
    ActiveSong
  },
  computed: {
    v: function() {
      return "MPD Protocol Version: " + this.version;
    },
    position: function() {
      if (this.activeSong == null) return -1;
      return this.activeSong.position;
    },
    queueLength: function() {
      // FIXME: implement
      return 0;
    }
  },
  methods: {
    openWebSocket: function(app) {
      var addr = document.getElementById("ws").value;
      if (addr == "{{.ws}}") {
        // not served from go, try this
        addr = "ws://" + window.location.hostname + ":8666/ws";
      }
      app.websocket = new WebSocket(addr); // FIXME: where to get the address
      app.websocket.onopen = function() {
        // eslint-disable-next-line
        console.log("OPEN");
      };
      app.websocket.onclose = function() {
        // eslint-disable-next-line
        console.log("CLOSE");
        app.websocket = null;
        // eslint-disable-next-line
        console.log("no connection&hellip; reconnecting in 4 seconds&hellip;");
        setTimeout(app.openWebSocket, 4000);
      };
      app.websocket.onmessage = function(evt) {
        // eslint-disable-next-line
        console.log(evt.data);
        const e = JSON.parse(evt.data);
        switch (e.type) {
          case "version":
            app.version = e.data;
            break;
          case "update":
            app.status = e.data.status.state;
            app.queue = e.data.queue;
            app.activeSong = e.data.activeSong;
            // eslint-disable-next-line
            console.log(app.status);
            break;
          case "search":
            app.search = e.data;
            break;
        }
      };
      app.websocket.onerror = function(evt) {
        // eslint-disable-next-line
        console.log({ evt });
      };
    },
    command: function(cmd, data, page = 1) {
      if (!this.websocket) return false;
      var d = data.toString();
      // eslint-disable-next-line
      console.log({ cmd, d, page });
      this.websocket.send(
        JSON.stringify({ command: cmd, data: d, page: page })
      );
      return true;
    }
  },
  mounted() {
    this.openWebSocket(this);
  }
};
</script>


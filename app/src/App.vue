<template>
  <div id="app">
    <info :msg="status" area="status"/>
    <info :msg="v" area="version"/>
    <player-control
      :status="status"
      :command="command"
      :position="position"
      :queue-length="queueLength"
      area="playerControl"
    />
    <view-control area="viewControl" @view-changed="setView($event)"/>
    <active-song area="activeSong" :song="activeSong" class="bordered"/>
    <queue v-if="view=='queue'" area="view"/>
    <search v-if="view=='search'" area="view"/>
    <browse v-if="view=='browse'" area="view"/>
    <playlist v-if="view=='playlist'" area="view"/>
  </div>
</template>

<script>
import Info from "./components/Info.vue";
import PlayerControl from "./components/PlayerControl.vue";
import ViewControl from "./components/ViewControl.vue";
import ActiveSong from "./components/ActiveSong.vue";
import Queue from "./components/Queue.vue";
import Browse from "./components/Browse.vue";
import Search from "./components/Search.vue";
import Playlist from "./components/Playlist.vue";

var ws = null;
var openWebSocket = function(app) {
  ws = new WebSocket("ws://192.168.0.111:8666/ws"); // FIXME: where to get the address
  ws.onopen = function() {
    // eslint-disable-next-line
    console.log("OPEN");
  };
  ws.onclose = function() {
    // eslint-disable-next-line
    console.log("CLOSE");
    ws = null;
    // eslint-disable-next-line
    console.log("no connection&hellip; reconnecting in 4 seconds&hellip;");
    setTimeout(openWebSocket, 4000);
  };
  ws.onmessage = function(evt) {
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
        break;
    }
  };
  ws.onerror = function(evt) {
    // eslint-disable-next-line
    console.log({ evt });
  };
};
var c = function(cmd, data) {
  if (!ws) return false;
  // eslint-disable-next-line
  var d = data.toString();
  console.log({ cmd, d });
  if (cmd == "search" && d.length < 3) {
    return false;
  }
  ws.send(JSON.stringify({ command: cmd, data: d }));
  return true;
};

export default {
  name: "app",
  data: function() {
    return {
      status: "",
      activeSong: null,
      queue: null,
      version: "unknown",
      command: c,
      view: "queue"
    };
  },
  components: {
    Info,
    PlayerControl,
    ViewControl,
    ActiveSong,
    Queue,
    Browse,
    Search,
    Playlist
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
    setView: function(v) {
      this.view = v;
    }
  },
  mounted() {
    openWebSocket(this);
  }
};
</script>

<style>
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

  /*  font-family: "Roboto Mono", monospace;*/
  font-family: "Righteous", cursive;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;

  font-size: 24px;
  margin: 4px;
  padding: 0px;
  color: var(--foreground);
  background: var(--background);
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

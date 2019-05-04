<template>
  <div class="activeSong" :style="style">
    <div
      v-for="(v, k) in attrs"
      class="item"
      :class="v.name"
      :key="k"
      :title="v.title + ': ' + attr(v.name)"
      :style="{gridArea:v.name}"
    >{{attr(v.name)}}</div>
    <div class="progress">
      <div :style="{width:progress+'%'}"></div>
    </div>
  </div>
</template>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.item {
  text-align: center;
  align-self: center;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.activeSong {
  padding: 4px;
  display: grid;
  grid-template-columns: 1fr;
  grid-template-areas: "title" "artist" "album" "progress";
}
.progress {
  grid-area: progress;
  border: 1px solid var(--foreground);
  align-self: center;
}
.progress div {
  background: var(--background);
  height: 0.4em;
  border-right: 1px solid var(--foreground);
}
</style>

<script>
export default {
  name: "ActiveSong",
  data: function() {
    return {
      attrs: [
        { name: "artist", title: "Artist" },
        { name: "title", title: "Title" },
        { name: "album", title: "Album" }
      ]
    };
  },
  props: {
    area: String,
    song: Object
  },
  components: {},
  computed: {
    style: function() {
      var result = {};
      if (this.area) result["gridArea"] = this.area;
      return result;
    },
    progress: function() {
      if (this.song && this.song.elapsed != -1) {
        return (100.0 * this.song.elapsed) / this.song.duration;
      }
      return 0;
    }
  },
  methods: {
    attr: function(a) {
      if (!this.song) {
        return "";
      }
      if (a == "album") {
        if (
          this.song["album_artist"] != "" &&
          this.song["album_artist"] != this.song["artist"]
        ) {
          return this.song["album"] + " [" + this.song["album_artist"] + "]";
        }
      } else if (a == "duration") {
        var value = this.song["duration"];
        var min = parseInt(value / 60);
        var sec = parseInt(value % 60);
        if (sec < 10) {
          sec = "0" + sec;
        }
        return min + ":" + sec;
      }
      return this.song[a];
    },
    tick: function() {
      if (this.song && this.song.playing) {
        this.song.elapsed += 1;
      }
      setTimeout(this.tick, 1000);
    }
  },
  mounted() {
    setTimeout(this.tick, 1000);
  }
};
</script>

        
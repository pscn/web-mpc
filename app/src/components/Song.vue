<template>
  <Fragment>
    <div
      v-for="(v, k) in attrs"
      class="item"
      :class="v.name"
      :key="k"
      :title="v.title + ': ' + attr(v.name)"
      :style="{gridArea:v.name}"
    >{{attr(v.name)}}</div>
  </Fragment>
</template>

<script>
import { Fragment } from "vue-fragment";

export default {
  name: "Song",
  data: function() {
    return {
      attrs: [
        { name: "artist", title: "Artist" },
        { name: "title", title: "Title" },
        { name: "album", title: "Album" },
        { name: "duration", title: "Duration" }
      ]
    };
  },
  props: {
    song: Object
  },
  components: {
    Fragment
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
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.duration {
  text-align: right;
}
.item {
  min-height: 0;
  /*max-height: 1.2em;*/
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
        
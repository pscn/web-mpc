<template>
  <Fragment>
    <div
      v-for="(v, k) in attrs"
      class="item"
      :key="k"
      :title="attr(v.name)"
      :class="v.name"
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
        { name: "album", title: "Album" }
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
      }
      return this.song[a];
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.artist {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.title {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.album {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
.item {
  min-height: 0;
  max-height: 1.2em;
}
.viewControl {
  display: grid;
  grid-template-columns: auto auto auto auto 1fr;
  grid-template-areas: "queue search browse playlist .";

  grid-gap: 2px;
  min-height: 0;
  min-width: 0;
}

@media (min-width: 768px) {
  .item {
    max-height: 1em;
  }
  .viewControl {
    grid-template-columns: 1fr;
    grid-template-areas:
      "queue"
      "search"
      "browse"
      "playlist";
  }
}
</style>
        
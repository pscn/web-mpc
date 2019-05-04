<template>
  <div class="activeSong" :style="style">
    <song-parts :song="song"/>
  </div>
</template>

<script>
import SongParts from "./SongParts.vue";

export default {
  name: "ActiveSong",
  props: {
    area: String,
    song: Object
  },
  components: { SongParts },
  computed: {
    style: function() {
      var result = {};
      if (this.area) result["gridArea"] = this.area;
      return result;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
.activeSong .item {
  text-align: center;
  align-self: center;
}
.activeSong {
  display: grid;
  grid-template-columns: 1fr 1fr;
  grid-template-areas:
    "title title"
    "artist artist"
    "album album"
    "duration duration";
}
</style>
        
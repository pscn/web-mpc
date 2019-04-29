<template>
  <div class="activeSong" :style="style">
    <song :song="song"/>
  </div>
</template>

<script>
import Song from "./Song.vue";

export default {
  name: "ActiveSong",
  props: {
    area: String,
    song: Object
  },
  components: { Song },
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
<style scoped>
.item {
  text-align: center;
  align-self: center;
}
.activeSong {
  display: grid;
  grid-template-columns: 1fr;
  grid-template-areas:
    "title"
    "artist"
    "album";
}
</style>
        
<template>
  <div class="song" :style="style">
    <song-parts :song="song"/>
  </div>
</template>

<script>
import SongParts from "./SongParts.vue";

export default {
  name: "Song",
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
  },
  methods: {}
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.duration {
  text-align: right;
}

.song {
  display: grid;
  grid-template-columns: 1fr auto;
  grid-template-areas:
    "title title"
    "artist artist"
    "album duration";

  padding: 2px 2px 2px 2px;
  grid-gap: 2px;
  min-height: 0;
  min-width: 0;
}

@media (min-width: 768px) {
  .song {
    grid-template-columns: 15fr 12fr 18fr 3fr;
    grid-template-areas: "title artist album duration";
  }
}
</style>
        
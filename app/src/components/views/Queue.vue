<template>
  <div class="view" :style="style">
    <div v-if="queue!=null" class="songs">
      <div v-for="(song, k) in queue.songs" :key="k" class="song bordered">
        <song :song="song"/>
      </div>
    </div>
  </div>
</template>

<script>
import Song from "./../Song.vue";

export default {
  name: "Queue",
  props: {
    area: String,
    queue: Object
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
.song {
  display: grid;
  grid-template-columns: 1fr;
  grid-template-areas: "title" "artist" "album";

  grid-gap: 2px;
  min-height: 0;
  min-width: 0;
}
@media (min-width: 768px) {
  .song {
    grid-template-columns: 1fr 1fr 1fr;
    grid-template-areas: "title artist album";
  }
}
</style>
        
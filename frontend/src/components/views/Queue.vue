<template>
  <div class="view" :style="style">
    <pagination
      v-if="queue!=null"
      :page="queue.page"
      :last-page="queue.lastPage"
      @goto="goto($event)"
    />
    <div v-if="queue!=null" class="songs">
      <div v-for="(song, k) in queue.songs" :key="k" class="song bordered">
        <song :song="song"/>
      </div>
    </div>
  </div>
</template>

<script>
import Song from "./../Song.vue";
import Pagination from "./Pagination.vue";

export default {
  name: "Queue",
  props: {
    area: String,
    queue: Object,
    command: Function
  },
  components: { Song, Pagination },
  computed: {
    style: function() {
      var result = {};
      if (this.area) result["gridArea"] = this.area;
      return result;
    }
  },
  methods: {
    goto: function(page) {
      // eslint-disable-next-line
      console.log("goto: " + page);
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
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
        
<template>
  <div class="view" :style="style">
    <pagination
      v-if="queue!=null"
      area="pagination"
      :page="queue.page"
      :last-page="queue.lastPage"
      @goto="goto($event)"
    />

    <div v-if="queue!=null" class="songs">
      <div v-for="(song, k) in queue.songs" :key="k" class="bordered row">
        <div class="control">Ctrl</div>
        <song :song="song" area="song"/>
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
.view {
  display: grid;
  grid-template-columns: 1fr;
  grid-template-areas:
    "pagination"
    "songs";
  grid-gap: 2px;
  min-height: 0;
  min-width: 0;
}
.control {
  grid-area: control;
}
.row {
  display: grid;
  grid-template-columns: auto 1fr;
  grid-template-areas: "control song";

  grid-gap: 2px;
  min-height: 0;
  min-width: 0;
}
</style>
        
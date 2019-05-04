<template>
  <div class="view" :style="style">
    <input
      type="text"
      v-model="searchText"
      v-on:input="cmd()"
      placeholder="any t:title a:artist al:album"
      style="grid-area:searchTxt"
    >
    <button-icon icon="search" area="searchBtn"/>

    <pagination
      v-if="search!=null"
      :page="search.page"
      :last-page="search.lastPage"
      area="pagination"
      @goto="goto($event)"
    />

    <div v-if="search!=null" class="songs">
      <song v-for="(song, k) in search.songs" :key="k" :song="song" area="song" class="bordered"/>
    </div>
  </div>
</template>

<script>
import ButtonIcon from "../ButtonIcon.vue";
import Pagination from "./Pagination.vue";
import Song from "../Song.vue";

export default {
  name: "Search",
  data: function() {
    return {
      searchText: ""
    };
  },
  props: {
    area: String,
    command: Function,
    search: Object
  },
  components: { ButtonIcon, Pagination, Song },
  computed: {
    style: function() {
      var result = {};
      if (this.area) result["gridArea"] = this.area;
      return result;
    }
  },
  methods: {
    cmd: function() {
      if (this.searchText.length >= 3) {
        // eslint-disable-next-line
        console.log(this.searchText);
        this.command("search", this.searchText, 1);
      }
    },
    goto: function(page) {
      // eslint-disable-next-line
      console.log("goto: " + page);
      this.command("search", this.searchText, page);
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.item {
  min-height: 0;
}
.songs {
  grid-area: songs;
}
.view {
  display: grid;
  grid-template-columns: 1fr auto;
  grid-template-areas:
    "searchTxt searchBtn"
    "pagination pagination"
    "songs songs";

  grid-gap: 2px;
  min-height: 0;
  min-width: 0;
}
</style>
        
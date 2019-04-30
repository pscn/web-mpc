<template>
  <Fragment>
    <div class="viewControl" :style="controlStyle">
      <button-icon
        v-for="(v, k) in views"
        class="item"
        :key="k"
        :disabled="view==v.name"
        :icon="v.name"
        :title="v.title"
        :area="v.name"
        @click="view=v.name"
      />
    </div>

    <queue v-if="view=='queue'" :queue="queue" :style="viewStyle"/>
    <search v-if="view=='search'" :search="search" :style="viewStyle"/>
    <browse v-if="view=='browse'" :browse="browse" :style="viewStyle"/>
    <playlist v-if="view=='playlist'" :playlist="playlist" :style="viewStyle"/>
  </Fragment>
</template>

<script>
import { Fragment } from "vue-fragment";

import ButtonIcon from "../ButtonIcon.vue";
import Queue from "./Queue.vue";
import Search from "./Search.vue";
import Browse from "./Browse.vue";
import Playlist from "./Playlist.vue";

export default {
  name: "ViewControl",
  data: function() {
    return {
      view: "queue",
      views: [
        { name: "queue", title: "Queue" },
        { name: "search", title: "Search" },
        { name: "browse", title: "Browse" },
        { name: "playlist", title: "Playlist" }
      ]
    };
  },
  props: {
    area: String,
    areaControl: String,
    queue: Object,
    browse: Object,
    search: Object,
    playlist: Object
  },
  components: {
    Fragment,
    ButtonIcon,
    Queue,
    Search,
    Browse,
    Playlist
  },
  computed: {
    controlStyle: function() {
      var result = {};
      if (this.area) result["gridArea"] = this.areaControl;
      return result;
    },
    viewStyle: function() {
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
        
<template>
  <Fragment>
    <div class="viewControl" :style="style">
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

    <queue v-if="view=='queue'" :queue="queue" :command="command" :area="area"/>
    <search v-if="view=='search'" :search="search" :command="command" :area="area"/>
    <folder v-if="view=='folder'" :folder="folder" :command="command" :area="area"/>
    <playlist v-if="view=='playlist'" :playlist="playlist" :command="command" :area="area"/>
  </Fragment>
</template>

<script>
import { Fragment } from "vue-fragment";

import ButtonIcon from "../ButtonIcon.vue";
import Queue from "./Queue.vue";
import Search from "./Search.vue";
import Folder from "./Folder.vue";
import Playlist from "./Playlist.vue";

export default {
  name: "ViewControl",
  data: function() {
    return {
      view: "queue",
      views: [
        { name: "queue", title: "Queue" },
        { name: "search", title: "Search" },
        { name: "folder", title: "Folder" },
        { name: "playlist", title: "Playlist" }
      ]
    };
  },
  props: {
    area: String,
    areaControl: String,
    queue: Object,
    folder: Object,
    search: Object,
    playlist: Object,
    command: Function
  },
  components: {
    Fragment,
    ButtonIcon,
    Queue,
    Search,
    Folder,
    Playlist
  },
  computed: {
    style: function() {
      var result = {};
      if (this.area) result["gridArea"] = this.areaControl;
      return result;
    },
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
  grid-template-areas: "queue search folder playlist .";

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
      "folder"
      "playlist";
  }
}
</style>
        
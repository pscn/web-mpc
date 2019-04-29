<template>
  <div class="viewControl" :style="style">
    <button-icon
      v-for="(v, k) in views"
      class="item"
      :key="k"
      :disabled="view==v.name"
      :icon="v.name"
      :title="v.title"
      :area="v.name"
      @click="view=v.name; $emit('view-changed', v.name)"
    />
  </div>
</template>

<script>
import ButtonIcon from "./ButtonIcon.vue";

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
    area: String
  },
  components: {
    ButtonIcon
  },
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
        
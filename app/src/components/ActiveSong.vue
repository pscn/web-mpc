<template>
  <div class="activeSong" :style="style">
    <div
      v-for="(v, k) in attrs"
      class="item"
      :key="k"
      :title="v.title"
      :area="v.name"
    >{{attr(v.name)}}</div>
  </div>
</template>

<script>
export default {
  name: "ActiveSong",
  data: function() {
    return {
      attrs: [
        { name: "artist", title: "Artist" },
        { name: "title", title: "Title" },
        { name: "album", title: "Album" }
      ]
    };
  },
  props: {
    area: String,
    song: Object
  },
  components: {},
  computed: {
    style: function() {
      var result = {};
      if (this.area) {
        result["gridArea"] = this.area;
      }
      return result;
    }
  },
  methods: {
    attr: function(a) {
      if (this.song) return this.song[a];
      return "";
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
    "artist"
    "title"
    "album";
}

</style>
        
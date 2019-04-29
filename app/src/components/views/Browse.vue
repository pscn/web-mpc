<template>
  <div class="view" :style="style">Browse</div>
</template>

<script>
export default {
  name: "Browse",
  props: {
    area: String
  },
  components: {},
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
  max-height: 0.8em;
}
.view {
  display: grid;
  grid-template-columns: 1fr;
  grid-template-areas: "item";

  grid-gap: 2px;
  min-height: 0;
  min-width: 0;
}
</style>
        
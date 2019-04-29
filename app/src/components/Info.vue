<template>
  <div class="info" :style="style">{{ msg }}</div>
</template>

<script>
export default {
  name: "Info",
  props: {
    msg: String,
    area: String
  },
  computed: {
    style: function() {
      if (this.area) {
        return { gridArea: this.area };
      }
      return {};
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.info {
  font-size: 0.5em;
}
</style>

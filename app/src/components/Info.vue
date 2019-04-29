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
      var result = {};
      if (this.area) result["gridArea"] = this.area;
      return result;
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

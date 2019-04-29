<template>
  <svg class="icon" :style="{stroke:color,fill:fill}" aria-hidden="true" focusable="false">
    <use :xlink:href="icon"></use>
  </svg>
</template>

<script>
export default {
  name: "Icon",
  props: {
    name: {
      type: String,
      default: ""
    },
    color: {
      type: String,
      default: "black"
    },
    fill: {
      type: String,
      default: "black"
    }
  },
  computed: {
    icon: function() {
      return "#icon-" + this.name;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.icon {
  width: 0.9em;
  height: 0.9em;
  overflow: hidden;
  display: block;
  stroke-width: 8%;
  stroke-linejoin: round;
}
</style>
        
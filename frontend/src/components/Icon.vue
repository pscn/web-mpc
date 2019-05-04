<template>
  <svg class="icon" :style="style" aria-hidden="true" focusable="false">
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
    color: String,
    fill: String
  },
  computed: {
    icon: function() {
      return "#icon-" + this.name;
    },
    style: function() {
      var result = {};
      if (this.color) result["stroke"] = this.color;
      if (this.fill) result["fill"] = this.fill;
      return result;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.icon {
  width: 1.1em;
  height: 1.1em;
  overflow: hidden;
  display: block;
  fill: none;
  stroke: var(--foreground);
  stroke-width: 8%;
  stroke-linejoin: round;
}

@media (min-width: 768px) {
  .icon {
    width: 0.9em;
    height: 0.9em;
  }
}
</style>
        
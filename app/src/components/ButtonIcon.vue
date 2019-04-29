<template>
  <button class="button" :style="style" @click="$emit('click')">
    <Icon :name="icon" :color="color" :fill="iconFill"/>
  </button>
</template>

<script>
import Icon from "./Icon.vue";

export default {
  name: "ButtonIcon",
  props: {
    icon: String,
    area: String,
    borderColor: String,
    color: String,
    iconFill: {
      type: String,
      default: "none"
    },
    background: String
  },
  computed: {
    border: function() {
      return "2px solid" + this.borderColor;
    },
    style: function() {
      var result = {};
      if (this.border) {
        result["border"] = this.border;
      }
      if (this.background) {
        result["background"] = this.background;
      }
      if (this.area) {
        result["gridArea"] = this.area;
      }
      return result;
    }
  },
  components: {
    Icon
  },
  methods: {
    doClick: function() {
      if (this.click != null) {
        this.click;
      }
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.button {
  padding: 2px;
  margin: 1px;
  box-sizing: unset;
  border-radius: 4px;
  width: 0.9em;
  height: 0.9em;
  text-align: center;
  font-size: 0.9em;
  border: 2px solid var(--foreground);
  background: var(--c5);
}

.button:hover {
  background: var(--c5);
}

.button:focus {
  background: var(--c5);
}

.button:disabled {
  opacity: 0.5;
  color: var(--c4);
  background: var(--c2);
}
</style>
        
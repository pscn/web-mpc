<template>
  <button class="button" :style="style" @click="$emit('click')">{{text}}</button>
</template>

<script>
export default {
  name: "ButtonText",
  props: {
    text: String,
    area: String,
    borderColor: String,
    color: String,
    background: String
  },
  computed: {
    style: function() {
      var result = {};
      if (this.border) result["border"] = "2px solid" + this.borderColor;
      if (this.background) result["background"] = this.background;
      if (this.area) result["gridArea"] = this.area;
      return result;
    }
  },
  components: {},
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
<!-- FIXME: use same style us ButtonIcon? -->
<style scoped>
.button {
  margin: 2px;
  box-sizing: unset;
  border-radius: 4px;
  text-align: center;
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
        
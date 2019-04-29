<template>
  <div class="playerControl" :style="style">
    <button-icon
      v-if="status=='stop'"
      class="item"
      icon="play"
      title="Play"
      icon-fill="black"
      area="play"
      @click="command('play','')"
    />
    <button-icon
      v-if="status=='pause'"
      class="item"
      icon="play"
      title="Resume"
      icon-fill="black"
      area="play"
      @click="command('resume','')"
    />
    <button-icon
      v-if="status=='play'"
      class="item"
      icon="pause"
      title="Pause"
      icon-fill="black"
      area="play"
      @click="command('pause','')"
    />
    <button-icon
      :disabled="status=='stop'"
      class="item"
      icon="stop"
      title="Stop"
      icon-fill="black"
      area="stop"
      @click="command('stop','')"
    />
    <button-icon
      class="item"
      icon="prev"
      title="Previous"
      icon-fill="black"
      area="prev"
      @click="command('previous','')"
    />
    <button-icon
      class="item"
      icon="next"
      title="Next"
      icon-fill="black"
      area="next"
      @click="command('next','')"
    />
    <button-icon
      :disabled="position==-1"
      class="item"
      icon="ban"
      title="Remove"
      area="remove"
      @click="command('remove',position)"
    />
    <button-icon
      :disabled="queueLength==0"
      class="item"
      icon="clean"
      title="Empty queue"
      area="clean"
      @click="command('clean','')"
    />
  </div>
</template>

<script>
import ButtonIcon from "./ButtonIcon.vue";

export default {
  name: "PlayerControl",
  props: {
    status: String,
    command: Function,
    area: String,
    position: Number,
    queueLength: Number
  },
  components: {
    ButtonIcon
  },
  computed: {
    style: function() {
      var result = {};
      if (this.area) {
        result["gridArea"] = this.area;
      }
      return result;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.item {
  min-width: 0;
  overflow: hidden;
  max-width: 1em;
}
.playerControl {
  display: grid;
  grid-template-columns: auto auto auto auto auto auto 1fr;
  grid-template-areas: "play stop prev next remove clean .";
  grid-gap: 2px;
  min-height: 0;
  min-width: 0;
}
</style>
        
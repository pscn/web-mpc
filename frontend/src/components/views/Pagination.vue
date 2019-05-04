<template>
  <div v-if="lastPage>1" class="pagination" :style="style">
    <div class="firstPage">
      <button-text :disabled="page<=1" v-on:click="$emit('goto', 1)" text="1"/>
    </div>
    <div class="previousPage">
      <button-text
        v-if="page!=1"
        :disabled="page<=1"
        v-on:click="$emit('goto', page-1)"
        :text="(page-1).toString()"
      />
    </div>
    <div id="currentPage" class="currentPage">{{page}}/{{lastPage}}</div>
    <div class="nextPage">
      <button-text
        v-if="(page+1)!=lastPage"
        :disabled="page>=lastPage"
        v-on:click="$emit('goto', page+1)"
        :text="(page+1).toString()"
      />
    </div>
    <div class="lastPage">
      <button-text
        :disabled="page>=lastPage"
        v-on:click="$emit('goto', lastPage)"
        :text="lastPage.toString()"
      />
    </div>
  </div>
</template>

<script>
import ButtonText from "../ButtonText.vue";

export default {
  name: "Pagination",
  props: {
    area: String,
    page: Number,
    lastPage: Number
  },
  components: { ButtonText },
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
.firstPage {
  grid-area: firstPage;
  text-align: left;
}
.previousPage {
  grid-area: previousPage;
  text-align: left;
}
.currentPage {
  grid-area: currentPage;
  align-self: center;
  text-align: center;
}
.nextPage {
  grid-area: nextPage;
  text-align: right;
}
.lastPage {
  grid-area: lastPage;
  text-align: right;
}
.pagination {
  display: grid;
  grid-template-columns: auto auto 1fr auto auto;
  grid-template-areas: "firstPage previousPage currentPage nextPage lastPage";
}
</style>
        
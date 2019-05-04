import Vue from "vue";
import App from "./App.vue";
// provide for "rootless" components aka components with multiple root nodes
// aka fragments
import Fragment from "vue-fragment";
Vue.use(Fragment.Plugin);

Vue.config.productionTip = false;

new Vue({
  render: h => h(App)
}).$mount("#app");

// eof

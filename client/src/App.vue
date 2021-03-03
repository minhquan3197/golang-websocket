<template>
  <div>
    <Header :randomName="state.randomName" />
    <HomePage title="Test Chat" />
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, reactive } from "vue";
import HomePage from "./pages/HomePage.vue";
import Header from "./components/common/Header.vue";

export default defineComponent({
  name: "App",
  components: {
    HomePage,
    Header,
  },
  setup: () => {
    const state = reactive({
      randomName: computed(() => {
        return "User " + new Date().getTime().toString();
      }),
    });

    return {
      state,
    };
  },
});
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
<template>
  <div>
    <Header :randomName="state.randomName" />
    <HomePage title="Test Chat" />
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, reactive, onMounted } from "vue";
import HomePage from "./pages/HomePage.vue";
import Header from "./components/common/Header.vue";
import { Config } from "./config/config";

export default defineComponent({
  name: "App",
  components: {
    HomePage,
    Header,
  },
  setup: () => {
    const state = reactive({
      connection: null,
      randomName: computed(() => {
        return "User " + new Date().getTime().toString();
      }),
    });

    onMounted(() => {
      console.log("Starting connection to WebSocket Server");
      state.connection = new WebSocket(Config.websocketUrl + "/ws");
      state.connection.onmessage = function (event) {
        console.log(event.data, "bbbbbbbb");
      };

      state.connection.onopen = function (event) {
        console.log("Connected");
      };

      setInterval(function () {
        state.connection.send("Hello, Server!");
      }, 1000);
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
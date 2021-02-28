<template>
  <div>
    <h1 class="capitalize text-2xl">{{ title }}</h1>
    <div class="flex">
      <Card class="w-1/4" />
      <Chat class="w-3/4" />
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import Card from "../components/cards/Card.vue";
import Chat from "../components/chat/Chat.vue";
export default defineComponent({
  name: "HomePage",
  components: {
    Card,
    Chat,
  },
  props: {
    title: {
      type: String,
      required: true,
    },
  },
});
</script>
<style scoped>
</style>

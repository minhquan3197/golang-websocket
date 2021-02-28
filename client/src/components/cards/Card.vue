<template>
  <div class="holder overflow-scroll">
    <div
      class="card border w-80 hover:shadow-none relative flex flex-col shadow-lg m-5z mb-5"
      v-for="user in state.users"
      :key="user.name"
    >
      <img
        class="max-h-20 w-full opacity-80 absolute top-0 background"
        style="z-index: -1"
        src="../../assets/images/image.jpg"
        alt=""
      />
      <div class="profile w-full flex m-3 ml-4 text-white">
        <img
          class="w-28 h-28 p-1 bg-white rounded-full"
          src="../../assets/images/avatar.jpg"
          alt=""
        />
        <div class="title mt-11 ml-3 font-bold flex flex-col">
          <div class="name break-words">Sarah</div>
          <!--  add [dark] class for bright background -->
          <div class="add font-semibold text-sm italic dark">Model</div>
        </div>
      </div>
      <div
        class="buttons flex absolute bottom-0 font-bold right-0 text-xs text-gray-500 space-x-0 my-3.5 mr-3"
      >
        <div
          class="add border rounded-l-2xl rounded-r-sm border-gray-300 p-1 px-4 cursor-pointer hover:bg-gray-700 hover:text-white"
        >
          Contact
        </div>
        <!-- <div class="add border rounded-r-2xl rounded-l-sm border-gray-300 p-1 px-4 cursor-pointer hover:bg-gray-700 hover:text-white">Bio</div> -->
      </div>
    </div>
    <!-- card end -->
  </div>
</template>
<script lang="ts">
import { computed, defineComponent, reactive, ref } from "vue";
export default defineComponent({
  name: "Card",
  setup: () => {
    const state = reactive({
      users: computed(() => {
        return [
          {
            name: "Test 1",
          },
          { name: "Test 2" },
          { name: "Test 3" },
          { name: "Test 4" },
          { name: "Test 5" },
          { name: "Test 6" },
        ];
      }),
    });
    return {
      state,
    };
  },
});
</script>
<style scoped>
.background {
  object-fit: cover;
}
.holder {
  height: 75vh;
}
</style>
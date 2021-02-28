<template>
  <!-- component -->
  <nav
    class="bg-white w-full flex relative justify-between items-center mx-auto px-8 h-20"
  >
    <!-- logo -->
    <div class="inline-flex">
      <a class="_o6689fn" href="/">
        <img class="w-12" alt="Vue logo" src="../../assets/images/logo.png" />
      </a>
    </div>

    <!-- end logo -->

    <!-- login -->
    <div class="flex-initial">
      <div class="flex justify-end items-center relative">
        <div class="block">
          <div class="inline relative">
            <button
              type="button"
              class="inline-flex items-center relative rounded-full hover:shadow-lg"
            >
              <div class="pl-2">
                <b>{{ randomName }}</b>
              </div>
              <div class="block flex-grow-0 flex-shrink-0 h-10 w-12 pl-5">
                <svg
                  viewBox="0 0 32 32"
                  xmlns="http://www.w3.org/2000/svg"
                  aria-hidden="true"
                  role="presentation"
                  focusable="false"
                  style="
                    display: block;
                    height: 100%;
                    width: 100%;
                    fill: currentcolor;
                  "
                >
                  <path
                    d="m16 .7c-8.437 0-15.3 6.863-15.3 15.3s6.863 15.3 15.3 15.3 15.3-6.863 15.3-15.3-6.863-15.3-15.3-15.3zm0 28c-4.021 0-7.605-1.884-9.933-4.81a12.425 12.425 0 0 1 6.451-4.4 6.507 6.507 0 0 1 -3.018-5.49c0-3.584 2.916-6.5 6.5-6.5s6.5 2.916 6.5 6.5a6.513 6.513 0 0 1 -3.019 5.491 12.42 12.42 0 0 1 6.452 4.4c-2.328 2.925-5.912 4.809-9.933 4.809z"
                  ></path>
                </svg>
              </div>
            </button>
          </div>
        </div>
      </div>
    </div>
    <!-- end login -->
  </nav>
</template>
<script lang="ts">
import { defineComponent } from "vue";
export default defineComponent({
  name: "Header",
  props: {
    randomName: {
      type: String,
      required: true,
    },
  },
});
</script>
<style scoped>
</style>
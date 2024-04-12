<script setup lang="ts">
import RecycleBin from '@/pages/home/main-nav/recycle-bin.vue';
</script>

<template>
  <div class="garbage">
    <RecycleBin :show-title="false"></RecycleBin>
  </div>
</template>

<style scoped lang="scss">
  .garbage >>> .w-full {
    display: none;
  }
</style>
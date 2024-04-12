<script setup lang="ts">
import UpdateLog from '@/components/update-log/index.vue';
import { NScrollbar } from 'naive-ui';
</script>

<template>
  <n-scrollbar style="height: 100%">
    <div id="dynamic-bar" class="dynamic-bar px-4 pt-3 animate__animated animate__fadeIn faster">
      <UpdateLog />
    </div>
  </n-scrollbar>

</template>

<style scoped lang="scss">

</style>
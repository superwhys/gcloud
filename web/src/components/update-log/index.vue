<template>
  <div class="text-sm">
    <div
      class="mb-3 flex justify-between hover:text-secondary"
      v-for="item in data"
      :key="item.content"
    >
      <div>
        <span class="text-primary">「{{ item.type }}」</span>{{ item.content }}
      </div>
      <div class="text-primary ml-2" style="min-width: 80px">{{ item.date }}</div>
    </div>
  </div>
</template>

<script setup lang="ts">
const data = [
  {
    type: '✨ 功能',
    content: '悄咪咪上线图集 /doge',
    date: '2024-03-14',
  },
  {
    type: '🦄 新增',
    content: '支持预览office文件啦~',
    date: '2024-03-08',
  },
  {
    type: '🌈 优化',
    content: '优化请求提示信息',
    date: '2024-03-01',
  },
  {
    type: '🦄 新增',
    content: '可以给你喜欢的帖子点赞+收藏啦~',
    date: '2024-03-01',
  },
  {
    type: '🦄 新增',
    content: '评论功能已上线~',
    date: '2024-02-23',
  },
  {
    type: '🦄 新增',
    content: 'Golang记录论坛可以发贴了~',
    date: '2024-02-14',
  },
  {
    type: '🦄 新增',
    content: '用户分享历史查询',
    date: '2024-02-04',
  },
  {
    type: '🦄 新增',
    content: '现在，你可以保存别人的分享文件辣~',
    date: '2024-01-30',
  },
  {
    type: '🦄 新增',
    content: '热门分享列表',
    date: '2024-01-28',
  },
  {
    type: '🦄 新增',
    content: '分享功能；分享详情功能；分享数据统计功能',
    date: '2024-01-27',
  },
  {
    type: '🦄 新增',
    content: '添加回收站列表',
    date: '2024-01-25',
  },
  {
    type: '🔨 重构',
    content: '重构使用说明模块，添加提示按钮',
    date: '2024-01-21',
  },
  {
    type: '✨ 功能',
    content: '新增预览文件功能（文本文件预览，不支持office文件）',
    date: '2024-01-20',
  },
  {
    type: '🐞 修复',
    content: '修复已知问题',
    date: '2024-01-19',
  },
  {
    type: '✨ 功能',
    content: '修复已知问题；禁止选中文件夹文本',
    date: '2024-01-13',
  },
  {
    type: '🐞 修复',
    content: '修复注册成功/失败跳转问题',
    date: '2024-01-01',
  },
  {
    type: '✨ 功能',
    content: '添加移动端访问受限提示',
    date: '2023-12-24',
  },
  {
    type: '✨ 功能',
    content: '新增加入天数显示',
    date: '2023-12-22',
  },
  {
    type: '✨ 功能',
    content: '新增Not Found页面',
    date: '2023-12-21',
  },
  {
    type: '🚀 维护',
    content: '新增日志查看; 新增视频播放功能; 修复图片预览畸变；其他...',
    date: '2023-12-20',
  },
];
</script>

<style lang="scss" scoped></style>

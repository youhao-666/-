<script setup lang="ts">
import Menu from '@/components/menu.vue';
import SelfInfo from './selfInfo.vue';

</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <Menu />
      </el-aside>
      <el-main>
        <SelfInfo />
      </el-main>
    </el-container>
  </div>
</template>

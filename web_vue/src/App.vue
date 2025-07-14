<script setup>
import headmenu from './components/headMenu.vue';
</script>

<template>
  <div>
    <el-container>
      <el-header>
        <div class="header">
          <headmenu></headmenu>
        </div>
      </el-header>
      <el-container>
        <el-container>
          <el-main class="main"><router-view/></el-main>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>



<style>
.header {
  margin-left: 36vw;
  color: #fff;
  align-items: center;
  height: 60px;
}

.main {
  width: 100%;
  margin-top: 7.2vh;
}

.footer {
  color: #000000;
  height: 60px;
  line-height: 60px;
}
</style>

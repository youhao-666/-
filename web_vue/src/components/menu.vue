<script lang="ts" setup>
import { computed } from 'vue';
import {
  Document,
  Menu as IconMenu,
  Location,
  Setting,
} from '@element-plus/icons-vue'
import { useRouter } from 'vue-router'

const props = defineProps({
  msg: String
})
const defaultActive = computed(() => {
    return props.msg ? "/manageCenter/articleManage" : "/manageCenter/selfInfo";
});

const router = useRouter()
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  router.push(key)
}
</script>

<template>
  <el-col>
    <el-menu :default-active="defaultActive" class="el-menu-vertical-demo" @select="handleSelect">
      <el-menu-item index="/manageCenter/selfInfo">
        <el-icon><Document /></el-icon>
        <span>个人信息</span>
      </el-menu-item>

      <el-menu-item index="/manageCenter/articleManage">
        <el-icon>
          <setting />
        </el-icon>
        <span>我的文章</span>
      </el-menu-item>
    </el-menu>
  </el-col>
</template>

<style>
.el-menu-vertical-demo {
  width: 200rw;
  height: 100%;
  background-color: #545c64;
}

.routerLink {
  color: rgb(0, 0, 0);
  text-decoration: none;
}
</style>

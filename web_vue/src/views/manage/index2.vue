<script setup lang="ts">
import Menu from '@/components/adminMenu.vue';
import UserManage from './userManage.vue';
import request from '@/utils/request';
import { message } from 'ant-design-vue';
import { ref } from 'vue';
const tableData = ref([]);
const handleRefresh = () => {
  request({
    url: '/admin/user/list',
    method: 'GET',
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    }
  }).then((res) => {
    console.log(res)
    if (res.status === 200) {
      message.success('获取用户成功')
      tableData.value = res.data.data;
      console.log(tableData.value)
    } else {
      message.success('获取用户失败')
    }
  })
}
handleRefresh();
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <Menu msg="用户管理" />
      </el-aside>
      <el-main>
        <UserManage :props="tableData" @refresh="handleRefresh" />
      </el-main>
    </el-container>
  </div>
</template>

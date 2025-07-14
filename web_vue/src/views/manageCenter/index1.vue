<script setup lang="ts">
import Menu from '@/components/menu.vue';
import ArticleManage from './articleManage.vue';
import { ref } from 'vue';
import request from '@/utils/request';
import { message } from 'ant-design-vue';
const tableData = ref([]);
const handleRefresh = () => {
  request({
    url: `/article/listByUser/${localStorage.getItem('id')}`,
    method: 'GET',
  }).then((res) => {
    console.log(res)
    if (res.status === 200) {
      message.success('获取文章成功')
      tableData.value = res.data.data;
      console.log(tableData.value)
    } else {
      message.success('获取文章失败')
    }
  })
}
handleRefresh();
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-aside width="200px">
        <Menu msg="manageCenter" />
      </el-aside>
      <el-main>
        <ArticleManage :props="tableData" @refresh="handleRefresh" />
      </el-main>
    </el-container>
  </div>
</template>

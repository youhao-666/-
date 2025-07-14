<script setup lang="ts">
import { ref } from 'vue';
import Menu from '@/components/adminMenu.vue';
import { message } from 'ant-design-vue';
import TagManage from '@/views/manage/tagManage.vue';
import request from '@/utils/request';
const tableData = ref([]);
const handleRefresh = () => {
    request({
        url: '/tag/list',
        method: 'GET',
    }).then((res) => {
        // console.log(res)
        if (res.status === 200) {
            message.success('获取标签成功')
            tableData.value = res.data.data;
            console.log(tableData.value)
        } else {
            message.success('获取标签失败')
        }
    })
}
handleRefresh();

</script>

<template>
    <div class="common-layout">
        <el-container>
            <el-aside width="200px">
                <Menu msg="标签管理" />
            </el-aside>
            <el-main>
                <TagManage :props="tableData"  @refresh="handleRefresh" />
            </el-main>
        </el-container>
    </div>
</template>

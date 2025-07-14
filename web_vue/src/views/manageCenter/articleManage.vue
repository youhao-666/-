<script lang="ts" setup>
import request from '@/utils/request'
import { message } from 'ant-design-vue'

const props = defineProps({
    props: {
        type: Array,
        default: () => []
    }
})

const emit = defineEmits(['edit', 'delete'])
const handleDelete = (index: any) => {
    console.log(index)
    request({
        url: `/admin/article/delete/${index}`,
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    }).then((res) => {
        console.log(res)
        if (res.status === 200) {
            message.success('删除成功')
            emit('refresh')
            // window.location.reload()
        } else {
            message.error('删除失败')
        }
    }
    )
}

</script>

<template>
    <el-table :data="props.props" style="width: 100%" height="500px">
        <el-table-column label="ID" prop="id" align="center" />
        <el-table-column label="标题" prop="title" align="center" />
        <el-table-column label="作者ID" prop="userId" align="center" />
        <el-table-column label="标签" prop="tags" align="center" />
        <el-table-column label="操作" align="center">
            <template #default="scope">
                <el-popconfirm title="你确定要删除该文章吗？" ok-text="确定" cancel-text="取消" @confirm="handleDelete(scope.row.id)">
                    <template #reference>
                        <el-button type="danger">删除</el-button>
                    </template>
                </el-popconfirm>
            </template>
        </el-table-column>
    </el-table>
</template>

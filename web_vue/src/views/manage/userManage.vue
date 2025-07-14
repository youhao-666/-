<script lang="ts" setup>
import { computed, ref, reactive, onMounted } from 'vue'
import { message } from 'ant-design-vue'
import request from '@/utils/request'
import dayjs from 'dayjs'

const formLabelWidth = '70px'
const isAdd = ref(true)
const Visible = ref(false)
const props = defineProps({
    props: {
        type: Array,
        default: () => []
    }
})

const emit = defineEmits(['edit', 'delete'])
const tagList = ref(props.props)
const search = ref('')

const formatDate = (dateStr: string) => {
    return dayjs(dateStr).format('YYYY-MM-DD HH:mm:ss')
}

const form = reactive({
    id: '',
    account: '',
    nickname: '',
    password: '',
})

// 记录当前编辑的行数据
const currentRow = ref(null)

// 点击编辑按钮时的处理函数
const handleEdit = (row: any) => {
    Visible.value = true
    isAdd.value = false
    currentRow.value = row
    form.id = row.ID
    form.account = row.account
    form.nickname = row.nickname
}

const handleDelete = (index: any) => {
    console.log(index)
    request({
        url: `/admin/delete/${index}`,
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    }).then((res) => {
        console.log(res)
        if (res.status === 200) {
            message.success('删除成功')
            emit('refresh')
            //  window.location.reload()
        } else {
            message.error('删除失败')
        }
    }
    )
}

const handleResetPassword = (index: any) => {
    console.log(index)
    request({
        url: `/admin/Reset/${index}`,
        method: 'PUT',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    }).then((res) => {
        console.log(res)
        if (res.status === 200) {
            message.success('重置密码成功')
            emit('refresh')
            //  window.location.reload()
        } else {
            message.error('重置密码失败')
        }
    }
    )
}


const submit = () => {
    Visible.value = false
    const json = JSON.parse(JSON.stringify({
        password: form.password,
        account: form.account,
        nickname: form.nickname,
    }
    ))
    console.log(form);
    const formData = new FormData()
    formData.append('data', JSON.stringify(json))
    request({
        url: `/user/update/${form.id}`,
        method: 'PUT',
        data: formData,
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    }).then((res) => {
        console.log(res)
        if (res.status === 200) {
            //  window.location.reload()
            emit('refresh')
            message.success('修改成功')
        } else {
            message.error('修改失败')
        }
    })
}

// 关闭对话框时清空表单
const closeDialog = () => {
    form.id = ''
    form.account = ''
    form.nickname = ''
    form.password = ''
    currentRow.value = null
}

// 监听props变化，更新tagList
onMounted(() => {
    tagList.value = props.props
})
</script>

<template>
    <el-table :data="props.props" style="width: 100%" height="500px">
        <el-table-column label="ID" prop="ID" align="center" />
        <el-table-column label="用户名" prop="account" align="center" />
        <el-table-column label="昵称" prop="nickname" align="center" />
        <el-table-column label="更新时间" align="center">
            <template #default="scope">
                {{ formatDate(scope.row.UpdatedAt) }}
            </template>
        </el-table-column>
        <el-table-column label="操作" align="center">
            <template #default="scope">
                <el-button size="small" @click="handleEdit(scope.row)">
                    编辑
                </el-button>
                <el-popconfirm title="你确定要删除该标签吗？" ok-text="确定" cancel-text="取消" @confirm="handleDelete(scope.row.ID)">
                    <template #reference>
                        <el-button size="small" type="danger">删除</el-button>
                    </template>
                </el-popconfirm>
                <el-popconfirm title="你确定要重置密码吗？" ok-text="确定" cancel-text="取消" @confirm="handleResetPassword(scope.row.ID)">
                    <template #reference>
                        <el-button size="small" type="danger">重置密码</el-button>
                    </template>
                </el-popconfirm>
            </template>
        </el-table-column>
    </el-table>

    <el-dialog v-model="Visible" title='编辑用户' width="500" @close="closeDialog">
        <el-form :model="form">
            <el-form-item label="用户名" :label-width="formLabelWidth" :rules="[{ required: true, message: '请输入用户名' }]">
                <el-input v-model="form.account" autocomplete="off" />
            </el-form-item>
            <el-form-item label="昵称" :label-width="formLabelWidth" :rules="[{ required: true, message: '请输入昵称' }]">
                <el-input v-model="form.nickname" autocomplete="off" />
            </el-form-item>
            <el-form-item label="密码" :label-width="formLabelWidth">
                <el-input v-model="form.password" type="password" placeholder="tips: 留空则不修改密码" autocomplete="off" />
            </el-form-item>
        </el-form>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="Visible = false">取消</el-button>
                <el-button type="primary" @click="submit">
                    确定
                </el-button>
            </div>
        </template>
    </el-dialog>
</template>
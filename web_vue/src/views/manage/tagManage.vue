<script lang="ts" setup>
import { computed, ref, reactive } from 'vue'
import request from '@/utils/request'
import { message } from 'ant-design-vue'
import dayjs from 'dayjs'

const formLabelWidth = '70px'
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
    ID: '',
    tagName: '',
})

const handleDelete = (index: any) => {
    console.log(index)
    request({
        url: `/admin/tag/delete/${index}`,
        method: 'DELETE',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    }).then((res) => {
        console.log(res)
        if (res.status === 200) {
            message.success('删除成功')
            // window.location.reload()
            emit('refresh')
        } else {
            message.error('删除失败')
        }
    }
    )

}

const submit = () => {
    Visible.value = false
    console.log(form);
    const json = JSON.parse(JSON.stringify({
        tagName: form.tagName,
    }
    ))
    const formData = new FormData()
    formData.append('data', JSON.stringify(json))
    request({
        url: `/admin/tag/create`,
        method: 'POST',
        data: formData,
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    }).then((res) => {
        console.log(res)
        if (res.status === 200) {
            // window.location.reload()
            emit('refresh')
            message.success('修改成功')
        } else {
            message.error('修改失败')
        }
    })
}

</script>

<template>
    <div>
        <el-button type="primary" @click="Visible = true"
            :style="{ right: '180px', top: '110px', position: 'fixed' }">新增标签</el-button>
    </div>
    <el-table :data="props.props" style="width: 100%" height="500px">
        <el-table-column label="ID" prop="id" align="center" />
        <el-table-column label="标签名称" prop="tagName" align="center" />
        <el-table-column label="更新时间" align="center">
            <template #default="scope">
                {{ formatDate(scope.row.UpdatedAt) }}
            </template>
        </el-table-column>
        <el-table-column label="操作" align="center">
            <!-- <template #header>
                <el-input v-model="search" size="small" placeholder="搜索标签" />
            </template> -->
            <template #default="scope">
                <el-popconfirm title="你确定要删除该标签吗？" ok-text="确定" cancel-text="取消" @confirm="handleDelete(scope.row.id)">
                    <template #reference>
                        <el-button type="danger">删除</el-button>
                    </template>
                </el-popconfirm>
            </template>
        </el-table-column>
    </el-table>



    <el-dialog v-model="Visible" title="添加标签" width="500">
        <el-form :model="form">
            <el-form-item label="标签名称" :label-width="formLabelWidth">
                <el-input v-model="form.tagName" autocomplete="off" />
            </el-form-item>
            <!-- <el-form-item label="Zones" :label-width="formLabelWidth">
        <el-select v-model="form.region" placeholder="Please select a zone" >
          <el-option label="Zone No.1" value="shanghai" />
          <el-option label="Zone No.2" value="beijing" />
        </el-select>
      </el-form-item> -->
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

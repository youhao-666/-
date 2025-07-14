<script lang="ts" setup>
import { reactive, ref } from 'vue'
import request from '@/utils/request'
import { message } from 'ant-design-vue'
import type { FormInstance, FormRules } from 'element-plus'

const form = reactive({
    nickName: '',
    signature: '',
    realName: '',
    school: '',
    class: '',
})
request({
    url: `/user/${localStorage.getItem('id')}`,
    method: 'GET',
    headers: {
        'Authorization': `Bearer ${localStorage.getItem('token')}`
    }
}).then(res => {
    console.log(res);
    if (res.status === 200) {
        form.nickName = res.data.data.nickname;
        message.success('获取用户信息成功');
    } else {
        message.error('获取用户信息失败');
    }
})



const submitForm = () => {
    console.log(form)

}


</script>

<template>
    <h2>个人信息</h2>
    <el-form ref="ruleFormRef" style="max-width: 600px" :model="form" status-icon label-width="auto"
        class="demo-ruleForm">
        <el-form-item label="社交昵称">
            <el-input v-model="form.nickName" autocomplete="off" />
        </el-form-item>
        <el-form-item label="个性签名">
            <el-input v-model="form.signature" />
        </el-form-item>
        <el-form-item label="真实姓名">
            <el-input v-model="form.realName" />
        </el-form-item>
        <el-form-item label="所在学校">
            <el-input v-model="form.school" />
        </el-form-item>
        <el-form-item label="所在班级">
            <el-input v-model="form.class" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="submitForm()">
                提交
            </el-button>
        </el-form-item>
    </el-form>
</template>

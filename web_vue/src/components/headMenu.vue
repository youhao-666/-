<script lang="ts" setup>
import { message } from 'ant-design-vue';
import request from '../utils/request';
import {
    Document,
    Menu as IconMenu,
    Plus,
    Setting,
    ZoomIn
} from '@element-plus/icons-vue'
import type { UploadFile } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
const isUser = ref(true)  // 是否是用户
const Visible = ref(false)  // 菜单是否显示
const isLogin = ref(false)  // 是否登录了
const login = ref(true)     // 登录还是注册
const userName = ref('') // 用户名
const dialogVisible = ref(false)
const dialogImageUrl = ref('')
const uploadFile = ref<UploadFile | null>(null)
const formLabelWidth = '65px'
const router = useRouter()

const handleSelect = (key: string, keyPath: string[]) => {
    if (!isLogin.value) {
        message.error('请先登录')
        Visible.value = true
        return
    }
    console.log(key, keyPath)
    router.push(key)
}
const gotoInfo = () => {
    router.push('/info')
}
const exitLogin = () => {
    isLogin.value = false
    isUser.value = true
    login.value = true
    message.success('退出登录成功')
    router.push('/')
}
const form = reactive({
    nickname: '',
    identify: 0,
    password: '',
    account: '',
    avatar: '',
})

const goLogin = () => {
    const json = JSON.parse(JSON.stringify({
        password: form.password,
        account: form.account,
    }
    ))
    console.log(json)
    const formData = new FormData()
    formData.append('data', JSON.stringify(json))
    const url = form.identify === 0 ? "/user/login" : "/admin/login";
    request({
        method: "POST",
        url: url,
        data: formData,
    }).then((res) => {
        console.log(res);
        console.log(res.data.data.token);
        localStorage.setItem('id', res.data.data.id);
        localStorage.setItem('token', res.data.data.token);
        userName.value = res.data.data.nickname;
        message.success('登录成功');
        isLogin.value = true;
        login.value = true;
        isUser.value = form.identify === 0 ? true : false;
        router.push('/');
        Visible.value = false;
    }).catch((error) => {
        console.error(error);
        message.error('登录失败');
    });
}

const goRegister = (file: UploadFile) => {
    const json = JSON.parse(JSON.stringify({
        password: form.password,
        account: form.account,
        nickname: form.nickname,
    }
    ))
    console.log(json);
    const formData = new FormData()
    formData.append('data', JSON.stringify(json))
    // for (let [key, value] of formData.entries()) {
    //     console.log(key, value);
    // }
    console.log("FormData", formData)
    request({
        method: "POST",
        url: "/user/create",
        data: formData,
    }).then((res) => {
        console.log(res)
        message.success('注册成功')
    })
    login.value = true
    Visible.value = false
}

// const handlePictureCardPreview = (file: UploadFile) => {
//     dialogImageUrl.value = file.url!
//     dialogVisible.value = true
//     console.log(file)
//     console.log(dialogImageUrl.value)
// }

// const handleFileChange = (file: UploadFile) => {
//     // 当文件状态为"ready"（已选中未上传）时，生成临时预览URL
//     if (file.status === 'ready') {
//         // 将文件转为本地可访问的URL（仅用于预览）
//         file.url = URL.createObjectURL(file.raw!)
//         uploadFile.value = file  // 保存文件信息
//         form.avatar = file.url  // 更新表单中的头像地址
//     }
// }

</script>


<template>
    <div :style="{ position: 'absolute', top: '2vh', left: '10vw', width: '100%', color: '#000', fontSize: '1.5vw' }">
        ZSDN</div>
    <el-menu default-active="/" class="el-menu-demo" mode="horizontal" @select="handleSelect">
        <el-menu-item index="/">
            <el-icon>
                <Document />
            </el-icon>
            <span>首页</span>
        </el-menu-item>
        <el-menu-item index="/article" :style="{ display: isUser ? 'block' : 'none' }">
            <el-icon>
                <Document />
            </el-icon>
            发布文章
        </el-menu-item>
        <el-menu-item index="/manageCenter/selfInfo" :style="{ display: isUser ? 'block' : 'none' }">
            <el-icon>
                <Setting />
            </el-icon>
            创作者中心
        </el-menu-item>
        <el-menu-item index="/manage/articleManage" :style="{ display: isUser ? 'none' : 'block' }">
            <el-icon>
                <Setting />
            </el-icon>
            管理中心
        </el-menu-item>
    </el-menu>
    <button class="user" :style="{ display: isLogin ? 'none' : 'block' }" @click="Visible = true">登录</button>
    <div :style="{ display: isLogin ? 'block' : 'none', position: 'absolute', top: 0, right: 0 }">
        <el-dropdown trigger="click">
            <img src="@/assets/logo.png" alt="user" class="user-img" />
            <template #dropdown>
                <el-dropdown-menu>
                    <el-dropdown-item @click="gotoInfo">个人信息</el-dropdown-item>
                    <el-dropdown-item @click="exitLogin">退出登录</el-dropdown-item>
                </el-dropdown-menu>
            </template>
        </el-dropdown>
        <div class="user-name">{{ userName }}</div>
    </div>

    <el-dialog v-model="Visible" :title="login ? '用户登录' : '用户注册'" width="500" draggable="true">
        <el-form :model="form" inline>
            <el-form-item label="用户名" :rules="[{ required: true, message: '请输入用户名' }]" :label-width="formLabelWidth">
                <el-input v-model="form.account" autocomplete="off" style="width: 390px" />
            </el-form-item>
            <el-form-item label="密码" :rules="[{ required: true, message: '请输入密码' }]" :label-width="formLabelWidth">
                <el-input v-model="form.password" type="password" autocomplete="off" style="width: 390px" />
            </el-form-item>
            <el-form-item label="昵称" :style="{ display: login ? 'none' : 'inline-flex' }" :label-width="formLabelWidth"
                :rules="[{ required: true, message: '请输入昵称' }]">
                <el-input v-model="form.nickname" autocomplete="off" style="width: 390px" />
            </el-form-item>
            <!-- <el-form-item label="头像" :style="{ display: login ? 'none' : 'inline-flex' }" :label-width="formLabelWidth">
                <el-upload action="/user/create" :limit="1" list-type="picture-card" :auto-upload="false"
                    @change="handleFileChange" ref="upload">
                    <el-icon>
                        <Plus />
                    </el-icon>
                    <template #file="{ file }">
                        <div>
                            <img class="el-upload-list__item-thumbnail" :src="file.url" alt="" />
                            <span class="el-upload-list__item-actions">
                                <span class="el-upload-list__item-preview" @click="handlePictureCardPreview(file)">
                                    <el-icon>
                                        <ZoomIn />
                                    </el-icon>
                                </span>
                            </span>
                        </div>
                    </template>
                </el-upload>
            </el-form-item> -->
            <el-form-item label="类型" :style="{ display: login ? 'inline-flex' : 'none' }" :label-width="formLabelWidth"
                :rules="[{ required: true, message: '请选择类型' }]">
                <el-radio-group v-model="form.identify">
                    <el-radio :value=0>用户</el-radio>
                    <el-radio :value=1>管理员</el-radio>
                </el-radio-group>
            </el-form-item>
        </el-form>
        <template #footer>
            <div class="dialog-footer">
                <div :style="{ display: login ? 'block' : 'none' }">
                    <el-button type="primary" @click="goLogin">
                        登录
                    </el-button>
                    <el-button @click="login = false">注册</el-button>
                </div>
                <div :style="{ display: login ? 'none' : 'block' }">
                    <el-button @click="login = true">
                        返回登录
                    </el-button>
                    <el-button type="primary" @click="goRegister">注册</el-button>
                </div>
            </div>
        </template>

        <el-dialog v-model="dialogVisible">
            <img w-full :src="dialogImageUrl" alt="Preview Image" />
        </el-dialog>
    </el-dialog>




</template>

<style>
.user {
    border: none;
    background-color: transparent;
    cursor: pointer;
    position: absolute;
    right: 5vw;
    top: 3vh;
    color: #3296e3;
    font-size: 14px;
}

.user-img {
    width: 25px;
    height: 25px;
    border-radius: 50%;
    position: absolute;
    right: 7.5vw;
    top: 4vh;
    cursor: pointer;
    transform: translate(-50%, -50%);
}

.user-name {
    position: absolute;
    width: 5vw;
    right: 3vw;
    top: 3vh;
    color: #3296e3;
    font-size: 14px;
}
</style>
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
    if (props.msg === '文章管理') {
        return '/manage/articleManage'
    } else if (props.msg === '标签管理') {
        return '/manage/tagManage'
    } else if (props.msg === '用户管理') {
        return '/manage/userManage'
    } else {
        return '/manage/articleManage'
    }
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
            <el-menu-item index="/manage/articleManage">
                <el-icon>
                    <Document />
                </el-icon>
                <span>文章管理</span>
            </el-menu-item>

            <el-menu-item index="/manage/tagManage">
                <el-icon>
                    <setting />
                </el-icon>
                <span>标签管理</span>
            </el-menu-item>

            <el-menu-item index="/manage/userManage">
                <el-icon>
                    <setting />
                </el-icon>
                <span>用户管理</span>
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

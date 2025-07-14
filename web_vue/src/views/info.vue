<script setup>
import { ref } from 'vue';
import { ArrowLeftBold } from '@element-plus/icons-vue'
import { message } from 'ant-design-vue'
import request from '@/utils/request';

const props = {};
request({
    url: `/user/${localStorage.getItem('id')}`,
    method: 'GET',
    headers: {
        'Authorization': `Bearer ${localStorage.getItem('token')}`
    }
}).then(res => {
    console.log(res);
    if (res.status === 200) {
        props.value = res.data.data;
        console.log(props.value);
        username.value = props.value.nickname;
        message.success('获取用户信息成功');
    } else {
        message.error('获取用户信息失败');
    }
})
const username = ref(props);
// const account = props.value.account;
const signature = ref('这个人很懒，什么都没有留下');
const followCount = ref(1);
const fanCount = ref(0);

</script>

<template>
    <el-icon :style="{ fontSize: '30px', position: 'absolute', left: '80px', top: '100px', cursor: 'pointer' }"
        @click="$router.back()">
        <ArrowLeftBold />
    </el-icon>
    <div class="user-info-container">
        <div class="user-card">
            <div class="avatar">
                <img src="@/assets/home.png" alt="avatar" />
            </div>
            <div class="user-info">
                <div class="username">{{ username }}</div>
                <div class="signature">{{ signature }}</div>
                <div class="stats">
                    <span class="stat-item">关注数：{{ followCount }}</span>
                    <span class="stat-item">粉丝数：{{ fanCount }}</span>
                </div>
            </div>
        </div>
    </div>
</template>

<style scoped>
.user-info-container {
    width: 80%;
    margin: 0 auto;
    padding: 20px;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0 0 8px rgba(0, 0, 0, 0.1);
}

.back-arrow {
    cursor: pointer;
    margin-bottom: 10px;
}

.user-card {
    display: flex;
    align-items: center;
}

.avatar {
    width: 80px;
    height: 80px;
    border-radius: 50%;
    overflow: hidden;
    margin-right: 20px;
}

.avatar img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.user-info {
    display: flex;
    flex-direction: column;
    justify-content: center;
}

.username {
    font-size: 18px;
    font-weight: bold;
    margin-bottom: 8px;
}

.signature {
    color: #999;
    margin-bottom: 8px;
}

.stats {
    display: flex;
    gap: 20px;
}

.stat-item {
    color: #666;
}
</style>
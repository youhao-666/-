<script lang="ts" setup>
import { reactive, ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import request from '@/utils/request'

const router = useRouter()
const input = ref('') // 查询标签
const items = ref([])
const originalItems = ref([]) // 保存原始数据用于搜索
const tags = ref([])
const activeTag = ref('全部')

const currentPage = ref(1)
const pageSize = ref(5)
const total = ref(0)


const displayedItems = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return items.value.slice(start, end)
})


request({
  url: '/article/quaryall',
  method: 'GET',
}).then((res) => {
  if (res.status === 200) {
    originalItems.value = res.data.data
    items.value = [...originalItems.value]
    total.value = items.value.length
  } else {
    message.error('获取文章失败')
  }
})


request({
  url: '/tag/list',
  method: 'GET',
}).then((res) => {
  if (res.status === 200) {
    tags.value = res.data.data
  } else {
    message.error('获取标签失败')
  }
})


const handleSelect = (record: any) => {
  localStorage.setItem('article', JSON.stringify(record))     //全球可飞
  router.push('/articleDetail')
}

const handleSelectTag = (tag: any) => {
  console.log(tag);
  activeTag.value = tag.tagName
  currentPage.value = 1
  input.value = ''                      //清空搜索框
  if (tag === 0) {
    request({
      url: '/article/quaryall',
      method: 'GET',
    }).then((res) => {
      if (res.status === 200) {
        originalItems.value = res.data.data
        items.value = [...originalItems.value]
        total.value = items.value.length
      } else {
        message.error('获取文章失败')
      }
    })
  } else {
    request({
      url: `/article/listByTag/${tag.id}`,
      method: 'GET',
    }).then((res) => {
      if (res.status === 200) {
        items.value = res.data.data
        originalItems.value = res.data.data
      } else {
        message.error('获取文章失败')
      }
      total.value = items.value.length
    })
  }
}

// 分页器页码改变事件
const handleCurrentChange = (page: number) => {
  currentPage.value = page
}

const handleSearch = () => {
  if (!input.value.trim()) {
    items.value = [...originalItems.value]
    return
  }

  const keyword = input.value.trim().toLowerCase()
  items.value = originalItems.value.filter(item =>
    item.title.toLowerCase().includes(keyword) ||
    item.content.toLowerCase().includes(keyword)
  )
  currentPage.value = 1
  total.value = items.value.length
}
const truncatedContent = (content) => {
  if (content.length > 35) {
    return content.slice(0, 35) + '...';
  }
  return content;
};

</script>

<template>
  <div class="tags-container">
    <div class="tags">
      <span class="tag" @click="handleSelectTag(0)" :class="{ 'active': activeTag === '全部' }">全部</span>
      <span class="tag" v-for="tag in tags" :key="tag.ID" :class="{ 'active': tag.tagName === activeTag }"
        @click="handleSelectTag(tag)">{{ tag.tagName }}</span>
    </div>
    <hr />
  </div>
  <div>
    <div class="head">
      <el-input v-model="input" style="width: 240px" placeholder="请输入搜索内容" />
      <el-button type="primary" @click="handleSearch"><el-icon>search</el-icon></el-button>
    </div>
    <div v-for="item in displayedItems" :key="item.id" class="item" @click="handleSelect(item)">
      <h3>{{ item.title }}</h3>
      <p>{{ truncatedContent(item.content) }}</p>
      <p>{{ item.author }}</p>
      <p>{{ item.time }}</p>
    </div>
    <div>
      <el-pagination background class="pagination" layout="prev, pager, next, jumper, total" :current-page="currentPage"
        :page-size="pageSize" :total="total" @current-change="handleCurrentChange" />
    </div>
  </div>
</template>

<style scoped>
/* 保持原有样式不变 */
.head {
  position: absolute;
  inset: 0px;
  background-image: url(@/assets/home.png);
  background-size: cover;
  background-position: center center;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 10vh;
  top: 7.2vh;
}

.item {
  width: 50vw;
  padding: 20px;
  margin: 3vh 10vw 0 10vw;
  width: 100%;
  border: 1px solid #ccc;
  border-radius: 2vw;
  max-width: 75%;
  background-color: #fff;
}

.item:hover {
  cursor: pointer;
  border: none;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.2);
}

.tags-container {
  width: 100%;
  overflow-x: auto;
  white-space: nowrap;
  padding: 8px 0;
}

.tags {
  display: flex;
  gap: 8px;
  justify-content: center;
}

.tag {
  padding: 6px 12px;
  background-color: #f2f2f2;
  border-radius: 4px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.tag.active {
  background-color: #409eff;
  color: #fff;
}

.tag:hover {
  background-color: #52a8ff;
  color: #fff;
}

.pagination {
  display: flex;
  justify-content: center;
  margin-top: 21px;
}
</style>
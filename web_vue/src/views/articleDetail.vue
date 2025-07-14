<script setup>
import { ref, onMounted } from 'vue';
import MarkdownIt from 'markdown-it';
import { ArrowLeftBold } from '@element-plus/icons-vue'

const md = new MarkdownIt();
const props = JSON.parse(localStorage.getItem('article')) || {};
console.log("props", props);

const title = ref(props.title);
const tags = ref(props.tags || []);
const content = ref(props.content || '');
const userId = ref(props.userId);
const markdownHtml = ref('');

onMounted(() => {
  // 将 Markdown 内容渲染为 HTML
  markdownHtml.value = md.render(content.value);
});
</script>

<template>
    <el-icon :style="{ fontSize: '30px', position: 'absolute', left: '80px', top: '100px', cursor: 'pointer' }"
        @click="$router.back()">
        <ArrowLeftBold />
    </el-icon>
    <div class="interview-record-container">
        <h1 class="title">{{ title }}</h1>
        <div class="tags">
            <span class="tag" v-for="tag in tags" :key="tag">{{ tag }}</span>
        </div>
        <hr />
        <section v-html="markdownHtml"></section>
    </div>
</template>

<style scoped>
.interview-record-container {
    width: 80%;
    margin: 0 auto;
    padding: 20px;
    font-family: Arial, sans-serif;
}

.title {
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 10px;
}

.tags {
    margin-bottom: 20px;
}

.tag {
    display: inline-block;
    background-color: #f0f0f0;
    padding: 5px 10px;
    margin-right: 5px;
    border-radius: 4px;
}

.section-title {
    font-size: 18px;
    color: #0066cc;
    margin-top: 20px;
    margin-bottom: 10px;
}

.question-item {
    margin-bottom: 20px;
}

.question {
    font-weight: bold;
    margin-bottom: 5px;
}

.answer,
.my-answer,
.description {
    margin-bottom: 8px;
    line-height: 1.6;
}

.my-answer {
    color: #0066cc;
}

.correct-answer {
    margin-top: 10px;
    padding-left: 10px;
    border-left: 3px solid #0099ff;
}

.correct-answer-title {
    font-weight: bold;
    margin-bottom: 5px;
}

.sub-section-title {
    font-size: 16px;
    color: #0066cc;
    margin-top: 15px;
    margin-bottom: 10px;
}
</style>
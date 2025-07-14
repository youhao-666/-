<script lang="ts" setup>
import { reactive, ref } from 'vue'
import { message } from 'ant-design-vue'
import type { FormInstance, FormRules } from 'element-plus'
import request from '@/utils/request';
import MarkdownIt from 'markdown-it';
const isMd = ref(true)
const options = ref([])
request({
  url: '/tag/list',
  method: 'GET',
}).then((res) => {
  console.log(res)
  if (res.status === 200) {
    message.success('获取标签成功')
    console.log(res.data.data)
    options.value = res.data.data.map((item: any) => {
      return {
        value: item.tagName,
        label: item.tagName,
      }
    })
  } else {
    message.success('获取标签失败')
  }
})



const md = new MarkdownIt();

const preview = ref('');

const updatePreview = () => {
  preview.value = md.render(ruleForm.content);
};

const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive({
  title: '',
  tag: '',
  content: '',
  location: 'Markdown编辑器',
  resource: '0',
})

const locationOptions = [
  { label: 'Markdown编辑器', value: 'Markdown编辑器' },
  { label: '富文本编辑器', value: '富文本编辑器' },
  // 其他选项
];

const rules = reactive<FormRules>({
  title: [
    { required: true, message: '请输入文章标题', trigger: 'blur' },
    { max: 20, message: '长度必须在20个字符以内', trigger: 'blur' },
  ],
  tag: [
    {
      required: true,
      message: '请选择文章标签',
      trigger: 'change',
    },
  ],
})

const submitForm = () => {
  console.log("submit", ruleForm)
  const json = JSON.parse(JSON.stringify({
    title: ruleForm.title,
    content: ruleForm.content,
    tags: ruleForm.tag,
    userId: Number(localStorage.getItem('id')),
  }
  ))
  const formData = new FormData()
  formData.append('data', JSON.stringify(json))
  request({
    url: `/article/create`,
    method: 'POST',
    data: formData,
    headers: {
      'Authorization': `Bearer ${localStorage.getItem('token')}`
    }
  }).then((res) => {
    console.log(res)
    if (res.status === 200) {
      // window.location.reload()
      message.success('发布成功')
      
    } else {
      message.error('发布失败')
    }
  })


}
const handleLocationChange = (val: string) => {
  isMd.value = val === 'Markdown编辑器'
  ruleForm.location = val
  console.log(val)
}

</script>


<template>
  <el-segmented v-model="ruleForm.location" default-value="Markdown编辑器" :options="locationOptions"
    @change="handleLocationChange" :style="{ position: 'absolute', top: '130px', right: '150px' }" />
  <el-form ref="ruleFormRef" style="max-width: 600px" :model="ruleForm" :rules="rules" label-width="auto">
    <el-form-item label="文章标题" prop="title">
      <el-input v-model="ruleForm.title" placeholder="请输入文章标题" />
    </el-form-item>
    <el-form-item label="文章标签" prop="tag">
      <el-select-v2 v-model="ruleForm.tag" placeholder="请选择文章标签" :options="options" multiple="true" />
    </el-form-item>
    <el-form-item label="文章内容" prop="content">
      <div class="markdown-editor" :style="{ display: isMd? 'block' : 'none' }">
        <textarea v-model="ruleForm.content" class="editor" @input="updatePreview"></textarea>
        <div class="preview" v-html="preview"></div>
      </div>
      <div :style="{ display: isMd? 'none' : 'block' }">
        <textarea v-model="ruleForm.content" class="editor" ></textarea>
      </div>
    </el-form-item>
    <el-form-item label="文章状态" prop="resource">
      <el-radio-group v-model="ruleForm.resource">
        <el-radio value="0">发布</el-radio>
        <el-radio value="1">保存为草稿</el-radio>
      </el-radio-group>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="submitForm()">
        提交
      </el-button>
    </el-form-item>
  </el-form>
</template>

<style>
.editor {
  width: 650px;
  height: 400px;
  border: 1px solid #ccc;
  padding: 10px;
  font-size: large;
  word-wrap: break-word;
  word-break: break-all;
  overflow: auto;
}

.preview {
  position: absolute;
  top: 0px;
  left: 700px;
  width: 650px;
  height: 400px;
  border: 1px solid #ccc;
  padding: 10px;
  word-wrap: break-word;
  word-break: break-all;
  overflow: auto;
}
</style>
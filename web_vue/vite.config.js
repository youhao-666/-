import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs[表情]ugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('.[表情]c', import.meta.url))
    },
  },
   server: {
    port: 8000, // 设置你想要的端口号
    host: '0.0.0.0', // 可选：允许局域网访问
    open: true,      // 可选：自动打开浏览器
  }
})
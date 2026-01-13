<template>
  <div id="app">
    <el-container style="height: 100vh">
      <el-aside width="auto">
        <Sidebar ref="sidebarRef" />
      </el-aside>
      <el-container>
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Sidebar from '@/components/Sidebar.vue'

const sidebarRef = ref<InstanceType<typeof Sidebar>>()
</script>

<style scoped>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  height: 100vh;
  margin: 0;
  padding: 0;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

.el-aside {
  overflow: hidden;
}

.el-main {
  padding: 0;
  overflow-y: auto;
}
</style>

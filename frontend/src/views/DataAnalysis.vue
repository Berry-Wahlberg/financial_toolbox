<template>
  <div class="data-analysis">
    <el-container>
      <el-main>
        <el-card>
          <template #header>
            <div class="card-header">
              <span>Under Development</span>
            </div>
          </template>
          <el-result icon="warning" title="Page Under Development" sub-title="This feature is currently being developed">
            <template #extra>
              <el-button type="primary" @click="$router.push('/')">Back to Home</el-button>
            </template>
          </el-result>
        </el-card>
      </el-main>
    </el-container>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.data-analysis {
  height: 100%;
}

.card-header {
  font-weight: bold;
}
</style>

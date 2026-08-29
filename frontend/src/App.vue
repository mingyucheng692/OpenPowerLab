<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { EndpointState } from './types/api'
import { fetchHealth, fetchPing } from './api/system'
import EndpointCard from './components/EndpointCard.vue'

const backendBaseUrl = 'http://localhost:8080'

const healthState = ref<EndpointState>({
  url: `${backendBaseUrl}/api/health`,
  status: null,
  loading: false,
  error: null,
  data: null,
})

const pingState = ref<EndpointState>({
  url: `${backendBaseUrl}/api/ping`,
  status: null,
  loading: false,
  error: null,
  data: null,
})

async function handleFetchHealth() {
  healthState.value.loading = true
  healthState.value.error = null
  try {
    const { status, data } = await fetchHealth()
    healthState.value.status = status
    healthState.value.data = data
  } catch (err: any) {
    healthState.value.error = err.message || '请求失败'
    healthState.value.data = null
  } finally {
    healthState.value.loading = false
  }
}

async function handleFetchPing() {
  pingState.value.loading = true
  pingState.value.error = null
  try {
    const { status, data } = await fetchPing()
    pingState.value.status = status
    pingState.value.data = data
  } catch (err: any) {
    pingState.value.error = err.message || '请求失败'
    pingState.value.data = null
  } finally {
    pingState.value.loading = false
  }
}

function fetchAll() {
  handleFetchHealth()
  handleFetchPing()
}

onMounted(() => {
  fetchAll()
})
</script>

<template>
  <div class="container">
    <h1>前后端通信验证 (OpenPowerLab)</h1>
    <p class="subtitle">前端: <code>http://localhost:5173</code> | 后端: <code>http://localhost:8080</code></p>

    <div class="actions">
      <button @click="fetchAll">重新请求全部</button>
      <button @click="handleFetchHealth">GET /api/health</button>
      <button @click="handleFetchPing">GET /api/ping</button>
    </div>

    <!-- Health Section -->
    <EndpointCard :endpoint="healthState" method="GET" />

    <!-- Ping Section -->
    <EndpointCard :endpoint="pingState" method="GET" />
  </div>
</template>

<style scoped>
.container {
  max-width: 720px;
  margin: 0 auto;
}

h1 {
  font-size: 20px;
  margin-bottom: 8px;
  color: #111827;
}

.subtitle {
  font-size: 13px;
  color: #6b7280;
  margin-bottom: 20px;
}

code {
  background-color: #e5e7eb;
  padding: 2px 6px;
  border-radius: 4px;
  font-family: monospace;
}

.actions {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

button {
  padding: 8px 14px;
  font-size: 13px;
  cursor: pointer;
  background-color: #2563eb;
  color: #ffffff;
  border: 1px solid #1d4ed8;
  border-radius: 4px;
}

button:hover {
  background-color: #1d4ed8;
}
</style>

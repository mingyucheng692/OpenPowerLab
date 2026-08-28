<script setup lang="ts">
import { ref, onMounted } from 'vue'

interface ApiResponse {
  status?: string
  service?: string
  [key: string]: any
}

interface EndpointState {
  url: string
  status: number | null
  loading: boolean
  error: string | null
  data: ApiResponse | null
}

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

async function fetchEndpoint(state: EndpointState) {
  state.loading = true
  state.error = null
  try {
    const res = await fetch(state.url)
    state.status = res.status
    if (!res.ok) {
      throw new Error(`HTTP error! status: ${res.status}`)
    }
    const json = await res.json()
    state.data = json
  } catch (err: any) {
    state.error = err.message || '请求失败'
    state.data = null
  } finally {
    state.loading = false
  }
}

function fetchAll() {
  fetchEndpoint(healthState.value)
  fetchEndpoint(pingState.value)
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
      <button @click="fetchEndpoint(healthState)">GET /api/health</button>
      <button @click="fetchEndpoint(pingState)">GET /api/ping</button>
    </div>

    <!-- Health Section -->
    <div class="card">
      <div class="card-header">
        <span class="method">GET</span>
        <span class="endpoint">{{ healthState.url }}</span>
        <span v-if="healthState.loading" class="tag loading">请求中...</span>
        <span v-else-if="healthState.status === 200" class="tag success">200 OK</span>
        <span v-else-if="healthState.error" class="tag error">Error</span>
      </div>
      <div v-if="healthState.error" class="error-msg">
        {{ healthState.error }}
      </div>
      <pre v-if="healthState.data">{{ JSON.stringify(healthState.data, null, 2) }}</pre>
    </div>

    <!-- Ping Section -->
    <div class="card">
      <div class="card-header">
        <span class="method">GET</span>
        <span class="endpoint">{{ pingState.url }}</span>
        <span v-if="pingState.loading" class="tag loading">请求中...</span>
        <span v-else-if="pingState.status === 200" class="tag success">200 OK</span>
        <span v-else-if="pingState.error" class="tag error">Error</span>
      </div>
      <div v-if="pingState.error" class="error-msg">
        {{ pingState.error }}
      </div>
      <pre v-if="pingState.data">{{ JSON.stringify(pingState.data, null, 2) }}</pre>
    </div>
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

.card {
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  background-color: #ffffff;
  padding: 14px;
  margin-bottom: 16px;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 10px;
  margin-bottom: 10px;
}

.method {
  font-size: 12px;
  font-weight: bold;
  background-color: #e0f2fe;
  color: #0369a1;
  padding: 2px 6px;
  border-radius: 4px;
}

.endpoint {
  font-size: 14px;
  font-family: monospace;
  color: #374151;
  flex: 1;
}

.tag {
  font-size: 12px;
  padding: 2px 8px;
  border-radius: 4px;
}

.tag.success {
  background-color: #dcfce7;
  color: #15803d;
}

.tag.loading {
  background-color: #fef9c3;
  color: #854d0e;
}

.tag.error {
  background-color: #fee2e2;
  color: #b91c1c;
}

.error-msg {
  color: #dc2626;
  font-size: 13px;
  margin-bottom: 8px;
}

pre {
  background-color: #f3f4f6;
  border: 1px solid #e5e7eb;
  border-radius: 4px;
  padding: 12px;
  font-size: 13px;
  font-family: monospace;
  overflow-x: auto;
  color: #1f2937;
}
</style>

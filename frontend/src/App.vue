<!-- SPDX-FileCopyrightText: 2026 mingyucheng692 -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

<script setup lang="ts">
import { onMounted } from 'vue'
import { healthApi, pingApi } from './api/system'
import { useEndpoint } from './composables/useEndpoint'
import EndpointCard from './components/EndpointCard.vue'

const healthEndpoint = useEndpoint(healthApi.url, healthApi.request)
const pingEndpoint = useEndpoint(pingApi.url, pingApi.request)

function fetchAll() {
  healthEndpoint.execute()
  pingEndpoint.execute()
}

onMounted(fetchAll)
</script>

<template>
  <div class="container">
    <h1>前后端通信验证 (OpenPowerLab)</h1>
    <p class="subtitle">前端: <code>/</code>（同源 + 开发代理） | 后端 API: <code>/api/*</code></p>

    <div class="actions">
      <button @click="fetchAll">重新请求全部</button>
      <button @click="healthEndpoint.execute">GET /api/health</button>
      <button @click="pingEndpoint.execute">GET /api/ping</button>
    </div>

    <!-- Health Section -->
    <EndpointCard :endpoint="healthEndpoint.state" method="GET" />

    <!-- Ping Section -->
    <EndpointCard :endpoint="pingEndpoint.state" method="GET" />
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

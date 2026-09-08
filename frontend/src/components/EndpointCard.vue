<!-- SPDX-FileCopyrightText: 2026 mingyucheng692 -->
<!-- SPDX-License-Identifier: Apache-2.0 -->

<script setup lang="ts">
import type { EndpointState } from '../types/api'

defineProps<{
  endpoint: EndpointState
  method?: string
}>()
</script>

<template>
  <div class="card">
    <div class="card-header">
      <span class="method">{{ method || 'GET' }}</span>
      <span class="endpoint">{{ endpoint.url }}</span>
      <span v-if="endpoint.loading" class="tag loading">请求中...</span>
      <span v-else-if="endpoint.status === 200" class="tag success">200 OK</span>
      <span v-else-if="endpoint.error" class="tag error">Error</span>
    </div>
    <div v-if="endpoint.error" class="error-msg">
      {{ endpoint.error }}
    </div>
    <pre v-if="endpoint.data">{{ JSON.stringify(endpoint.data, null, 2) }}</pre>
  </div>
</template>

<style scoped>
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

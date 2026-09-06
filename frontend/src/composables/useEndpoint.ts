import { reactive } from 'vue'
import type { EndpointState } from '../types/api'

export function useEndpoint<T>(url: string, fetchFn: () => Promise<T>) {
  const state = reactive({
    url,
    status: null,
    loading: false,
    error: null,
    data: null,
  }) as EndpointState<T>

  async function execute() {
    state.loading = true
    state.error = null
    try {
      const data = await fetchFn()
      state.status = 200
      state.data = data
    } catch (err) {
      state.status = null
      state.error = err instanceof Error ? err.message : '请求失败'
      state.data = null
    } finally {
      state.loading = false
    }
  }

  return {
    state,
    execute,
  }
}

import { ref, type Ref } from 'vue'
import type { EndpointState } from '../types/api'

export function useEndpoint<T>(url: string, fetchFn: () => Promise<T>) {
  const state: Ref<EndpointState<T>> = ref({
    url,
    status: null,
    loading: false,
    error: null,
    data: null,
  })

  async function execute() {
    state.value.loading = true
    state.value.error = null
    try {
      const data = await fetchFn()
      state.value.status = 200
      state.value.data = data as any
    } catch (err: any) {
      state.value.status = null
      state.value.error = err.message || '请求失败'
      state.value.data = null
    } finally {
      state.value.loading = false
    }
  }

  return {
    state,
    execute,
  }
}

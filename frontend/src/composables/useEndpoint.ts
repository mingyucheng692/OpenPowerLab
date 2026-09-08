// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

import { reactive } from 'vue'
import type { EndpointState } from '../types/api'

/**
 * Manages the request lifecycle of a single API endpoint.
 *
 * Returns a reactive {@link EndpointState} and an `execute` function that
 * triggers the request: while running, `loading` is true and prior errors are
 * cleared; on completion, either `data` holds the parsed response or `error`
 * holds the failure message. `status` is 200 on success, null otherwise.
 */
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

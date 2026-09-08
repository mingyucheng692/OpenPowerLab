// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

import { fetchJson, getFullUrl } from './client'
import type { HealthResponse, PingResponse } from '../types/api'

export const healthApi = {
  url: getFullUrl('/api/health'),
  request: () => fetchJson<HealthResponse>('/api/health'),
}

export const pingApi = {
  url: getFullUrl('/api/ping'),
  request: () => fetchJson<PingResponse>('/api/ping'),
}

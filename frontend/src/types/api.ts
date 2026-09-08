// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

export interface HealthResponse {
  status: string
  service: string
}

export interface PingResponse {
  message: string
}

export interface EndpointState<T = unknown> {
  url: string
  status: number | null
  loading: boolean
  error: string | null
  data: T | null
}

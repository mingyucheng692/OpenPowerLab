import type { ApiResponse } from '../types/api'

const backendBaseUrl = 'http://localhost:8080'

export async function fetchHealth(): Promise<{ status: number; data: ApiResponse }> {
  const res = await fetch(`${backendBaseUrl}/api/health`)
  if (!res.ok) {
    throw new Error(`HTTP error! status: ${res.status}`)
  }
  const data = await res.json()
  return { status: res.status, data }
}

export async function fetchPing(): Promise<{ status: number; data: ApiResponse }> {
  const res = await fetch(`${backendBaseUrl}/api/ping`)
  if (!res.ok) {
    throw new Error(`HTTP error! status: ${res.status}`)
  }
  const data = await res.json()
  return { status: res.status, data }
}

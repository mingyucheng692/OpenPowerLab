export const BASE_URL = 'http://localhost:8080'

export function getFullUrl(path: string): string {
  return `${BASE_URL}${path}`
}

export async function fetchJson<T>(path: string): Promise<T> {
  const res = await fetch(getFullUrl(path))
  if (!res.ok) {
    throw new Error(`HTTP ${res.status}: ${res.statusText}`)
  }
  return res.json()
}

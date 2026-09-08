// SPDX-FileCopyrightText: 2026 mingyucheng692
// SPDX-License-Identifier: Apache-2.0

// All browser requests use relative paths (same-origin or via dev proxy).
// Do not hardcode localhost / protocol / port here or anywhere else.

/**
 * Fetches a JSON resource and resolves with the parsed body.
 *
 * @throws Error when the response status is not 2xx, with a
 *         `HTTP <status>: <statusText>` message.
 */
export function fetchJson<T>(path: string, init?: RequestInit): Promise<T> {
  return doFetch<T>(path, init)
}

export function getFullUrl(path: string): string {
  return path
}

async function doFetch<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(path, init)
  if (!res.ok) {
    throw new Error(`HTTP ${res.status}: ${res.statusText}`)
  }
  return res.json()
}

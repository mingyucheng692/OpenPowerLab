export interface HealthResponse {
  status: string
  service: string
}

export interface PingResponse {
  message: string
}

export interface EndpointState<T = any> {
  url: string
  status: number | null
  loading: boolean
  error: string | null
  data: T | null
}

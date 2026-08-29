export interface ApiResponse {
  status?: string
  service?: string
  message?: string
  [key: string]: any
}

export interface EndpointState {
  url: string
  status: number | null
  loading: boolean
  error: string | null
  data: ApiResponse | null
}

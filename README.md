# OpenPowerLab

OpenPowerLab is an open-source platform for simulating power devices and experimenting with industrial communication protocols — Modbus, IEC 60870-5-104, and IEC 61850 — for protocol development, testing, and learning in power and energy applications.

## Status

> **Early stage.** The project currently validates the frontend–backend skeleton only (health/ping endpoints and a communication verification page). Protocol simulation is on the roadmap. Do not use in production.

## Quick Start

**Prerequisites:** Go 1.25+ and Node.js 20+.

1. Start the backend (listens on `:8080` by default; override with the `PORT` environment variable):

   ```bash
   cd backend
   go run ./cmd/server
   ```

2. Start the frontend dev server (listens on `:5173`, proxies `/api` to the backend):

   ```bash
   cd frontend
   npm install
   npm run dev
   ```

3. Open `http://localhost:5173` in your browser and use the verification page to call `/api/health` and `/api/ping`.

Stop either process with `Ctrl+C` (the backend shuts down gracefully).

## Project Structure

```
OpenPowerLab/
├── backend/       # Go HTTP backend (standard library only)
│   ├── cmd/       # Entry points
│   └── internal/  # Handlers, middleware, server
├── frontend/      # Vue 3 + TypeScript + Vite SPA
│   └── src/       # API layer, composables, components, types
└── docs/          # Public documentation
```

## Documentation

- [Architecture](docs/architecture.md) — system overview, directory layout, running topology, API reference, and roadmap.

## License

See [LICENSE](LICENSE).

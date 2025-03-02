# Video Streaming Platform

A modern video streaming application built with a polyglot architecture using Vue.js for the frontend and multiple backend options (Go and Rust).

## 🚀 Technology Stack

### 📦 Monorepo Structure
- **Turborepo**: High-performance build system for JavaScript/TypeScript monorepos
- **pnpm**: Fast, disk space efficient package manager
- **Workspaces**: Organized as multiple packages within a single repository

### 🖥️ Frontend
- **Vue.js 3**: Progressive JavaScript framework with Composition API
- **TypeScript**: Static typing for improved developer experience
- **Vite**: Next generation frontend tooling for faster development and optimized builds
- **Pinia**: Intuitive, type-safe store for Vue
- **Vue Router**: Official router for Vue.js

### 🔧 Backend Options

#### Go Backend
- **Go**: High-performance, statically typed, compiled language
- **Standard Go modules**: For dependency management

#### Rust Backend
- **Rust**: Systems programming language focused on safety, speed, and concurrency
- **Actix-web**: Powerful, pragmatic web framework for Rust
- **Cargo**: Rust's package manager and build system

### 🛠️ Development Tools
- **ESLint**: Linting utility for JavaScript and TypeScript
- **Prettier**: Code formatter
- **Vitest**: Unit testing framework for Vite
- **Vue Test Utils**: Official testing utilities for Vue.js

## 🏗️ Project Structure

```
video-stream/
├── app/
│   ├── frontend/         # Vue.js application
│   │   ├── src/          # Source code
│   │   ├── public/       # Static assets
│   │   └── ...
│   └── backend/
│       ├── go-backend/   # Go backend implementation
│       └── rust-backend/ # Rust backend implementation
├── turbo.json            # Turborepo configuration
├── package.json          # Root package.json with scripts
└── pnpm-workspace.yaml   # pnpm workspace configuration
```

## 🚀 Getting Started

### Prerequisites
- Node.js (v18+)
- pnpm (v10.5.2+)
- Go (v1.18+)
- Rust (v1.65+)

### Installation

```bash
# Clone the repository
git clone <repository-url>
cd video-stream

# Install dependencies
pnpm install
```

### Development

```bash
# Start all services in development mode
pnpm dev

# Start only frontend
pnpm --filter frontend dev

# Start only Go backend
pnpm --filter go-backend dev

# Start only Rust backend
pnpm --filter rust-backend dev
```

### Building

```bash
# Build all packages
pnpm build

# Build specific packages
pnpm --filter frontend build
pnpm go:build
pnpm rust:build
```

## 🧪 Testing

```bash
# Run all tests
pnpm test

# Run frontend tests
pnpm --filter frontend test
```

## 🧹 Linting and Formatting

```bash
# Run linting
pnpm lint

# Format code
pnpm --filter frontend format
```

## 🔄 CI/CD Integration

The monorepo structure with Turborepo enables efficient CI/CD pipelines:
- Incremental builds: Only rebuild what changed
- Caching: Reuse build artifacts across CI runs
- Parallel execution: Run tasks in parallel when possible

## 📚 Learn More

- [Vue.js Documentation](https://vuejs.org/)
- [Go Documentation](https://golang.org/doc/)
- [Rust Documentation](https://www.rust-lang.org/learn)
- [Turborepo Documentation](https://turbo.build/repo/docs)
- [pnpm Documentation](https://pnpm.io/motivation)

# CodeSmoker Test: Go API + React Frontend (#23)

A test repository for the CodeSmoker test suite demonstrating a mixed-language project with Go backend API and React frontend.

## Project Structure

```
├── backend/
│   ├── main.go           # Go Gin API server
│   └── go.mod            # Go module definition
├── frontend/
│   ├── src/
│   │   ├── App.jsx       # Main React component
│   │   ├── App.css       # Styles
│   │   ├── main.jsx      # Entry point
│   │   └── index.css     # Global styles
│   ├── index.html        # HTML template
│   ├── package.json      # Node dependencies
│   └── vite.config.js    # Vite configuration
└── .gitignore
```

## Features

### Backend (Go + Gin)
- **Gin Framework**: High-performance HTTP web framework
- **RESTful API**: Full CRUD operations
- **CORS Support**: Cross-origin requests enabled
- **Concurrent Safe**: Mutex-protected data access

### Frontend (React + Vite)
- **React 18**: Modern React with hooks
- **Vite**: Fast development server with HMR
- **Fetch API**: HTTP client for API calls
- **Responsive UI**: Clean, modern styling

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/todos` | Get all todos |
| GET | `/api/todos/:id` | Get todo by ID |
| POST | `/api/todos` | Create a new todo |
| PUT | `/api/todos/:id` | Update a todo |
| DELETE | `/api/todos/:id` | Delete a todo |

## Getting Started

### Prerequisites

- Go >= 1.21
- Node.js >= 18
- npm or pnpm

### Backend Setup

```bash
cd backend

# Install dependencies
go mod tidy

# Run server (port 8080)
go run main.go
```

### Frontend Setup

```bash
cd frontend

# Install dependencies
npm install

# Run dev server (port 5173)
npm run dev
```

### Access the App

- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

## Development

The frontend Vite config includes a proxy to forward `/api` requests to the Go backend:

```javascript
// vite.config.js
export default defineConfig({
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
})
```

## Documentation

Built using latest documentation from:
- [Gin Web Framework](https://gin-gonic.com) - Go Gin documentation
- [React](https://react.dev) - React documentation
- [Vite](https://vitejs.dev) - Vite documentation

---

*This is a CodeSmoker test repository*

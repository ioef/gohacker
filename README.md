# 🎮 GoHacker - Learn Go by Hacking!

A gamified platform for learning Go programming through interactive challenges. Built with Go (Gin) backend and modern web frontend.

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![Vue Version](https://img.shields.io/badge/Vue-3.4+-4FC08D?style=flat&logo=vue.js)
![License](https://img.shields.io/badge/License-MIT-green.svg)

## ✨ Features

- 🎯 **18+ Progressive Challenges** - From basics to concurrency
- 🔐 **JWT Authentication** - Secure user accounts
- 💾 **In-Memory Database** - Fast with JSON persistence (go-memdb)
- 🏆 **Achievement System** - Unlock badges and earn XP
- 📊 **Progress Tracking** - Monitor your learning journey
- 🏅 **Leaderboard** - Compete with other learners
- ⚡ **Real-time Code Execution** - Instant feedback on your code
- 🎨 **Cyberpunk Theme** - Cool hacker aesthetic
- 🧠 **Time-Friendly** - Short challenges, instant feedback, visual progress

## 🚀 Quick Start

### Prerequisites

- [Go 1.21+](https://golang.org/dl/)
- [Node.js 20+](https://nodejs.org/)
- [Task](https://taskfile.dev/) (optional but recommended)

### Using Task (Recommended)

```bash
# Install dependencies
task install

# Run backend (Terminal 1)
task backend:run

# Run frontend (Terminal 2)
task frontend:dev

# Visit http://localhost:3000
```

### Manual Setup

#### Backend
```bash
cd backend
go mod download
go run main.go
```

#### Frontend
```bash
cd frontend
npm install
npm run dev
```

## 📋 Available Tasks

```bash
task --list                 # Show all available tasks

# Build
task build                  # Build both backend and frontend
task backend:build          # Build backend binary
task frontend:build         # Build frontend for production

# Test
task test                   # Run all tests
task backend:test           # Run backend unit tests

# Lint
task lint                   # Lint all code
task backend:lint           # Lint backend code
task frontend:lint          # Lint frontend code

# Run
task backend:run            # Run backend server
task frontend:dev           # Run frontend dev server

# Clean
task clean                  # Clean all build artifacts
task backend:clean          # Clean backend artifacts
task frontend:clean         # Clean frontend artifacts

# Docker
task docker:build           # Build Docker image
task docker:run             # Run Docker container

# Kubernetes
task k8s:deploy             # Deploy to Kubernetes
task k8s:status             # Check deployment status
task k8s:logs               # View logs
task k8s:delete             # Delete deployment
```

## 🐳 Docker

### Build and Run

```bash
# Build image
task docker:build
# or
docker build -t gohacker:latest .

# Run container
task docker:run
# or
docker run -p 8080:8080 gohacker:latest
```

### Docker Compose (Coming Soon)

```bash
docker-compose up
```

## ☸️ Kubernetes Deployment

### Deploy to Kubernetes

```bash
# Deploy all resources
task k8s:deploy
# or
kubectl apply -f k8s/

# Check status
task k8s:status

# View logs
task k8s:logs

# Delete deployment
task k8s:delete
```

### Kubernetes Resources

- **Namespace**: `gohacker`
- **Deployment**: 2 replicas with health checks
- **Service**: LoadBalancer on port 80
- **PVC**: 1Gi persistent storage
- **Ingress**: NGINX with TLS support
- **ConfigMap**: Environment configuration

## 🏗️ Project Structure

```
gohacker/
├── backend/                 # Go backend
│   ├── api/                # API handlers, middleware, routes
│   ├── config/             # Configuration
│   ├── game/               # Game engine & challenges
│   ├── models/             # Data models
│   ├── storage/            # Database (go-memdb + persistence)
│   ├── utils/              # Utilities (JWT, sandbox)
│   ├── main.go             # Entry point
│   └── *_test.go           # Unit tests
├── frontend/               # Vue.js frontend
│   ├── src/
│   │   ├── components/    # Vue components
│   │   ├── views/         # Page views
│   │   ├── stores/        # Pinia stores
│   │   ├── services/      # API services
│   │   └── router/        # Vue Router
│   └── package.json
├── k8s/                    # Kubernetes manifests
│   ├── namespace.yaml
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── pvc.yaml
│   ├── ingress.yaml
│   └── configmap.yaml
├── Dockerfile              # Multi-stage Docker build
├── Taskfile.yml            # Task automation
└── README.md
```

## 🎯 Challenge Categories

### 🟢 Newbie Hacker (Level 1-10)
- Variables & Data Types
- Control Flow (if/else, loops)
- Functions
- Arrays & Slices
- Maps
- String Manipulation

### 🟡 Script Kiddie (Level 11-20)
- Structs & Methods
- Pointers
- Interfaces
- Error Handling
- File I/O

### 🟠 Elite Coder (Level 21-35)
- Goroutines
- Channels
- WaitGroups
- Mutexes
- Context

### 🔴 Master Hacker (Level 36-50)
- HTTP Servers
- REST APIs
- Testing
- Real-world Projects

## 🧪 Testing

```bash
# Run all tests
task test

# Run with coverage
cd backend && go test -v -cover ./...

# Run specific package tests
cd backend && go test -v ./models
cd backend && go test -v ./utils
```

## 📊 API Endpoints

### Public
- `GET /api/health` - Health check
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login

### Protected (JWT Required)
- `GET /api/auth/me` - Get current user
- `GET /api/challenges` - Get all challenges
- `GET /api/challenges/:id` - Get specific challenge
- `POST /api/challenges/:id/execute` - Execute code
- `GET /api/progress` - Get user progress
- `GET /api/progress/achievements` - Get achievements
- `GET /api/leaderboard` - Get top users

## 🔧 Configuration

Backend configuration in `backend/config/config.go`:

- **Server Port**: 8080 (default)
- **JWT Secret**: Change in production!
- **Execution Timeout**: 5 seconds
- **Auto-save Interval**: 60 seconds
- **Max Code Length**: 10,000 characters

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🎓 Learning Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Tour](https://tour.golang.org/)

## 🐛 Bug Reports

Found a bug? Please open an issue with:
- Description of the bug
- Steps to reproduce
- Expected behavior
- Screenshots (if applicable)

## 💡 Feature Requests

Have an idea? Open an issue with the `enhancement` label!

## 🙏 Acknowledgments

- Inspired by [Bitburner](https://github.com/danielyxie/bitburner)
- Built with [Gin](https://github.com/gin-gonic/gin)
- Frontend powered by [Vue.js](https://vuejs.org/)
- Database: [go-memdb](https://github.com/hashicorp/go-memdb)

## 📧 Contact

Created with ❤️ for busy developers who want to master Go in bite-sized chunks!

---

**Happy Hacking! 🚀💻✨**

# 🚀 GoHacker Quick Start Guide

Get up and running in 5 minutes!

## 📦 Prerequisites

Install these first:
- [Go 1.21+](https://golang.org/dl/)
- [Node.js 20+](https://nodejs.org/)
- [Task](https://taskfile.dev/#/installation) (optional but recommended)

## ⚡ Fastest Way to Start

### Option 1: Using Task (Recommended)

```bash
cd /home/jeff/gohacker

# Install all dependencies
task install

# Terminal 1: Start backend
task backend:run

# Terminal 2: Start frontend (in a new terminal)
task frontend:dev

# Open browser to http://localhost:3000
```

### Option 2: Manual Start

```bash
cd /home/jeff/gohacker

# Terminal 1: Backend
cd backend
go run main.go

# Terminal 2: Frontend (in a new terminal)
cd frontend
npm install
npm run dev

# Open browser to http://localhost:3000
```

## 🎮 First Steps

1. **Register an Account**
   - Open http://localhost:3000
   - Click "Register"
   - Create your hacker username!

2. **Start Your First Challenge**
   - Click on "Challenge 001: First Breach"
   - Read the mission brief
   - Write your Go code
   - Click "Run Code"
   - Watch your XP grow! 🎯

3. **Track Your Progress**
   - View your level and rank
   - Check unlocked achievements
   - Compete on the leaderboard

## 🧪 Run Tests

```bash
# Run all backend tests
task test

# Or manually
cd backend
go test -v ./...
```

## 🐳 Docker Quick Start

```bash
# Build and run with Docker
task docker:build
task docker:run

# Or manually
docker build -t gohacker:latest .
docker run -p 8080:8080 gohacker:latest
```

## ☸️ Kubernetes Quick Deploy

```bash
# Deploy to your cluster
task k8s:deploy

# Check status
task k8s:status

# View logs
task k8s:logs
```

## 📋 Common Tasks

```bash
task --list              # Show all available tasks
task build               # Build everything
task clean               # Clean build artifacts
task lint                # Lint code
task backend:test        # Run tests
```

## 🔧 Configuration

Backend runs on: `http://localhost:8080`
Frontend runs on: `http://localhost:3000`

To change ports, edit:
- Backend: `backend/config/config.go`
- Frontend: `frontend/vite.config.js`

## 🆘 Troubleshooting

### Backend won't start
```bash
cd backend
go mod tidy
go run main.go
```

### Frontend won't start
```bash
cd frontend
rm -rf node_modules
npm install
npm run dev
```

### Tests failing
```bash
cd backend
go mod download
go test -v ./...
```

### Port already in use
```bash
# Kill process on port 8080
lsof -ti:8080 | xargs kill -9

# Kill process on port 3000
lsof -ti:3000 | xargs kill -9
```

## 🎯 What's Next?

- Complete all 18+ challenges
- Unlock all achievements
- Reach Master Hacker rank (Level 50)
- Top the leaderboard!

## 📚 Learn More

- [Full README](README.md)
- [Backend API Docs](backend/README.md)
- [Task Commands](Taskfile.yml)

---

**Happy Hacking! 🎮💻✨**

# GoHacker Backend API

A gamified Go learning platform backend built with Gin framework.

## Features

- 🔐 JWT Authentication
- 🎮 18+ Progressive Challenges
- 💾 In-memory database (go-memdb) with JSON persistence
- 🏆 Achievement System
- 📊 Progress Tracking
- 🏅 Leaderboard
- ⚡ Real-time Code Execution

## API Endpoints

### Public
- `GET /api/health` - Health check
- `POST /api/auth/register` - Register new user
- `POST /api/auth/login` - Login

### Protected (requires JWT token)
- `GET /api/auth/me` - Get current user
- `GET /api/challenges` - Get all challenges
- `GET /api/challenges/:id` - Get specific challenge
- `POST /api/challenges/:id/execute` - Execute code for challenge
- `GET /api/progress` - Get user progress
- `GET /api/progress/achievements` - Get achievements
- `GET /api/leaderboard` - Get top users

## Running

```bash
cd backend
go run main.go
```

Server will start on `http://localhost:8080`

## Configuration

Edit `config/config.go` to change:
- Server port
- JWT secret
- Execution timeout
- Auto-save interval

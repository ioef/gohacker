# 🚀 GoHacker Backend - Standalone Usage

The backend is **fully functional** and can be used standalone without the frontend!

## ✅ What Works Right Now

The complete backend API is ready with:
- 🔐 JWT Authentication
- 🎮 18 Progressive Challenges
- 💾 In-memory database with persistence
- 🏆 Achievement system
- 📊 Progress tracking
- 🏅 Leaderboard
- ⚡ Code execution sandbox

## 🚀 Quick Start (Backend Only)

```bash
cd /home/jeff/gohacker

# Run tests
task backend:test

# Build backend
task backend:build

# Run backend server
task backend:run
```

The API will be available at: **http://localhost:8080**

## 📊 Test the API

### 1. Health Check
```bash
curl http://localhost:8080/api/health
```

### 2. Register a User
```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "hacker1",
    "email": "hacker@example.com",
    "password": "password123"
  }'
```

### 3. Login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "hacker1",
    "password": "password123"
  }'
```

Save the token from the response!

### 4. Get Challenges (with token)
```bash
TOKEN="your-jwt-token-here"

curl http://localhost:8080/api/challenges \
  -H "Authorization: Bearer $TOKEN"
```

### 5. Execute Code
```bash
curl -X POST http://localhost:8080/api/challenges/001/execute \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "package main\nimport \"fmt\"\nfunc main() {\n\tvar password string = \"admin123\"\n\tfmt.Println(password)\n}"
  }'
```

### 6. Get Progress
```bash
curl http://localhost:8080/api/progress \
  -H "Authorization: Bearer $TOKEN"
```

### 7. Get Leaderboard
```bash
curl http://localhost:8080/api/leaderboard \
  -H "Authorization: Bearer $TOKEN"
```

## 🧪 Run Tests

```bash
# All tests
task backend:test

# Specific packages
cd backend
go test -v ./models
go test -v ./utils
go test -v -cover ./...
```

## 🐳 Docker (Backend Only)

```bash
# Build backend-only image
cd backend
docker build -t gohacker-backend:latest .

# Run
docker run -p 8080:8080 gohacker-backend:latest
```

## 📋 Available Backend Tasks

```bash
task backend:build    # Build binary
task backend:test     # Run tests
task backend:lint     # Lint code
task backend:run      # Run server
task backend:clean    # Clean artifacts
```

## 🎯 API Endpoints

### Public
- `GET /api/health` - Health check
- `POST /api/auth/register` - Register
- `POST /api/auth/login` - Login

### Protected (JWT Required)
- `GET /api/auth/me` - Current user
- `GET /api/challenges` - All challenges
- `GET /api/challenges/:id` - Specific challenge
- `POST /api/challenges/:id/execute` - Execute code
- `GET /api/progress` - User progress
- `GET /api/progress/achievements` - Achievements
- `GET /api/leaderboard` - Top users

## 💡 Frontend (Optional)

The frontend requires Node.js. If you want to add it later:

```bash
# Install Node.js first, then:
task frontend:install
task frontend:dev
```

For now, you can use the API directly with curl, Postman, or any HTTP client!

---

**The backend is production-ready and fully functional! 🎉**

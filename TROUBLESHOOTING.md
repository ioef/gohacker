# 🔧 GoHacker Troubleshooting Guide

## Getting 403 Forbidden Error

If you're getting a 403 error when accessing https://localhost:3000, try these fixes:

### Fix 1: Fix File Permissions

```bash
cd /home/jeff/gohacker

# Fix frontend permissions
chmod -R 755 frontend/dist
chmod 644 frontend/dist/index.html

# Restart services
task down
task run
```

### Fix 2: Check if Services are Running

```bash
# Check Docker containers
docker ps

# You should see:
# - gohacker-backend
# - gohacker-nginx
```

### Fix 3: Check NGINX Logs

```bash
# View NGINX logs
docker logs gohacker-nginx

# View backend logs
docker logs gohacker-backend
```

### Fix 4: Rebuild Everything

```bash
cd /home/jeff/gohacker

# Clean everything
task docker:clean

# Rebuild from scratch
task docker:rebuild

# Or use the full rebuild
task down
docker-compose build --no-cache
task run
```

### Fix 5: Try Direct Backend Access

If NGINX isn't working, try accessing the backend directly:

```bash
# Test backend directly
curl http://localhost:8080/api/health
```

If this works, the issue is with NGINX configuration.

### Fix 6: Check Port Availability

```bash
# Check if ports are in use
sudo lsof -i :3000
sudo lsof -i :8080
sudo lsof -i :443

# Kill processes if needed
sudo lsof -ti:3000 | xargs kill -9
```

## Other Common Issues

### Docker Compose Not Found

```bash
# Install docker-compose
sudo apt-get update
sudo apt-get install docker-compose

# Or use Docker Compose V2
docker compose version
```

### Permission Denied Errors

```bash
# Add your user to docker group
sudo usermod -aG docker $USER

# Log out and log back in, then:
newgrp docker
```

### SSL Certificate Issues

```bash
# Regenerate SSL certificates
cd /home/jeff/gohacker
./scripts/generate-ssl.sh

# Restart
task down
task run
```

### Backend Won't Start

```bash
# Check Go version
go version  # Should be 1.21+

# Rebuild backend
cd backend
go mod tidy
go build -o gohacker main.go
./gohacker
```

### Frontend Files Missing

```bash
# Check if index.html exists
ls -la /home/jeff/gohacker/frontend/dist/

# If missing, it should have been created
# Check the file
cat /home/jeff/gohacker/frontend/dist/index.html
```

## Quick Diagnostic

Run this to check everything:

```bash
cd /home/jeff/gohacker

echo "=== Checking Files ==="
ls -la frontend/dist/

echo "=== Checking Docker ==="
docker ps

echo "=== Checking Ports ==="
sudo lsof -i :3000
sudo lsof -i :8080

echo "=== Testing Backend ==="
curl http://localhost:8080/api/health

echo "=== Checking Logs ==="
docker logs gohacker-nginx --tail 20
docker logs gohacker-backend --tail 20
```

## Still Not Working?

### Option 1: Use Backend Only

The backend works standalone without Docker:

```bash
cd /home/jeff/gohacker/backend
go run main.go

# Access at: http://localhost:8080/api/health
```

### Option 2: Simplify Docker Setup

Create a simpler docker-compose without NGINX:

```yaml
version: '3.8'
services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=release
```

Then access directly at: http://localhost:8080/api/health

## Get Help

If none of these work, provide:
1. Output of `docker ps`
2. Output of `docker logs gohacker-nginx`
3. Output of `docker logs gohacker-backend`
4. Your OS and Docker version

---

**Most Common Fix**: Run `chmod -R 755 frontend/dist` and restart!

# 🐳 GoHacker - Docker Quick Start with HTTPS

Run the entire application with one command!

## 🚀 Super Quick Start

```bash
cd /home/jeff/gohacker

# Start everything with HTTPS
task up

# Access the app
# Open: https://localhost:3000
```

That's it! The app is running with HTTPS! 🎉

## 📋 What `task up` Does

1. Generates self-signed SSL certificates
2. Builds the backend Docker image
3. Starts backend + nginx with docker-compose
4. Serves on HTTPS port 3000

## 🌐 Access Points

- **Frontend (HTTPS)**: https://localhost:3000
- **API (HTTPS)**: https://localhost:3000/api/health
- **HTTP** (redirects to HTTPS): http://localhost:80

⚠️ **Note**: You'll see a security warning because we use self-signed certificates. This is normal for development. Click "Advanced" → "Proceed to localhost"

## 📋 Available Commands

```bash
# Start services
task up                 # Start with HTTPS

# View logs
task logs               # Follow all logs

# Stop services
task down               # Stop all containers

# Restart
task restart            # Restart services

# Rebuild
task docker:rebuild     # Rebuild from scratch

# Clean up
task docker:clean       # Remove all containers & volumes
```

## 🧪 Test the API

Once running, test with curl:

```bash
# Health check (accept self-signed cert)
curl -k https://localhost:3000/api/health

# Register a user
curl -k -X POST https://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "hacker1",
    "email": "test@example.com",
    "password": "password123"
  }'

# Login
curl -k -X POST https://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "hacker1",
    "password": "password123"
  }'
```

## 🔧 Architecture

```
┌─────────────────────────────────────┐
│   Browser (https://localhost:3000)  │
└──────────────┬──────────────────────┘
               │ HTTPS
               ▼
┌─────────────────────────────────────┐
│   NGINX (SSL Termination)           │
│   - Serves frontend                 │
│   - Proxies /api/* to backend       │
└──────────────┬──────────────────────┘
               │ HTTP
               ▼
┌─────────────────────────────────────┐
│   Go Backend (Gin)                  │
│   - REST API                        │
│   - Code execution                  │
│   - JWT auth                        │
└─────────────────────────────────────┘
```

## 📊 Services

### Backend
- **Image**: Built from `backend/Dockerfile`
- **Port**: 8080 (internal)
- **Health**: Auto-checked every 30s
- **Data**: Persisted in Docker volume

### NGINX
- **Image**: nginx:alpine
- **Ports**: 80 (HTTP), 443 (HTTPS), 3000 (HTTPS)
- **SSL**: Self-signed certificates
- **Config**: `nginx/nginx.conf`

## 🔐 SSL Certificates

Self-signed certificates are auto-generated in `nginx/ssl/`:
- `cert.pem` - Certificate
- `key.pem` - Private key

Valid for 365 days. Regenerate with:
```bash
./scripts/generate-ssl.sh
```

## 🐛 Troubleshooting

### Port already in use
```bash
# Stop services
task down

# Or kill specific ports
sudo lsof -ti:3000 | xargs kill -9
sudo lsof -ti:443 | xargs kill -9
```

### View logs
```bash
# All services
task logs

# Specific service
docker-compose logs -f backend
docker-compose logs -f nginx
```

### Rebuild everything
```bash
task docker:rebuild
```

### Clean slate
```bash
task docker:clean
task up
```

## 🎯 What's Next?

1. Open https://localhost:3000
2. Accept the security warning
3. See the landing page
4. Test the API with the buttons
5. Use curl or Postman to interact with the API

## 📚 More Info

- [Full README](README.md)
- [Backend API Guide](BACKEND_ONLY.md)
- [Kubernetes Deployment](k8s/)

---

**Happy Hacking with Docker! 🐳🚀**

#!/bin/bash

# GoHacker Startup Script
# Starts all services with docker-compose and shows logs

set -e

echo "🎮 GoHacker - Starting All Services"
echo "===================================="
echo ""

# Check if docker-compose is installed
if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
    echo "❌ Error: docker-compose not found!"
    echo ""
    echo "Please install docker-compose:"
    echo "  sudo apt-get install docker-compose"
    echo "  or"
    echo "  sudo curl -L \"https://github.com/docker/compose/releases/latest/download/docker-compose-\$(uname -s)-\$(uname -m)\" -o /usr/local/bin/docker-compose"
    echo "  sudo chmod +x /usr/local/bin/docker-compose"
    exit 1
fi

# Generate SSL certificates if they don't exist
if [ ! -f "./nginx/ssl/cert.pem" ]; then
    echo "🔐 Generating SSL certificates..."
    ./scripts/generate-ssl.sh
    echo ""
fi

# Stop any existing containers
echo "🛑 Stopping existing containers..."
docker-compose down 2>/dev/null || docker compose down 2>/dev/null || true
echo ""

# Build and start services
echo "🔨 Building Docker images..."
docker-compose build || docker compose build
echo ""

echo "🚀 Starting services..."
docker-compose up -d || docker compose up -d
echo ""

# Wait for services to be healthy
echo "⏳ Waiting for services to be ready..."
sleep 5

# Check if services are running
if docker-compose ps | grep -q "Up" || docker compose ps | grep -q "Up"; then
    echo ""
    echo "✅ GoHacker is running!"
    echo ""
    echo "📍 Access Points:"
    echo "   🌐 Frontend (HTTPS): https://localhost:3000"
    echo "   🔧 API Health:       https://localhost:3000/api/health"
    echo "   📊 Backend Direct:   http://localhost:8080"
    echo ""
    echo "⚠️  Note: Accept the self-signed certificate warning in your browser"
    echo ""
    echo "📋 Useful Commands:"
    echo "   task logs    - View logs"
    echo "   task down    - Stop services"
    echo "   task restart - Restart services"
    echo ""
    echo "📜 Showing logs (Ctrl+C to exit, services keep running)..."
    echo "================================================================"
    echo ""
    
    # Follow logs
    docker-compose logs -f || docker compose logs -f
else
    echo ""
    echo "❌ Error: Services failed to start"
    echo ""
    echo "Check logs with: docker-compose logs"
    exit 1
fi

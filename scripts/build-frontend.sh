#!/bin/bash
set -e

echo "🎨 Building GoHacker Frontend..."

# Ensure dist directory exists
mkdir -p frontend/dist

# The index.html is our static frontend - it's already built!
# Just verify it exists
if [ -f "frontend/dist/index.html" ]; then
    echo "✅ Frontend ready: frontend/dist/index.html"
    echo "📊 Size: $(du -h frontend/dist/index.html | cut -f1)"
else
    echo "❌ Error: frontend/dist/index.html not found!"
    echo "💡 The static HTML file should be in git. Try: git checkout frontend/dist/index.html"
    exit 1
fi

echo "🎉 Frontend build complete!"

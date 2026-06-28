#!/bin/bash

# Generate self-signed SSL certificates for development

SSL_DIR="./nginx/ssl"
mkdir -p "$SSL_DIR"

echo "🔐 Generating self-signed SSL certificates..."

openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout "$SSL_DIR/key.pem" \
  -out "$SSL_DIR/cert.pem" \
  -subj "/C=US/ST=State/L=City/O=GoHacker/CN=localhost"

chmod 600 "$SSL_DIR/key.pem"
chmod 644 "$SSL_DIR/cert.pem"

echo "✅ SSL certificates generated in $SSL_DIR"
echo "   - Certificate: $SSL_DIR/cert.pem"
echo "   - Private Key: $SSL_DIR/key.pem"
echo ""
echo "⚠️  These are self-signed certificates for development only!"
echo "   Your browser will show a security warning - this is normal."

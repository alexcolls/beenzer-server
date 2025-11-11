# Beenzer Server v3 - Deployment Guide

This guide covers different deployment strategies for Beenzer Server v3.

## Table of Contents

- [Quick Start](#quick-start)
- [Local Development](#local-development)
- [Docker Deployment](#docker-deployment)
- [Production Deployment](#production-deployment)
- [Environment Variables](#environment-variables)
- [Monitoring](#monitoring)
- [Troubleshooting](#troubleshooting)

## Quick Start

### Prerequisites

- Go 1.21+ installed
- PostgreSQL 14+ running
- Valid `.env` file configured

### Run Locally

```bash
# Install dependencies
make install

# Run the server
make run

# Or use hot reload for development
make dev
```

The server will be available at `http://localhost:4000`

## Local Development

### Setup

1. **Install Go dependencies**:
   ```bash
   go mod download
   ```

2. **Set up environment variables**:
   ```bash
   cp .env.sample .env
   # Edit .env with your configuration
   ```

3. **Create PostgreSQL databases**:
   ```sql
   CREATE DATABASE beenzer_users;
   CREATE DATABASE beenzer_messages;
   CREATE DATABASE beenzer_nfts;
   CREATE DATABASE beenzer_tokens;
   ```

4. **Run the server**:
   ```bash
   go run cmd/server/main.go
   ```

### Development Commands

```bash
make help           # Show all available commands
make dev            # Run with hot reload (installs air if needed)
make test           # Run tests
make test-cover     # Run tests with coverage report
make fmt            # Format code
make lint           # Run linter
make swagger        # Generate Swagger docs
```

## Docker Deployment

### Using Docker Compose (Recommended)

Docker Compose will set up the complete stack including all 4 PostgreSQL databases:

```bash
# Build and start all services
docker-compose up -d

# View logs
docker-compose logs -f beenzer-server

# Stop all services
docker-compose down

# Stop and remove volumes (⚠️  deletes all data)
docker-compose down -v
```

### Using Docker Only

```bash
# Build the image
docker build -t beenzer-server:v3 .

# Run the container (assuming PostgreSQL is running on host)
docker run -d \
  --name beenzer-server \
  -p 4000:4000 \
  --env-file .env \
  beenzer-server:v3
```

### Docker Configuration

#### Environment Variables

Create a `.env` file for Docker Compose:

```env
# Database password
DB_PASSWORD=your_secure_password_here

# Solana Configuration
SOLANA_RPC_URL=https://api.mainnet-beta.solana.com
MASTER_WALLET=YourSolanaPublicKey
MASTER_WALLET_KEYPAIR=comma,separated,keypair,bytes

# Google Maps
GOOGLE_MAPS_API_KEY=your_api_key_here
```

#### Health Checks

Both the server and databases have health checks configured:
- **Server**: HTTP check on `/health` every 30s
- **Databases**: PostgreSQL ready check every 10s

## Production Deployment

### Binary Deployment

1. **Build the binary**:
   ```bash
   CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -o beenzer-server cmd/server/main.go
   ```

2. **Copy to production server**:
   ```bash
   scp beenzer-server user@server:/opt/beenzer/
   scp .env user@server:/opt/beenzer/
   ```

3. **Create systemd service** (`/etc/systemd/system/beenzer-server.service`):
   ```ini
   [Unit]
   Description=Beenzer Server v3
   After=network.target postgresql.service

   [Service]
   Type=simple
   User=beenzer
   WorkingDirectory=/opt/beenzer
   ExecStart=/opt/beenzer/beenzer-server
   Restart=always
   RestartSec=10

   # Security
   NoNewPrivileges=true
   PrivateTmp=true
   ProtectSystem=strict
   ProtectHome=true
   ReadWritePaths=/opt/beenzer

   [Install]
   WantedBy=multi-user.target
   ```

4. **Enable and start the service**:
   ```bash
   sudo systemctl daemon-reload
   sudo systemctl enable beenzer-server
   sudo systemctl start beenzer-server
   sudo systemctl status beenzer-server
   ```

### Nginx Reverse Proxy

Configure Nginx as a reverse proxy:

```nginx
upstream beenzer_backend {
    server localhost:4000;
}

server {
    listen 80;
    server_name api.beenzer.com;

    # Redirect HTTP to HTTPS
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name api.beenzer.com;

    # SSL Configuration
    ssl_certificate /etc/letsencrypt/live/api.beenzer.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/api.beenzer.com/privkey.pem;

    # Logging
    access_log /var/log/nginx/beenzer_access.log;
    error_log /var/log/nginx/beenzer_error.log;

    # WebSocket support
    location /ws {
        proxy_pass http://beenzer_backend;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

    # API endpoints
    location / {
        proxy_pass http://beenzer_backend;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;

        # CORS headers (if needed)
        add_header 'Access-Control-Allow-Origin' '*' always;
        add_header 'Access-Control-Allow-Methods' 'GET, POST, PUT, DELETE, OPTIONS' always;
        add_header 'Access-Control-Allow-Headers' 'DNT,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Range,Authorization' always;

        if ($request_method = 'OPTIONS') {
            add_header 'Access-Control-Allow-Origin' '*';
            add_header 'Access-Control-Allow-Methods' 'GET, POST, PUT, DELETE, OPTIONS';
            add_header 'Access-Control-Max-Age' 1728000;
            add_header 'Content-Type' 'text/plain; charset=utf-8';
            add_header 'Content-Length' 0;
            return 204;
        }
    }
}
```

### Kubernetes Deployment (Advanced)

Example Kubernetes manifests:

**deployment.yaml**:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: beenzer-server-v3
  labels:
    app: beenzer-server
    version: v3
spec:
  replicas: 3
  selector:
    matchLabels:
      app: beenzer-server
  template:
    metadata:
      labels:
        app: beenzer-server
        version: v3
    spec:
      containers:
      - name: beenzer-server
        image: beenzer-server:v3
        ports:
        - containerPort: 4000
        env:
        - name: GO_ENV
          value: "production"
        envFrom:
        - secretRef:
            name: beenzer-secrets
        livenessProbe:
          httpGet:
            path: /health
            port: 4000
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 4000
          initialDelaySeconds: 5
          periodSeconds: 5
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "512Mi"
            cpu: "1000m"
```

**service.yaml**:
```yaml
apiVersion: v1
kind: Service
metadata:
  name: beenzer-server-service
spec:
  selector:
    app: beenzer-server
  ports:
  - protocol: TCP
    port: 80
    targetPort: 4000
  type: LoadBalancer
```

## Environment Variables

### Required

- `PORT` - Server port (default: 4000)
- `USERS_DB_URL` - PostgreSQL connection string for users database
- `MESSAGES_DB_URL` - PostgreSQL connection string for messages database
- `NFTS_DB_URL` - PostgreSQL connection string for NFTs database
- `TOKEN_DB_URL` - PostgreSQL connection string for tokens database

### Optional

- `GO_ENV` - Environment (development/production)
- `SOLANA_RPC_URL` - Solana RPC endpoint
- `MASTER_WALLET` - Solana wallet public key
- `MASTER_WALLET_KEYPAIR` - Solana wallet secret key
- `GOOGLE_MAPS_API_KEY` - Google Maps API key

See `.env.sample` for complete list.

## Monitoring

### Health Check Endpoint

```bash
curl http://localhost:4000/health
```

### Logs

The server uses structured JSON logging in production:

```bash
# View logs
journalctl -u beenzer-server -f

# Or with Docker
docker logs -f beenzer-server-v3
```

### Metrics (Future)

Prometheus metrics endpoint (planned):
```
http://localhost:4000/metrics
```

## Troubleshooting

### Server Won't Start

1. **Check database connections**:
   ```bash
   # Test database connectivity
   psql "postgresql://username:password@localhost:5432/beenzer_users"
   ```

2. **Verify environment variables**:
   ```bash
   # Check if .env is loaded
   cat .env
   ```

3. **Check port availability**:
   ```bash
   # See if port 4000 is in use
   lsof -i :4000
   ```

### Database Connection Errors

- Verify database URLs in `.env`
- Ensure PostgreSQL is running
- Check firewall rules
- Verify database exists

### WebSocket Connection Issues

- Check CORS configuration
- Verify WebSocket upgrade is working
- Check reverse proxy configuration
- Review browser console for errors

### Performance Issues

1. **Check database connection pool**:
   - Adjust `MaxConns` in `database.go`
   - Monitor active connections

2. **Profile the application**:
   ```bash
   # Run with profiling
   go run -cpuprofile=cpu.prof cmd/server/main.go
   
   # Analyze profile
   go tool pprof cpu.prof
   ```

3. **Check system resources**:
   ```bash
   # Memory usage
   docker stats beenzer-server-v3
   
   # Or with systemd
   systemctl status beenzer-server
   ```

### Getting Help

- **GitHub Issues**: [Report bugs](https://github.com/beenzer/beenzer-server/issues)
- **Discussions**: [Ask questions](https://github.com/beenzer/beenzer-server/discussions)
- **Logs**: Always include logs when reporting issues

---

**Note**: Always backup your databases before major updates or migrations.

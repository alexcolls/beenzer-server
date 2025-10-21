# 🌍 Beenzer Server

> A geo-social backend platform with Solana blockchain integration for NFTs and SPL tokens

[![TypeScript](https://img.shields.io/badge/TypeScript-4.6.4-blue.svg)](https://www.typescriptlang.org/)
[![Node.js](https://img.shields.io/badge/Node.js-16-green.svg)](https://nodejs.org/)
[![Socket.io](https://img.shields.io/badge/Socket.io-4.5.4-black.svg)](https://socket.io/)
[![Solana](https://img.shields.io/badge/Solana-Web3.js-purple.svg)](https://solana.com/)
[![License](https://img.shields.io/badge/License-ISC-yellow.svg)](LICENSE)

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Environment Setup](#environment-setup)
- [Database Setup](#database-setup)
- [Development](#development)
- [API Documentation](#api-documentation)
- [Socket.io Events](#socketio-events)
- [Docker Deployment](#docker-deployment)
- [Contributing](#contributing)
- [License](#license)
- [Related Projects](#related-projects)

## 🎯 Overview

Beenzer Server is the backend infrastructure for a geo-social mobile application that combines location-based social features with Solana blockchain technology. It provides real-time communication, user management, NFT minting/trading, and SPL token operations.

The server enables users to:
- Connect with others based on geographic proximity
- Mint location-tagged NFTs with metadata stored on Arweave
- Trade SPL tokens and NFTs on the Solana blockchain
- Engage in real-time messaging and social interactions
- Track wallet balances and transaction history

## ✨ Features

### Core Functionality
- 🔐 **User Management** - Registration, authentication, profiles, and social connections
- 💬 **Real-time Messaging** - WebSocket-based instant messaging between users
- 📍 **Geolocation Services** - Google Maps API integration for location-based features
- 👥 **Social Network** - Friend connections, followers, and user discovery

### Blockchain Integration
- 🎨 **NFT Minting** - Create geo-tagged NFTs using Metaplex Foundation
- 🪙 **SPL Token Operations** - Mint, transfer, and manage Solana tokens
- 💰 **Wallet Management** - Balance tracking for SOL and USDC
- 🔗 **Transaction Handling** - Secure blockchain transaction processing
- 📦 **Arweave Storage** - Decentralized metadata and asset storage via Bundlr

### Technical Features
- ⚡ **Real-time Communication** - Socket.io for bidirectional event-based communication
- 🗄️ **PostgreSQL Database** - Robust data persistence with multiple database connections
- 🐳 **Docker Support** - Containerized deployment for consistent environments
- 📊 **Scheduled Tasks** - Automated balance checks and periodic updates
- 🔄 **Version Control** - Dual version architecture (v1 and v2) for iterative development

## 🏗️ Architecture

### Version Comparison

The project maintains two versions with distinct architectural approaches:

**v1 (Monolithic)**
- Consolidated service layer
- Unified socket handling
- Single controller structure

**v2 (Modular)**
- Separated socket handlers by domain
- Distinct query and schema files
- Enhanced separation of concerns
- Improved maintainability and scalability

### System Components

```text
┌─────────────────────────────────────────────────────────┐
│                    Client Applications                   │
│              (Mobile App, Web Dashboard)                 │
└────────────────────────┬────────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────────┐
│                  Beenzer Server (Express)                │
│  ┌────────────┐  ┌────────────┐  ┌────────────┐        │
│  │   REST     │  │  Socket.io │  │   Cron     │        │
│  │   API      │  │  WebSocket │  │   Jobs     │        │
│  └─────┬──────┘  └──────┬─────┘  └──────┬─────┘        │
│        │                │                │               │
│  ┌─────▼────────────────▼────────────────▼─────┐        │
│  │           Controllers & Services            │        │
│  │  (Users, Messages, NFTs, Tokens)            │        │
│  └─────┬────────────────┬────────────────┬─────┘        │
└────────┼────────────────┼────────────────┼──────────────┘
         │                │                │
         ▼                ▼                ▼
┌──────────────┐  ┌──────────────┐  ┌──────────────┐
│  PostgreSQL  │  │    Solana    │  │  Google Maps │
│   Database   │  │  Blockchain  │  │     API      │
└──────────────┘  └──────┬───────┘  └──────────────┘
                         │
                         ▼
                  ┌──────────────┐
                  │   Arweave/   │
                  │    Bundlr    │
                  └──────────────┘
```

## 🛠️ Tech Stack

### Backend Framework
- **Node.js** v16 - JavaScript runtime
- **TypeScript** 4.6.4 - Type-safe development
- **Express.js** 4.18.2 - Web application framework
- **Socket.io** 4.5.4 - Real-time bidirectional communication

### Database
- **PostgreSQL** - Primary data store
- **pg** 8.8.0 - PostgreSQL client for Node.js

### Blockchain
- **@solana/web3.js** 1.66.2 - Solana blockchain interaction
- **@solana/spl-token** 0.3.6 - SPL token program interface
- **@metaplex-foundation/js** 0.17.6 - NFT minting and metadata
- **Bundlr Network** - Decentralized storage for NFT assets

### External Services
- **@googlemaps/google-maps-services-js** 3.3.16 - Geocoding and location services
- **@swiftcomplete/reverse-geocode** 1.1.2 - Reverse geocoding

### Development Tools
- **ts-node** 10.9.1 - TypeScript execution
- **nodemon** - Hot reload during development
- **Jest** 29.3.1 - Testing framework
- **Docker** - Containerization

### Utilities
- **dotenv** 16.0.3 - Environment variable management
- **cors** 2.8.5 - Cross-origin resource sharing
- **lodash** 4.17.21 - Utility functions

## 📁 Project Structure

```text
beenzer-server/
├── v1/                          # Version 1 (legacy)
│   ├── src/
│   │   ├── controllers/         # Request handlers and database logic
│   │   │   ├── users.controller.ts
│   │   │   ├── messages.controller.ts
│   │   │   ├── nfts.controller.ts
│   │   │   ├── token.controller.ts
│   │   │   ├── db.connections.ts
│   │   │   ├── queries.ts
│   │   │   └── schemas.ts
│   │   ├── services/            # Business logic and external integrations
│   │   │   ├── mintNFT.ts       # NFT minting with Metaplex
│   │   │   ├── sendNFT.ts       # NFT transfer operations
│   │   │   ├── mintToken.ts     # SPL token minting
│   │   │   ├── sendToken.ts     # SPL token transfers
│   │   │   ├── getBalances.ts   # Wallet balance queries
│   │   │   ├── sockets.ts       # Socket.io event handlers
│   │   │   └── index.ts
│   │   ├── utils/               # Helper functions
│   │   │   └── index.ts
│   │   ├── app.ts               # Express and Socket.io setup
│   │   └── index.ts             # Application entry point
│   ├── DockerFile
│   ├── package.json
│   ├── tsconfig.json
│   ├── ormconfig.json           # Database configuration
│   ├── jest.config.cjs          # Test configuration
│   ├── nodemon.json             # Dev server configuration
│   └── .nvmrc                   # Node version specification
├── v2/                          # Version 2 (current)
│   ├── src/
│   │   ├── controllers/         # Modular controllers with separated concerns
│   │   │   ├── users.controller.ts
│   │   │   ├── users.queries.ts
│   │   │   ├── users.schemas.ts
│   │   │   ├── messages.controller.ts
│   │   │   ├── messages.queries.ts
│   │   │   ├── messages.schemas.ts
│   │   │   ├── nfts.controller.ts
│   │   │   ├── nfts.queries.ts
│   │   │   ├── nfts.schemas.ts
│   │   │   ├── token.controller.ts
│   │   │   ├── token.queries.ts
│   │   │   ├── token.schemas.ts
│   │   │   └── db.connections.ts
│   │   ├── sockets/             # Separated socket handlers by domain
│   │   │   ├── users.socket.ts
│   │   │   ├── messages.socket.ts
│   │   │   ├── nfts.socket.ts
│   │   │   ├── token.socket.ts
│   │   │   └── index.ts
│   │   ├── services/            # Enhanced service layer
│   │   │   ├── mintNFT.ts
│   │   │   ├── sendNFT.ts
│   │   │   ├── mintToken.ts
│   │   │   ├── sendToken.ts
│   │   │   ├── getBalances.ts
│   │   │   ├── getTokenAccounts.ts
│   │   │   ├── getTokenHolders.ts
│   │   │   └── index.ts
│   │   ├── utils/
│   │   │   └── index.ts
│   │   ├── app.ts
│   │   └── index.ts
│   ├── DockerFile
│   ├── package.json
│   ├── tsconfig.json
│   ├── ormconfig.json
│   ├── jest.config.cjs
│   ├── nodemon.json
│   └── .nvmrc
├── .gitignore
└── README.md
```

## 🚀 Getting Started

### Prerequisites

- **Node.js** v16.x (specified in `.nvmrc`)
- **PostgreSQL** 12+ installed and running
- **npm** or **yarn** package manager
- **Solana CLI** (optional, for wallet management)
- **Git** for version control

### Installation

1. **Clone the repository**

```bash
git clone <repository-url>
cd beenzer-server
```

2. **Choose your version** (v2 recommended)

```bash
cd v2
```

3. **Install Node.js dependencies**

```bash
npm install
# or
yarn install
```

4. **Set Node.js version** (if using nvm)

```bash
nvm use
```

### Environment Setup

Create a `.env` file in your chosen version directory (v1 or v2). Reference the `.env.sample` file for all required variables.

```bash
cp .env.sample .env
```

Edit `.env` with your configuration:

```env
# Server Configuration
PORT=4000

# PostgreSQL Database URLs
USERS_DB_URL=postgresql://username:password@localhost:5432/beenzer_users
MESSAGES_DB_URL=postgresql://username:password@localhost:5432/beenzer_messages
NFTS_DB_URL=postgresql://username:password@localhost:5432/beenzer_nfts
TOKEN_DB_URL=postgresql://username:password@localhost:5432/beenzer_tokens

# Solana Configuration
SOLANA_RPC_URL=https://api.mainnet-beta.solana.com
MASTER_WALLET=<your_solana_public_key>
MASTER_WALLET_KEYPAIR=<your_solana_secret_key_array>

# Google Maps API
GOOGLE_MAPS_API_KEY=<your_google_maps_api_key>

# Optional: Additional Configuration
NODE_ENV=development
```

#### Environment Variables Reference

| Variable | Description | Required | Example |
|----------|-------------|----------|---------|
| `PORT` | Server listening port | Yes | `4000` |
| `USERS_DB_URL` | PostgreSQL connection string for users database | Yes | `postgresql://user:pass@localhost:5432/db` |
| `MESSAGES_DB_URL` | PostgreSQL connection string for messages database | Yes | `postgresql://user:pass@localhost:5432/db` |
| `NFTS_DB_URL` | PostgreSQL connection string for NFTs database | Yes | `postgresql://user:pass@localhost:5432/db` |
| `TOKEN_DB_URL` | PostgreSQL connection string for tokens database | Yes | `postgresql://user:pass@localhost:5432/db` |
| `SOLANA_RPC_URL` | Solana RPC endpoint | Yes | `https://api.mainnet-beta.solana.com` |
| `MASTER_WALLET` | Solana wallet public key (base58) | Yes | `7xKXtg2CW87d97TXJSDpbD5jBkheTqA83TZRuJosgAsU` |
| `MASTER_WALLET_KEYPAIR` | Solana wallet secret key (comma-separated numbers) | Yes | `[123,45,67,...]` |
| `GOOGLE_MAPS_API_KEY` | Google Maps API key for geocoding | Optional | `AIzaSy...` |

⚠️ **Security Warning**: Never commit `.env` files or expose secret keys. Always use environment-specific configurations and keep sensitive data secure.

## 🗄️ Database Setup

### PostgreSQL Databases

Beenzer Server uses four separate PostgreSQL databases:

1. **Users Database** - User profiles, authentication, and social connections
2. **Messages Database** - Direct messages and conversation history
3. **NFTs Database** - NFT metadata, ownership, and transaction records
4. **Tokens Database** - SPL token balances and transfer history

### Database Schema

The database configuration is managed through `ormconfig.json`:

```json
{
  "type": "postgres",
  "host": "host.docker.internal",
  "port": 5432,
  "username": "postgres",
  "password": "postgres",
  "database": "stripe-example",
  "synchronize": true,
  "logging": false
}
```

### Creating Databases

```bash
# Connect to PostgreSQL
psql -U postgres

# Create databases
CREATE DATABASE beenzer_users;
CREATE DATABASE beenzer_messages;
CREATE DATABASE beenzer_nfts;
CREATE DATABASE beenzer_tokens;

# Grant privileges (optional)
GRANT ALL PRIVILEGES ON DATABASE beenzer_users TO your_username;
GRANT ALL PRIVILEGES ON DATABASE beenzer_messages TO your_username;
GRANT ALL PRIVILEGES ON DATABASE beenzer_nfts TO your_username;
GRANT ALL PRIVILEGES ON DATABASE beenzer_tokens TO your_username;
```

### Tables and Schema

Key tables include:

- **Users**: `__pubkey__`, `__username__`, `__appuser__`, timestamps, flags
- **Friends/Connections**: `__pubkey1__`, `__pubkey2__` (relationship table)
- **Messages**: Dynamic tables per conversation (`_pubkey1_pubkey2_`)
- **NFTs**: `__mint__`, `__owner__`, `__creator__`, `__likes__`, metadata
- **Tokens**: `__pubkey__`, `__balance__`, transaction history

Refer to controller schema files (`*.schemas.ts`) for detailed table definitions.

## 💻 Development

### Available Scripts

```bash
# Start development server with hot reload
npm run dev

# Build TypeScript to JavaScript
npm run build

# Start production server
npm start

# Run tests
npm test

# Deploy (build only - v2)
npm run deploy
```

### Development Workflow

1. **Start PostgreSQL** and ensure all databases are created
2. **Configure environment** variables in `.env`
3. **Run development server**:

```bash
cd v2
npm run dev
```

4. **Monitor console output** for:
   - Server port confirmation
   - Database connection status
   - Balance updates (every 15 minutes)
   - Socket.io connection events

### Hot Reload

The development server uses `nodemon` for automatic restarts:

```json
{
  "watch": ["src"],
  "ext": "ts",
  "exec": "ts-node-esm ./src/index.ts"
}
```

### Testing

Jest is configured for TypeScript testing:

```bash
# Run all tests
npm test

# Run tests in watch mode
npm test -- --watch

# Run tests with coverage
npm test -- --coverage
```

Test configuration is in `jest.config.cjs`.

### Common Development Issues

**Port Already in Use**
```bash
# Find process using port 4000
lsof -i :4000
# Kill the process
kill -9 <PID>
```

**Database Connection Failed**
- Verify PostgreSQL is running: `sudo service postgresql status`
- Check database URLs in `.env`
- Ensure databases exist: `psql -U postgres -l`

**Solana RPC Errors**
- Check RPC URL is accessible
- Verify master wallet has sufficient SOL for transactions
- Consider using a dedicated RPC provider for production

## 📡 API Documentation

### REST Endpoints

The server exposes a minimal REST API:

#### Health Check
```http
GET /
Response: <h1>Beenzer Server</h1>
```

### Controller Methods

Most functionality is accessed via Socket.io events, but controllers provide the following methods:

#### Users Controller
- `getUser(pubkey)` - Retrieve user profile
- `newUser(pubkey, username, appuser)` - Create new user
- `updateUser(pubkey, field, value)` - Update user field
- `isUserName(username)` - Check username availability
- `isNewUser(pubkey)` - Check if user exists
- `addFriends(pubkey1, pubkey2)` - Create friendship
- `deleteFriends(pubkey1, pubkey2)` - Remove friendship
- `getFriends(pubkey)` - Get user's friends list
- `getUserFriends(pubkey)` - Get friends with full details
- `searchUsers(query)` - Search users by username
- `getUsersFlags()` - Get flagged users

#### Messages Controller
- `createMessages(tableName)` - Create conversation table
- `deleteMessages(tableName)` - Delete conversation
- `getMessages(tableName)` - Retrieve messages
- `newMessage(tableName, sender, message, timestamp)` - Send message

#### NFTs Controller
- `getNFT(mint)` - Get NFT details
- `newNFT(nftData)` - Register new NFT
- `updateNFTOwner(mint, newOwner)` - Transfer ownership
- `updateNFTLikes(mint, likes)` - Update like count
- `createNFTtransactions(txData)` - Record transaction

#### Token Controller
- `getTokenBalance(pubkey)` - Get user token balance
- `updateTokenBalance(pubkey, amount)` - Update balance
- `getTokenHolders()` - List all token holders
- `getTokenAccounts()` - Get token account information

## 🔌 Socket.io Events

### Client → Server Events

#### User Events
```typescript
socket.emit('newUser', { pubkey, username, appuser })
socket.emit('updateUser', { pubkey, field, value })
socket.emit('usernameExists', { username })
socket.emit('searchUsers', { query })
```

#### Friend/Connection Events
```typescript
socket.emit('addFriend', { pubkey1, pubkey2 })
socket.emit('deleteFriend', { pubkey1, pubkey2 })
socket.emit('getFriends', { pubkey })
```

#### Message Events
```typescript
socket.emit('sendMessage', { table, sender, message, timestamp })
socket.emit('getMessages', { table })
```

#### NFT Events
```typescript
socket.emit('mintNFT', { nftData })
socket.emit('sendNFT', { mint, from, to })
socket.emit('getNFT', { mint })
socket.emit('likeNFT', { mint })
```

#### Token Events
```typescript
socket.emit('mintToken', { recipient, amount })
socket.emit('sendToken', { from, to, amount })
socket.emit('getBalance', { pubkey })
socket.emit('getTokenHolders')
```

### Server → Client Events

#### Connection Events
```typescript
socket.on('serverConnection', (message) => {})
socket.on('nUsers', (count) => {})  // Active user count
```

#### Response Events
```typescript
socket.on('userCreated', (data) => {})
socket.on('userUpdated', (data) => {})
socket.on('friendsData', (friends) => {})
socket.on('messagesData', (messages) => {})
socket.on('nftData', (nft) => {})
socket.on('balanceData', (balance) => {})
```

#### Progress Events (NFT Minting)
```typescript
socket.on('mintLogs', (log) => {})  // Progress updates during minting
```

### Connection Management

```javascript
import io from 'socket.io-client';

const socket = io('http://localhost:4000');

socket.on('connect', () => {
  console.log('Connected to Beenzer Server');
});

socket.on('serverConnection', (message) => {
  console.log(message);
});

socket.on('nUsers', (count) => {
  console.log(`Active users: ${count}`);
});

socket.on('disconnect', () => {
  console.log('Disconnected from server');
});
```

## 🐳 Docker Deployment

### Building the Image

```bash
# From v1 or v2 directory
cd v2
docker build -t beenzer-server:latest .
```

### Running the Container

```bash
docker run -d \
  --name beenzer-server \
  -p 4000:4000 \
  --env-file .env \
  beenzer-server:latest
```

### Docker Configuration

The `DockerFile` includes:
- Node.js base image
- Dependency installation
- TypeScript compilation
- Production build setup
- Port 4000 exposure

### Docker Compose (Recommended)

Create `docker-compose.yml` for full stack deployment:

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:14
    environment:
      POSTGRES_PASSWORD: postgres
      POSTGRES_USER: postgres
    ports:
      - "5432:5432"
    volumes:
      - postgres_data:/var/lib/postgresql/data

  beenzer-server:
    build: ./v2
    ports:
      - "4000:4000"
    depends_on:
      - postgres
    env_file:
      - ./v2/.env
    restart: unless-stopped

volumes:
  postgres_data:
```

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f beenzer-server

# Stop services
docker-compose down
```

### Production Considerations

- Use environment-specific `.env` files
- Implement proper secrets management
- Set up reverse proxy (nginx/traefik) for HTTPS
- Configure database backups
- Monitor resource usage and scale horizontally
- Implement health checks and restart policies
- Use managed Solana RPC providers for reliability

## 🤝 Contributing

Contributions are welcome! Please follow these guidelines:

### Development Process

1. **Fork the repository**
2. **Create a feature branch**
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Make your changes**
4. **Run tests**
   ```bash
   npm test
   ```
5. **Commit your changes**
   ```bash
   git commit -m "✨ Add amazing feature"
   ```
6. **Push to your branch**
   ```bash
   git push origin feature/amazing-feature
   ```
7. **Open a Pull Request**

### Code Style

- Follow existing TypeScript patterns
- Use meaningful variable and function names
- Add JSDoc comments for public methods
- Keep functions focused and concise
- Handle errors gracefully with try/catch
- Use async/await for asynchronous operations

### Commit Messages

Use conventional commit format with emojis:

- ✨ `:sparkles:` - New feature
- 🐛 `:bug:` - Bug fix
- 📝 `:memo:` - Documentation
- ♻️ `:recycle:` - Refactoring
- ✅ `:white_check_mark:` - Tests
- 🔧 `:wrench:` - Configuration

### Pull Request Guidelines

- Keep PRs focused on a single concern
- Update documentation as needed
- Add tests for new functionality
- Ensure all tests pass
- Update CHANGELOG if applicable

## 📄 License

This project is licensed under the **ISC License**.

```text
ISC License

Copyright (c) 2024 Beenzer

Permission to use, copy, modify, and/or distribute this software for any
purpose with or without fee is hereby granted, provided that the above
copyright notice and this permission notice appear in all copies.

THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH
REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY
AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER RESULTING FROM
LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT, NEGLIGENCE OR
OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR
PERFORMANCE OF THIS SOFTWARE.
```

## 🔗 Related Projects

### Beenzer Ecosystem

- **[beenzer-expo](../beenzer-expo)** - React Native mobile application
- **[beenzer-landing](../beenzer-landing)** - Marketing website and landing page
- **[beenzer-dao](../beenzer-dao)** - DAO governance contracts and tooling

### Documentation

- [Solana Developer Docs](https://docs.solana.com/)
- [Metaplex Docs](https://docs.metaplex.com/)
- [Socket.io Documentation](https://socket.io/docs/v4/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)

## 📞 Support

For questions, issues, or contributions:

- **Issues**: [GitHub Issues](issues)
- **Discussions**: [GitHub Discussions](discussions)

---

**Built with ❤️ for the Solana ecosystem**

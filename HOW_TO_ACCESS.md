# 🚀 How to Access GoHacker

## Step 1: Start the Application

```bash
cd /home/jeff/gohacker
task run
```

Wait for the message: **"✅ GoHacker is running!"**

## Step 2: Open Your Browser

Open your web browser and go to:

```
https://localhost:3000
```

⚠️ **Important**: You'll see a security warning because we use self-signed SSL certificates. This is normal for development!

### How to Bypass the Security Warning:

**Chrome/Brave:**
1. Click "Advanced"
2. Click "Proceed to localhost (unsafe)"

**Firefox:**
1. Click "Advanced"
2. Click "Accept the Risk and Continue"

**Safari:**
1. Click "Show Details"
2. Click "visit this website"

## Step 3: What You'll See

### 🎮 Landing Page

You'll see a **cyberpunk-themed landing page** with:

- **GoHacker logo** (glowing green)
- **"Learn Go by Hacking"** subtitle
- **API Health Status** (should show "✅ Connected!")
- **Feature cards**: 18+ Challenges, Achievements, Live Coding, Secure
- **Quick Start code examples**
- **Test API Connection button**

### 🔧 Test the API

Click the **"Test API Connection"** button on the page to verify the backend is working.

You should see a JSON response like:
```json
{
  "status": "ok",
  "message": "GoHacker API is running! 🚀"
}
```

## Step 4: Use the API

### Option 1: Using curl (Terminal)

```bash
# Health check
curl -k https://localhost:3000/api/health

# Register a user
curl -k -X POST https://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "hacker1",
    "email": "test@example.com",
    "password": "password123"
  }'

# Login (save the token!)
curl -k -X POST https://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "hacker1",
    "password": "password123"
  }'

# Get challenges (use your token)
curl -k https://localhost:3000/api/challenges \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

### Option 2: Using Postman or Insomnia

1. Import the base URL: `https://localhost:3000`
2. Disable SSL verification in settings
3. Start making requests to the API endpoints

### Option 3: Using the Browser Console

Open the browser console (F12) and run:

```javascript
// Test API
fetch('/api/health')
  .then(r => r.json())
  .then(data => console.log(data));

// Register
fetch('/api/auth/register', {
  method: 'POST',
  headers: {'Content-Type': 'application/json'},
  body: JSON.stringify({
    username: 'hacker1',
    email: 'test@example.com',
    password: 'password123'
  })
}).then(r => r.json()).then(data => console.log(data));
```

## What You Can Do

### 📊 Available API Endpoints

**Public (no auth required):**
- `GET /api/health` - Check if API is running
- `POST /api/auth/register` - Create account
- `POST /api/auth/login` - Login and get JWT token

**Protected (requires JWT token):**
- `GET /api/auth/me` - Get your user info
- `GET /api/challenges` - List all 18 challenges
- `GET /api/challenges/:id` - Get specific challenge
- `POST /api/challenges/:id/execute` - Submit Go code
- `GET /api/progress` - Your progress & stats
- `GET /api/progress/achievements` - Your achievements
- `GET /api/leaderboard` - Top users

### 🎯 Try Your First Challenge

1. **Register** a user
2. **Login** to get your JWT token
3. **Get challenges** to see Challenge 001
4. **Execute code** to solve it:

```bash
TOKEN="your-jwt-token-here"

curl -k -X POST https://localhost:3000/api/challenges/001/execute \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "package main\nimport \"fmt\"\nfunc main() {\n\tvar password string = \"admin123\"\n\tfmt.Println(password)\n}"
  }'
```

## Troubleshooting

### Can't access https://localhost:3000?

1. Check if services are running:
   ```bash
   docker ps
   ```

2. Check logs:
   ```bash
   task logs
   ```

3. Restart services:
   ```bash
   task down
   task run
   ```

### Port already in use?

```bash
# Stop services
task down

# Or kill the port
sudo lsof -ti:3000 | xargs kill -9
```

### Docker not running?

```bash
# Start Docker
sudo systemctl start docker

# Then try again
task run
```

## Next Steps

1. ✅ Access https://localhost:3000
2. ✅ See the landing page
3. ✅ Test API connection
4. ✅ Register a user via API
5. ✅ Login and get JWT token
6. ✅ Get list of challenges
7. ✅ Solve your first challenge!
8. 🎮 Build the full Vue.js frontend (optional)

## Full Vue.js Frontend (Coming Soon)

The current landing page is a simple HTML page. To build the full Vue.js frontend with:
- User registration/login UI
- Challenge browser
- Code editor (Monaco)
- Progress dashboard
- Achievement gallery
- Leaderboard

You'll need Node.js installed, then:
```bash
cd frontend
npm install
npm run dev
```

---

**Happy Hacking! 🎮💻✨**

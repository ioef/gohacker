# 🎮 How to Play GoHacker

## Current Status

✅ **Backend is running** - You can see the landing page at https://localhost:3000
✅ **API is working** - The "Test API Connection" button shows it's connected
✅ **18 challenges are loaded** - Ready to play!

## How to Play (Using the API)

The current landing page is just a **status page**. To actually play the game, you interact with the **REST API**. Here's how:

### Step 1: Register a User

```bash
curl -k -X POST https://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "player1",
    "email": "player1@example.com",
    "password": "mypassword123"
  }'
```

You'll get a response with your user info and a JWT token.

### Step 2: Login (Get Your Token)

```bash
curl -k -X POST https://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "player1",
    "password": "mypassword123"
  }'
```

**Save the token from the response!** You'll need it for all other requests.

### Step 3: Get List of Challenges

```bash
# Replace YOUR_TOKEN with the token from login
TOKEN="your-jwt-token-here"

curl -k https://localhost:3000/api/challenges \
  -H "Authorization: Bearer $TOKEN"
```

You'll see all 18 challenges with their descriptions!

### Step 4: View a Specific Challenge

```bash
curl -k https://localhost:3000/api/challenges/001 \
  -H "Authorization: Bearer $TOKEN"
```

This shows Challenge 001: "First Breach" - Extract a password from a variable.

### Step 5: Solve the Challenge!

Write Go code to solve it:

```bash
curl -k -X POST https://localhost:3000/api/challenges/001/execute \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "package main\nimport \"fmt\"\nfunc main() {\n\tvar password string = \"admin123\"\n\tfmt.Println(password)\n}"
  }'
```

If correct, you'll get:
- ✅ Success message
- 🎯 XP points
- 🏆 Possible achievement unlock
- 📈 Level up (if you earned enough XP)

### Step 6: Check Your Progress

```bash
# View your stats
curl -k https://localhost:3000/api/progress \
  -H "Authorization: Bearer $TOKEN"

# View achievements
curl -k https://localhost:3000/api/progress/achievements \
  -H "Authorization: Bearer $TOKEN"

# View leaderboard
curl -k https://localhost:3000/api/leaderboard \
  -H "Authorization: Bearer $TOKEN"
```

## 🎯 All 18 Challenges

1. **001** - First Breach (Variables)
2. **002** - Loop Breaker (For loops)
3. **003** - Conditional Access (If/else)
4. **004** - Function Exploit (Functions)
5. **005** - Array Indexer (Arrays)
6. **006** - Slice Manipulator (Slices)
7. **007** - Map Decoder (Maps)
8. **008** - String Reverser (Strings)
9. **009** - Type Converter (Type conversion)
10. **010** - Range Iterator (Range loops)
11. **011** - Struct Builder (Structs)
12. **012** - Method Caller (Methods)
13. **013** - Pointer Tracker (Pointers)
14. **014** - Interface Implementer (Interfaces)
15. **015** - Error Handler (Error handling)
16. **016** - Goroutine Launcher (Goroutines)
17. **017** - Channel Communicator (Channels)
18. **018** - WaitGroup Synchronizer (WaitGroups)

## 🎮 Want a Full UI?

The current setup uses the **REST API** directly. To build a full game UI with:
- 🖥️ Visual challenge browser
- ✍️ Code editor (Monaco)
- 📊 Progress dashboard
- 🏆 Achievement gallery
- 🏅 Leaderboard

You would need to build the Vue.js frontend (requires Node.js):

```bash
cd /home/jeff/gohacker/frontend
npm install
npm run dev
```

But for now, you can play the entire game using curl or tools like:
- **Postman** - Visual API client
- **Insomnia** - REST client
- **HTTPie** - Better curl
- **Browser Console** - Use fetch() in DevTools

## 📝 Example Game Session

```bash
# 1. Register
curl -k -X POST https://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"hacker1","email":"h@ex.com","password":"pass123"}'

# 2. Login (save the token!)
TOKEN=$(curl -k -X POST https://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"hacker1","password":"pass123"}' | jq -r '.token')

# 3. Get challenges
curl -k https://localhost:3000/api/challenges \
  -H "Authorization: Bearer $TOKEN" | jq

# 4. Solve Challenge 001
curl -k -X POST https://localhost:3000/api/challenges/001/execute \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"code":"package main\nimport \"fmt\"\nfunc main(){\n\tvar password string=\"admin123\"\n\tfmt.Println(password)\n}"}' | jq

# 5. Check progress
curl -k https://localhost:3000/api/progress \
  -H "Authorization: Bearer $TOKEN" | jq
```

## 🎯 Quick Test

Try this complete example:

```bash
cd /home/jeff/gohacker

# Create a test script
cat > test-game.sh << 'EOF'
#!/bin/bash
echo "🎮 Testing GoHacker Game..."

# Register
echo "1. Registering user..."
curl -k -X POST https://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"testplayer","email":"test@test.com","password":"test123"}' \
  2>/dev/null | jq

# Login
echo -e "\n2. Logging in..."
TOKEN=$(curl -k -X POST https://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testplayer","password":"test123"}' \
  2>/dev/null | jq -r '.token')

echo "Token: $TOKEN"

# Get challenges
echo -e "\n3. Getting challenges..."
curl -k https://localhost:3000/api/challenges \
  -H "Authorization: Bearer $TOKEN" \
  2>/dev/null | jq '.[] | {id, title, difficulty}'

echo -e "\n✅ Game is working! You can now solve challenges!"
EOF

chmod +x test-game.sh
./test-game.sh
```

---

**The game is fully functional via the API! The landing page is just a status indicator.** 🎮

# GoHacker Database Backup Guide

## 📊 Database Overview

GoHacker uses an **in-memory database with automatic disk persistence**:
- Data is stored in memory for fast access
- Automatically saved to disk every 5 minutes
- Saved on graceful shutdown
- Stored in JSON format for easy backup

## 💾 Where is the Data Stored?

### Docker Volume
The database is stored in a Docker volume named `gohacker_backend-data`:
```bash
docker volume ls | grep gohacker
# Output: gohacker_backend-data
```

### File Location
Inside the container: `/app/data/gohacker.json`

## 🔍 View Current Data

### Check if data exists:
```bash
docker exec gohacker-backend ls -la /app/data/
```

### View database content:
```bash
docker exec gohacker-backend cat /app/data/gohacker.json
```

### View summary (requires jq):
```bash
docker exec gohacker-backend cat /app/data/gohacker.json | jq '{
  last_saved: .last_saved,
  total_users: (.users | length),
  total_progress: (.progress | length),
  total_challenges_completed: (.challenge_progress | length)
}'
```

## 💾 Backup Methods

### Method 1: Copy from Docker Volume (Recommended)

**Backup:**
```bash
# Create backup directory
mkdir -p ~/gohacker-backups

# Copy database file
docker cp gohacker-backend:/app/data/gohacker.json ~/gohacker-backups/gohacker-$(date +%Y%m%d-%H%M%S).json

# Or use docker volume backup
docker run --rm -v gohacker_backend-data:/data -v ~/gohacker-backups:/backup \
  debian:bookworm-slim tar czf /backup/gohacker-data-$(date +%Y%m%d-%H%M%S).tar.gz -C /data .
```

**Restore:**
```bash
# Copy backup back to container
docker cp ~/gohacker-backups/gohacker-20260628-120000.json gohacker-backend:/app/data/gohacker.json

# Restart container to load data
docker restart gohacker-backend
```

### Method 2: Mount Local Directory

**Modify docker-compose.yml** to use a local directory instead of a volume:

```yaml
services:
  backend:
    volumes:
      # Replace this line:
      # - backend-data:/app/data
      # With this:
      - ./data:/app/data
```

Then your data will be in `/home/jeff/gohacker/data/gohacker.json` on your host machine.

### Method 3: Automated Backup Script

Create `backup.sh`:
```bash
#!/bin/bash
BACKUP_DIR=~/gohacker-backups
mkdir -p $BACKUP_DIR

# Create timestamped backup
TIMESTAMP=$(date +%Y%m%d-%H%M%S)
docker cp gohacker-backend:/app/data/gohacker.json \
  $BACKUP_DIR/gohacker-$TIMESTAMP.json

# Keep only last 30 backups
cd $BACKUP_DIR
ls -t gohacker-*.json | tail -n +31 | xargs -r rm

echo "Backup created: gohacker-$TIMESTAMP.json"
```

Make it executable and run:
```bash
chmod +x backup.sh
./backup.sh
```

Add to crontab for daily backups:
```bash
# Run daily at 2 AM
0 2 * * * /home/jeff/gohacker/backup.sh
```

## 🔄 Auto-Save Configuration

The database auto-saves every **5 minutes** by default. This is configured in `backend/main.go`:

```go
pm.StartAutoSave(5 * time.Minute)
```

To change the interval, modify this value and rebuild the container.

## 📋 What Data is Saved?

The `gohacker.json` file contains:
- **Users**: All registered users with credentials and stats
- **Progress**: User progress, XP, levels, ranks
- **Challenge Progress**: Completed challenges per user
- **Achievements**: Unlocked achievements
- **Last Saved**: Timestamp of last save

## 🚨 Important Notes

1. **Data persists** across container restarts
2. **Data is lost** if you run `docker compose down -v` (removes volumes)
3. **Always backup** before major updates
4. **JSON format** makes it easy to inspect and edit manually
5. **Automatic saves** happen every 5 minutes + on shutdown

## 🔧 Troubleshooting

### Data not persisting?
```bash
# Check if volume exists
docker volume inspect gohacker_backend-data

# Check if file exists in container
docker exec gohacker-backend ls -la /app/data/

# Check container logs
docker logs gohacker-backend | grep -i "save\|load\|persist"
```

### Restore from backup
```bash
# Stop containers
docker compose down

# Copy backup to volume
docker run --rm -v gohacker_backend-data:/data -v ~/gohacker-backups:/backup \
  debian:bookworm-slim cp /backup/gohacker-20260628-120000.json /data/gohacker.json

# Start containers
docker compose up -d
```

## 📊 Database Statistics

View your database stats:
```bash
docker exec gohacker-backend cat /app/data/gohacker.json | jq '{
  users: (.users | length),
  progress_records: (.progress | length),
  challenges_completed: (.challenge_progress | length),
  last_saved: .last_saved
}'
```

## 🎯 Best Practices

1. **Regular Backups**: Set up automated daily backups
2. **Before Updates**: Always backup before updating the application
3. **Test Restores**: Periodically test your backup restoration process
4. **Multiple Copies**: Keep backups in multiple locations
5. **Monitor Size**: Check database file size growth over time

---

**Your data is safe!** The database automatically saves to disk and persists in a Docker volume. 🎮💾

# Grant Execution Permission
# Make the script executable by running:
# chmod +x cleanup_docker_compose.sh

#!/bin/bash

# Stop and remove all Docker Compose services
echo "Stopping and removing Docker Compose services..."
docker-compose down -v --rmi all

# Remove all unused volumes
echo "Removing unused volumes..."
docker volume prune -f

# Remove all unused networks
echo "Removing unused networks..."
docker network prune -f

# Remove all unused images
echo "Removing unused images..."
docker image prune -af

# Clean up builder cache
echo "Cleaning up builder cache..."
docker builder prune -f

echo "Docker Compose cleanup completed!"
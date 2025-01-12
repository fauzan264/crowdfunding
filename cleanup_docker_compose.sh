# Grant Execution Permission
# Make the script executable by running:
# chmod +x cleanup_docker_compose.sh

#!/bin/bash

# Check if docker-compose.yml exists in the current directory
if [ ! -f "docker-compose.yml" ]; then
    echo "Error: No docker-compose.yml file found in the current directory."
    echo "Please run this script in a directory with a valid docker-compose.yml file."
    exit 1
fi

# Stop and remove Docker Compose services for the current project
echo "Stopping and removing Docker Compose services for the current project..."
docker-compose down -v --rmi all

# Remove builder cache (global cleanup, but usually minimal impact)
echo "Cleaning up builder cache..."
docker builder prune -f

echo "Docker Compose cleanup for the current project completed!"
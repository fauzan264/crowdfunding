# Use Node.js 18 (or any version >= 16)
FROM node:18-alpine AS builder

# Set working directory
WORKDIR /usr/src/app

# Copy package.json and package-lock.json first to leverage Docker cache
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application files
COPY . .

# Build the Nuxt application (you must run this step to create .output and nitro.json)
RUN npx nuxi build

# Final stage for production image
FROM node:18-alpine

# Set the working directory
WORKDIR /usr/src/app

# Copy necessary files from the builder stage
COPY --from=builder /usr/src/app /usr/src/app

# Expose the port that Nuxt will use
EXPOSE 3000

# Start the Nuxt application in development mode
CMD ["npm", "run", "dev"]

# Start the Nuxt application in production mode
# CMD ["npm", "run", "start"]

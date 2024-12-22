FROM golang:1.20-alpine

# Define build arguments
ARG DB_USER
ARG DB_PASSWORD
ARG DB_NAME

# working directory
WORKDIR /app
COPY go.mod go.sum ./

# download go dependencies
RUN go mod download

# copy the entire backend project into the container
COPY . .

EXPOSE 8080

# command to run the application in production mode
CMD ["sh", "-c", "migrate -path ./migrations -database mysql://${DB_USER}:${DB_PASSWORD}@tcp(db:3306)/${DB_NAME} up && ./main"]
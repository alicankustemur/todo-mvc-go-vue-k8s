# Create build stage, this stage will just build and create "todo" named a binary file.
FROM golang:1.13-alpine as BUILDER

# Add git for fetch dependencies
RUN apk add \
        ca-certificates \
        git \
        build-base \
        --no-cache

WORKDIR /build/

# Copy Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY src/todo ./
RUN go build -v -o /build/todo

# Create runtime stage, this stage will just run builded "todo" binary file
FROM alpine:3.9 as RUNTIME

# This is important because application will not work.
# It gives this error : Env var 'PORT' must be set
ENV PORT 8000

WORKDIR /root

# Copy "todo" binary file from BUILDER stage
COPY --from=BUILDER /build/todo .

EXPOSE $PORT

ENTRYPOINT ["./todo"]
#!/bin/bash

# Create the main directories
mkdir -p bin cmd docs internal

# Create the nested directories
mkdir -p internal/adapter/cache/redis
mkdir -p internal/adapter/handler/http
mkdir -p internal/adapter/repository/postgres/migrations
mkdir -p internal/adapter/token/paseto
mkdir -p internal/core/domain
mkdir -p internal/core/port
mkdir -p internal/core/service
mkdir -p internal/core/util

# Print success message
echo "Folder structure created successfully!"
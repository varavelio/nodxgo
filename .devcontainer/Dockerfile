# Use golang (debian trixie) as the base image
FROM golang:1.26-trixie

# Install Deno
COPY --from=denoland/deno:bin-2.9.5 /deno /usr/local/bin/deno

# Install golangci-lint
COPY --from=golangci/golangci-lint:v2.12.2 /usr/bin/golangci-lint /usr/local/bin/golangci-lint

# Install Veta
COPY --from=varavel/veta:0.1.1 /usr/local/bin/veta /usr/local/bin/veta

# Set environment variables
ENV \
  DEBIAN_FRONTEND="noninteractive" \
  PIP_BREAK_SYSTEM_PACKAGES=1 \
  CGO_ENABLED=0

RUN set -e && \
  # Update and install system dependencies
  apt-get update -qq && \
  apt-get install -yqq --no-install-recommends \
  ca-certificates wget curl zip unzip p7zip-full tzdata git tree ripgrep \
  python3 python3-pip && \
  # Install tools from npm
  deno install --global --allow-all --allow-scripts --quiet npm:@go-task/cli@latest && \
  # Git config
  git config --global --add safe.directory '*' && \
  # Cleanup
  rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Just to keep the container running
CMD ["sleep", "infinity"]

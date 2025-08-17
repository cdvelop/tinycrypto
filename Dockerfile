FROM tinygo/tinygo:0.38.0

# Install Playwright dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    libgtk-4-1 \
    libgraphene-1.0-0 \
    libxslt1.1 \
    libwoff1 \
    libvpx9 \
    libevent-2.1-7 \
    libopus0 \
    libgstreamer1.0-0 \
    libgstreamer-plugins-base1.0-0 \
    libflite1 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go run github.com/playwright-community/playwright-go/cmd/playwright install --with-deps

CMD ["go", "test", "-v", "./..."]

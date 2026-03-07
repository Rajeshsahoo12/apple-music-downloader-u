#!/bin/bash

# CCExtractor installation script for EIA-608 caption extraction

echo "========================================="
echo "Installing CCExtractor from source"
echo "========================================="

# Install dependencies
echo "[1/5] Installing build dependencies..."
sudo apt-get install -y libclang-dev clang libtesseract-dev

if [ $? -ne 0 ]; then
    echo "Failed to install dependencies"
    exit 1
fi

# Install Rust/Cargo (required by CCExtractor)
echo ""
echo "[2/5] Installing Rust/Cargo..."
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y
source "$HOME/.cargo/env"
cargo --version

if [ $? -ne 0 ]; then
    echo "Failed to install Rust/Cargo"
    exit 1
fi

# Clone CCExtractor repository
echo ""
echo "[3/5] Cloning CCExtractor repository..."
git clone https://github.com/CCExtractor/ccextractor
cd ccextractor/linux
./build
cd ../..

if [ $? -ne 0 ]; then
    echo "Failed to clone/build CCExtractor"
    exit 1
fi

# Install CCExtractor
echo ""
echo "[4/5] Installing CCExtractor..."
sudo cp ccextractor/linux/ccextractor /usr/local/bin/
sudo chmod +x /usr/local/bin/ccextractor

if [ $? -ne 0 ]; then
    echo "Failed to install CCExtractor"
    exit 1
fi

# Verify installation
echo ""
echo "[5/5] Verifying installation..."
ccextractor --version

if [ $? -eq 0 ]; then
    echo ""
    echo "========================================="
    echo "✓ CCExtractor installed successfully!"
    echo "========================================="
    echo ""
    echo "Location: /usr/local/bin/ccextractor"
    echo ""
    echo "Next steps:"
    echo "1. Rebuild apple-music-downloader: go build -o apple-music-downloader main.go"
    echo "2. Test extraction: ./apple-music-downloader <music-video-url>"

    exit 0
else
    echo "Installation verification failed"
    exit 1
fi

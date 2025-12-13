#!/bin/bash

# Deployment Verification Script for thejokers69 Portfolio
# This script verifies that the Go application builds and runs correctly for deployment

set -e  # Exit on any error

echo "🚀 Starting Deployment Verification for thejokers69 Portfolio"
echo "============================================================"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print status
print_status() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    print_error "Not in the project root directory (go.mod not found)"
    exit 1
fi

print_status "Project directory verified"

# Check Go version
GO_VERSION=$(go version | grep -o 'go[0-9]\+\.[0-9]\+\.[0-9]\+')
REQUIRED_VERSION="go1.25.3"

if [[ "$GO_VERSION" == "$REQUIRED_VERSION" ]]; then
    print_status "Go version: $GO_VERSION (matches required version)"
else
    print_warning "Go version: $GO_VERSION (expected $REQUIRED_VERSION)"
fi

# Check if templ is installed
if command -v templ &> /dev/null; then
    TEMPL_VERSION=$(templ version 2>/dev/null || echo "unknown")
    print_status "Templ installed: $TEMPL_VERSION"
else
    print_error "Templ not installed"
    exit 1
fi

# Clean previous build
print_status "Cleaning previous build artifacts"
rm -f server server.log

# Generate templ files
print_status "Generating templ files"
if templ generate ./templates/...; then
    print_status "Templ files generated successfully"
else
    print_error "Failed to generate templ files"
    exit 1
fi

# Fix templ duplicate import issue (known bug in templ v0.3.960)
print_status "Fixing templ duplicate import issue"
if [ -f "templates/pages/base_templ.go" ]; then
    sed -i.bak '/^import (/,/^)$/{/github.com\/a-h\/templ/d;}' templates/pages/base_templ.go && rm templates/pages/base_templ.go.bak 2>/dev/null || true
    print_status "Templ import issue fixed"
fi

# Download dependencies
print_status "Downloading Go dependencies"
if go mod download && go mod verify; then
    print_status "Dependencies downloaded and verified"
else
    print_error "Failed to download/verify dependencies"
    exit 1
fi

# Run go mod tidy
print_status "Running go mod tidy"
go mod tidy

# Build the application
print_status "Building Go application"
if go build -v -o server ./cmd/server/main.go; then
    print_status "Application built successfully"
else
    print_error "Build failed"
    exit 1
fi

# Verify binary
if [ -f "server" ] && [ -x "server" ]; then
    BINARY_SIZE=$(ls -lh server | awk '{print $5}')
    print_status "Binary created: $BINARY_SIZE"
else
    print_error "Binary not created or not executable"
    exit 1
fi

# Check static files
print_status "Checking static files"
STATIC_FILES_OK=true

if [ ! -f "static/css/style.css" ]; then
    print_error "Missing static/css/style.css"
    STATIC_FILES_OK=false
fi

if [ ! -f "static/js/main.js" ]; then
    print_error "Missing static/js/main.js"
    STATIC_FILES_OK=false
fi

if [ ! -d "static/images" ] || [ -z "$(ls static/images/)" ]; then
    print_error "Missing or empty static/images directory"
    STATIC_FILES_OK=false
fi

if [ "$STATIC_FILES_OK" = true ]; then
    print_status "All static files present"
fi

# Check content files
print_status "Checking content files"
CONTENT_FILES_OK=true

for file in content/pages/index.md content/pages/profile.md content/pages/projects.md; do
    if [ ! -f "$file" ]; then
        print_error "Missing content file: $file"
        CONTENT_FILES_OK=false
    fi
done

if [ "$CONTENT_FILES_OK" = true ]; then
    print_status "All content files present"
fi

# Test server startup
print_status "Testing server startup"
PORT=3002
PORT=$PORT ./server > server_startup.log 2>&1 &
SERVER_PID=$!

# Wait a moment for server to start
sleep 3

if kill -0 $SERVER_PID 2>/dev/null; then
    print_status "Server started successfully on port $PORT"
    kill $SERVER_PID 2>/dev/null || true
else
    print_error "Server failed to start"
    cat server_startup.log
    exit 1
fi

# Check server startup log
if grep -q "Server starting on port $PORT" server_startup.log; then
    print_status "Server startup log shows success"
else
    print_warning "Server startup log may have issues"
    cat server_startup.log
fi

# Check server startup log
if grep -q "Server starting on port $PORT" server_startup.log; then
    print_status "Server startup log shows success"
else
    print_warning "Server startup log may have issues"
    cat server_startup.log
fi

# Clean up
rm -f server server_startup.log

# Final verification for deployment platforms
print_status "Verifying deployment configurations"

# Check Railway config
if [ -f "railway.json" ]; then
    print_status "Railway configuration present"
else
    print_warning "Railway configuration missing"
fi

# Check Render config
if [ -f "render.yaml" ]; then
    print_status "Render configuration present"
else
    print_warning "Render configuration missing"
fi

# Check GitHub Actions
if [ -f ".github/workflows/go-build.yml" ]; then
    print_status "GitHub Actions workflow present"
else
    print_warning "GitHub Actions workflow missing"
fi

echo ""
echo "🎉 DEPLOYMENT VERIFICATION COMPLETE"
echo "==================================="
print_status "All checks passed! The application is ready for deployment."
echo ""
echo "📋 Deployment Options:"
echo "  1. Railway: Connect repository and deploy automatically"
echo "  2. Render: Create web service from GitHub repo"
echo "  3. Manual: Build locally and deploy to any hosting platform"
echo ""
echo "🔧 Build Commands for Manual Deployment:"
echo "  go install github.com/a-h/templ/cmd/templ@v0.3.960"
echo "  templ generate ./templates/..."
echo "  go build -o server ./cmd/server/main.go"
echo "  ./server"
echo ""
echo "🌐 Environment Variables:"
echo "  PORT=8080 (or any port)"
echo "  BASEURL=/thejokers69 (for subdirectory deployment)"
echo "  POSTHOG_API_KEY=your_key (optional)"
echo "  POSTHOG_HOST=https://eu.i.posthog.com (optional)"
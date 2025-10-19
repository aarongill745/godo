#!/bin/bash
# Development script for godo
# Manages test workspaces and rebuilds the binary

set -e  # Exit on error

GODO_BIN="./godo"
TEST_DIR="test-workspace"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to show usage
show_help() {
    echo "Usage: ./dev.sh [command]"
    echo ""
    echo "Commands:"
    echo "  build          Build the godo binary"
    echo "  clean          Remove test workspaces and binary"
    echo "  test           Create fresh test workspace and rebuild"
    echo "  run [args]     Build and run godo with arguments in test workspace"
    echo "  help           Show this help message"
    echo ""
    echo "Examples:"
    echo "  ./dev.sh test                    # Create fresh test workspace"
    echo "  ./dev.sh run init                # Run 'godo init' in test workspace"
    echo "  ./dev.sh run add \"my task\" +++  # Add a task in test workspace"
    echo "  ./dev.sh clean                   # Clean up everything"
}

# Build the binary
build() {
    echo -e "${BLUE}Building godo...${NC}"
    go build -o "$GODO_BIN"
    echo -e "${GREEN}✓ Build complete${NC}"
}

# Clean up test workspaces and binary
clean() {
    echo -e "${BLUE}Cleaning up...${NC}"
    rm -rf "$TEST_DIR"
    rm -f "$GODO_BIN"
    echo -e "${GREEN}✓ Cleaned test workspace and binary${NC}"
}

# Create fresh test workspace
create_test_workspace() {
    echo -e "${BLUE}Creating fresh test workspace...${NC}"
    rm -rf "$TEST_DIR"
    mkdir -p "$TEST_DIR"
    echo -e "${GREEN}✓ Test workspace created at ./$TEST_DIR${NC}"
}

# Run godo in test workspace
run_in_test() {
    build
    create_test_workspace
    cd "$TEST_DIR"
    echo -e "${BLUE}Running: godo $@${NC}"
    echo ""
    ../"$GODO_BIN" "$@"
}

# Main command dispatcher
case "$1" in
    build)
        build
        ;;
    clean)
        clean
        ;;
    test)
        build
        create_test_workspace
        echo ""
        echo "Test workspace ready. Run commands with:"
        echo "  cd $TEST_DIR && ../$GODO_BIN [command]"
        ;;
    run)
        shift  # Remove 'run' from arguments
        run_in_test "$@"
        ;;
    help|--help|-h)
        show_help
        ;;
    *)
        if [ -z "$1" ]; then
            show_help
        else
            echo -e "${RED}Unknown command: $1${NC}"
            echo ""
            show_help
            exit 1
        fi
        ;;
esac
#!/bin/bash

./build.sh
echo ""
echo "-------------------"
echo ""
echo "Running Project..."
echo ""
echo "-------------------"
echo ""
./bin/backup-cli "$@"

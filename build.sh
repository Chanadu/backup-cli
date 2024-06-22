go mod tidy
echo "Building Project..."
go build -o ./bin/
echo "Build success."
echo "Running Project..."
echo ""
echo "-------------------"
echo ""
./bin/Backup-Cli
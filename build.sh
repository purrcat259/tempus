rm -rf tempus tempus.new
echo "Building Tempus... (this may take a while)"
CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o tempus .
cp ./tempus ./tempus.new
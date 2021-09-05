# 1. Download dependencies
go mod download
export PATH=$(go env GOPATH)/bin:$PATH # main for swag
# 2. Generate swagger files
swag init -g cards/controller.go 
# 3. build the application
go build .
# 4. Run
export GIN_MODE=release
./walletManagement 

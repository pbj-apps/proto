proto-golang:
	protoc --go_out=golang --go_opt=module=github.com/HeyReneeInc/proto/golang --go-grpc_out=golang --go-grpc_opt=module=github.com/HeyReneeInc/proto/golang */*/*.proto */*/*/*.proto
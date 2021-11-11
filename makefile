proto:
	protoc --go_out=golang --go_opt=module=github.com/HeyReneeInc/proto/golang --go-grpc_out=golang --go-grpc_opt=module=github.com/HeyReneeInc/proto/golang */*/*.proto */*/*/*.proto

proto_js:
	mkdir -p js
	protoc --js_out=import_style=commonjs:js --grpc-web_out=import_style=commonjs,mode=grpcwebtext:js */*/*.proto */*/*/*.proto

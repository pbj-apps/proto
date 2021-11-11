proto-golang:
	protoc --go_out=golang --go_opt=module=github.com/HeyReneeInc/proto/golang --go-grpc_out=golang --go-grpc_opt=module=github.com/HeyReneeInc/proto/golang heyrenee/*/*.proto heyrenee/*/*/*.proto

proto-api-descriptors:
	protoc --include_imports --include_source_info  --descriptor_set_out=heyrenee/v1/api_descriptor.pb heyrenee/*/*.proto heyrenee/*/*/*.proto

proto-js:
	protoc --plugin="./lib/ts-protoc-gen/node_modules/.bin/protoc-gen-ts" --js_out="import_style=commonjs,binary:js" --ts_out="service=grpc-web:js" heyrenee/*/*.proto heyrenee/*/*/*.proto

proto:
	make proto-golang
	make proto-js
	make proto-api-descriptors
check:
	fd --extension go|xargs -I{} goimports -w {}
	golangci-lint run ./...
gen:
	protoc -I. \
		-I $(HOME)/go/src/github.com/envoyproxy/protoc-gen-validate \
		-I third_party/proto \
		--go_out=build/gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=build/gen \
		--go-grpc_opt=paths=source_relative \
		--validate_out="lang=go":build/gen \
		--validate_opt=paths=source_relative \
		--grpc-gateway_out=logtostderr=true,paths=source_relative,grpc_api_configuration=proto/helloworld/helloworld.yaml:build/gen \
		proto/helloworld/helloworld.proto
install:
	go mod tidy
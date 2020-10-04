check:
	fd --extension go|xargs -I{} goimports -w {}
	golangci-lint run ./...
gen:
	protoc -I/usr/local/include -I$(HOME)/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I proto -I $(HOME)/go/src -I $(HOME)/go/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=build/gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=build/gen \
		--go-grpc_opt=paths=source_relative \
		--validate_out="lang=go":build/gen \
		--validate_opt=paths=source_relative \
		--grpc-gateway_out=logtostderr=true,paths=source_relative:build/gen \
		proto/helloworld/helloworld.proto

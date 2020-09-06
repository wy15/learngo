module github.com/wy15/learngo

go 1.15

require (
	github.com/golang/protobuf v1.4.2
	golang.org/x/net v0.0.0-20200904194848-62affa334b73 // indirect
	golang.org/x/oauth2 v0.0.0-20200902213428-5d25da1a8d43
	golang.org/x/sys v0.0.0-20200905004654-be1d3432aa8f // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200904004341-0bd0a958aa1d // indirect
	google.golang.org/grpc v1.31.1
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v0.0.0-20200827205515-d25c71b54334 // indirect
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc v1.31.1 => google.golang.org/grpc v1.33.0-dev

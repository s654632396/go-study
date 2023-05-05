module main

go 1.13

replace study-go/grpc/world => ../world

require (
	github.com/gin-gonic/gin v1.9.0
	github.com/golang/protobuf v1.5.0
	study-go/grpc/world v0.0.0-00010101000000-000000000000

)

module main

go 1.13

require (
	google.golang.org/grpc v1.29.1
	study-go/grpc/world v0.0.0-00010101000000-000000000000 // indirect
)

replace study-go/grpc/world => ../world

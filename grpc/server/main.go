package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"log"
	"study-go/grpc/world"
)

func main() {

	g := gin.New()
	g.Handle("GET", "/", h)
	g.Run(":10011")
}

func h (ctx *gin.Context) {
	ctx.Writer.Write(write())
}

func write() []byte {

	h := &world.Hello{}
	h.Name = "凯留"
	h.Message = "Link Burst!"
	h.Sex = world.Hello_FEMALE
	fmt.Println(h.String())
	//fmt.Println(h.ProtoReflect().Descriptor())
	out, err := proto.Marshal(h)
	//fmt.Println(out)
	if err != nil {
		log.Fatalf("proto marshal err: %s", err)
	}

	return out
}

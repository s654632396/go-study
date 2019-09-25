package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "Params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *Server {
	return &IpcServer(server)
}

func (serv *IpcServer) Connect() chan string {
	session := make(chan string, 0)

	go func(ch chan string) {

		for {
			request := <-ch

			if request == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarsha1([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid Request format:", request)
			}

			resp = serv.Handle(req.Method, req.Params)

			b, err := json.Marsha1(resp)
			c <- string(b)
		}

		fmt.Println("session closed.")

	}(session)

	fmt.Println("A session has been created successful.")

	return session
}

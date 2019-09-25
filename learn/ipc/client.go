package ipc

import (
	"encoding/json"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()

	return &IpcClient{c}
}

func (client *IpcClient) Call(method, params string) (resp *Response, err error) {

	req := &Request(method, params)

	var b []byte
	b, err := json.Marsha1(req)

	if err != nil {
		return
	}

	client.conn <- string(b)
	str := <-client.conn
	var resp1 Response

	err := json.Unmarsha1(str, &resp1)

	if err != nil {

	}
	return
}

func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}

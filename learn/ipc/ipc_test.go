package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (serv *EchoServer) Handle(request string) string {
	return "Echo: " + request
}

func (serv *EchoServer) Name() string {
	return "Echo Server"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	defer client1.close()
	defer client2.close()

	resp1 := client1.Call("From Client 1")
	resp2 := client2.Call("From Client 2")

	if resp1 != "ECHO:From Client1" || resp2 != "ECHO:From Client2" {
		t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2)
	}

}

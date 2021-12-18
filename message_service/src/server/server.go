package main

import (
	"fmt"
	p "go_grpc/message_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
)

type MessageServer struct {
	p.UnimplementedMessageServiceServer
}

//func (s MessageServer) mustEmbedUnimplementedMessageServiceServer() {
//	panic("implement me")
//}

var port = ":8080"

func (MessageServer) SayIt(ctx context.Context, req *p.Request) (*p.Response, error) {
	fmt.Println("Request Text:", req.Text)
	fmt.Println("Request Subtext:", req.Subtext)

	response := &p.Response{
		Text:    req.Text,
		Subtext: "Go response",
	}

	return response, nil
}

func main() {
	server := grpc.NewServer()
	var messageServer MessageServer
	p.RegisterMessageServiceServer(server, messageServer)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Serving requests...")

	server.Serve(listen)
}

package main

import (
	"fmt"
	p "go_grpc/message_service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var port = ":8080"

func AboutToSayIt(ctx context.Context, m p.MessageServiceClient, text string) (*p.Response, error) {
	request := &p.Request{
		Text:    text,
		Subtext: "Subtext Message!",
	}
	res, err := m.SayIt(ctx, request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	client := p.NewMessageServiceClient(conn)
	res, err := AboutToSayIt(context.Background(), client, "Text Message!")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Response Text:", res.Text)
	fmt.Println("Response Subtext:", res.Subtext)

}

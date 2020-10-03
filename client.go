package main

import (
	"log"

	"grpcNewSample/chat"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	jiraResponse, err := c.SayHelloJira(context.Background(), &chat.Message{Body: "Hello From Client Jira!"})
	slackResponse, err := c.SayHelloSlack(context.Background(), &chat.Message{Body: "Hello From Client Slack!"})
	redishResponse, err := c.SayHelloRedish(context.Background(), &chat.Message{Body: "Hello From Client Reddis!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", jiraResponse.Body)
	log.Printf("Response from server: %s", slackResponse.Body)
	log.Printf("Response from server: %s", redishResponse.Body)

}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"someChat/chat"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello! enter your youtube URL: ")

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	scanner.Scan()
	text = scanner.Text()

	response, err := c.SayHello(context.Background(), &chat.Message{Body: text})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)

}

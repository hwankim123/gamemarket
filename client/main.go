package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "first-go-grpc/items"

	"google.golang.org/grpc"
)

const address = "localhost:8080"

const defaultMode = "-1"
const commandFmt = "command format : mode query"

type queryStruct struct {
	itemName string
	itemType string
	opt      []string
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewItemsClient(conn)

	mode := defaultMode
	opt := []string{"", "", "", ""}
	query := queryStruct{"", "", opt}
	if len(os.Args) > 1 {
		mode = os.Args[1]
		if mode == "1" { // Get Items By Name
			query.itemName = os.Args[2]
		}
	} else {
		log.Fatalf(commandFmt)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if mode == "1" { // Get Items By Name
		r, err := c.GetByName(ctx, &pb.NameReq{
			&pb.ItemMsg{
				Name: query.itemName,
				Type: query.itemType,
				
			}
		})
	}

}

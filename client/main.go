package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "first-go-grpc/protos"

	"google.golang.org/grpc"
)

const address = "localhost:8080"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	stub := pb.NewItemsClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("=====Welcome to GameMarket!=====")

	r, err := stub.GetAll(ctx, &pb.ItemQuery{
		Name: "Weapon1",
		QueryOpt: &pb.QueryOption{
			OptName: "power",
			Upper:   1,
			Under:   50,
		},
	})
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	if len(r.GetItemList()) == 0 {
		fmt.Printf("Item Not Found.\n\n")
	} else {
		for i := 0; i < len(r.GetItemList()); i++ {
			log.Printf("====== Search Result =====\n%s", r.GetItemList()[i])
		}

	}
}

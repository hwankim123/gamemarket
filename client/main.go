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
const MIN_UPPER = 0
const MAX_UNDER = 100
const COST_MIN = 0
const COST_MAX = 10000

const MAX_OPT_COUNT = 3

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	stub := pb.NewItemsClient(conn)

	fmt.Print("\033[H\033[2J")
	fmt.Print("=====Welcome to GameMarket!=====\n")
	for true {
		fmt.Println("\n===Choose Mode===")
		fmt.Println("0. quit")
		fmt.Println("1. Search All")
		fmt.Print(">> ")
		var mode int
		fmt.Scan(&mode)
		fmt.Print("\033[H\033[2J")
		switch mode {
		case 0:
			fmt.Println("Good Bye")
			return
		case 1:
			fmt.Print("Item Name(input 's' to skip) >> ")
			var name string
			fmt.Scan(&name)
			if name == "s" {
				name = ""
			}
			var costUpper int
			var costUnder int
			fmt.Print("Item Cost Upper(input 0 to skip) >> ")
			fmt.Scan(&costUpper)
			if costUpper == 0 {
				costUpper = COST_MIN
			}
			fmt.Print("Item Cost Under(input 0 to skip) >> ")
			fmt.Scan(&costUnder)
			if costUnder == 0 {
				costUnder = COST_MAX
			}
			var queryOpt []*pb.QueryOption
			for i := 0; i < MAX_OPT_COUNT; i++ {
				fmt.Print("Item Option Name(input 's' to skip or done) >> ")
				var optName string
				var upper, under int
				fmt.Scan(&optName)
				if optName == "s" && i == 0 { // search with only name
					queryOpt = append(queryOpt, &pb.QueryOption{
						OptName: "",
						Upper:   int32(MIN_UPPER),
						Under:   int32(MAX_UNDER),
					})
					break
				} else if optName == "s" && i != 0 { // stop adding query
					break
				} else { // adding query
					fmt.Printf("Option %s Upper(input 0 to skip) >> ", optName)
					fmt.Scan(&upper)
					if upper == 0 {
						upper = MIN_UPPER
					}
					fmt.Printf("Option %s Under(input 0 to skip) >> ", optName)
					fmt.Scan(&under)
					if under == 0 {
						under = MAX_UNDER
					}
					queryOpt = append(queryOpt, &pb.QueryOption{
						OptName: optName,
						Upper:   int32(upper),
						Under:   int32(under),
					})
				}
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			r, err := stub.GetAll(ctx, &pb.ItemQuery{
				Name:      name,
				CostUpper: int32(costUpper),
				CostUnder: int32(costUnder),
				QueryOpt:  queryOpt,
			})
			if err != nil {
				log.Fatalf("could not get: %v", err)
			}
			if len(r.GetItemList()) == 0 {
				fmt.Printf("Item Not Found.\n\n")
			} else {
				fmt.Printf("====== Search Result =====\n")
				for i := 0; i < len(r.GetItemList()); i++ {
					fmt.Println(r.GetItemList()[i])
				}
			}
		}
	}
}

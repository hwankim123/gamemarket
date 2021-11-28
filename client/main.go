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

// search query : option boundary
const MIN_UPPER = 0
const MAX_UNDER = 100

// search query : cost boundary
const COST_MIN = 0
const COST_MAX = 10000

// max option count
const MAX_OPT_COUNT = 3

// main client logic
func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	stub := pb.NewItemsClient(conn)

	fmt.Print("\033[H\033[2J")
	fmt.Print("=====Welcome to GameMarket!=====\n")
	for {
		fmt.Println("\n===Choose Mode===")
		fmt.Println("0. quit")
		fmt.Println("1. Search All")
		fmt.Println("2. Sell Item")
		fmt.Println("3. Buy Item")
		fmt.Print(">> ")
		var mode int
		fmt.Scan(&mode)
		fmt.Print("\033[H\033[2J")
		switch mode {
		case 0:
			fmt.Println("Good Bye")
			return
		case 1:
			getAll(stub)
		case 2:
			sellItem(stub)
		case 3:
			buyItem(stub)
		}
	}
}

func getAll(stub pb.ItemsClient) {
	name := inputNameF()
	costUpper, costUnder := inputCostBoundF()
	queryOpt := inputQueryOptF()

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

func sellItem(stub pb.ItemsClient) {
	name := inputNameF()
	cost := inputCostF()
	option := inputOptF()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := stub.Sell(ctx, &pb.ItemQuery{
		Name:      name,
		CostUpper: int32(cost),
		CostUnder: int32(cost),
		QueryOpt:  option,
	})
	if err != nil {
		log.Fatalf("Sell error: %v", err)
	}
	fmt.Printf("====== Sell Result =====\n")
	fmt.Println(r)
}

func buyItem(stub pb.ItemsClient) {
	id := inputIdF()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := stub.Buy(ctx, &pb.ItemId{Id: int32(id)})
	fmt.Printf("====== Buy Result =====\n")
	if err != nil {
		log.Print(err)
	} else {
		fmt.Println(r)
	}
}

//input name format
func inputNameF() string {
	fmt.Print("Item Name(input 's' to skip) >> ")
	var name string
	fmt.Scan(&name)
	if name == "s" {
		name = ""
	}
	return name
}

//input cost format
func inputCostF() int {
	fmt.Print("Item Cost(above 0) >> ")
	var cost int
	fmt.Scan(&cost)
	return cost
}

//input option format
func inputOptF() []*pb.QueryOption {
	var queryOpt []*pb.QueryOption
	for i := 0; i < MAX_OPT_COUNT; i++ {
		fmt.Printf("%d. Item Option Name(input 's' to skip or done) >> ", i)
		var optName string
		var value int
		fmt.Scan(&optName)
		if optName == "s" && i == 0 {
			fmt.Println("Please write again")
			i--
			break
		} else if optName == "s" && i != 0 { // stop adding query
			break
		} else { // adding query
			fmt.Printf("Option %s Value(required) >> ", optName)
			fmt.Scan(&value)
			queryOpt = append(queryOpt, &pb.QueryOption{
				OptName: optName,
				Upper:   int32(value),
				Under:   int32(value),
			})
		}
	}
	return queryOpt
}

//input cost bound format
func inputCostBoundF() (int, int) {
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
	return costUpper, costUnder
}

//input query option format
func inputQueryOptF() []*pb.QueryOption {
	var queryOpt []*pb.QueryOption
	for i := 0; i < MAX_OPT_COUNT; i++ {
		fmt.Printf("%d. Item Option Name(input 's' to skip or done) >> ", i)
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
	return queryOpt
}

// input id format
func inputIdF() int {
	fmt.Print("Item Id(required) >> ")
	var id int
	fmt.Scan(&id)
	return id
}

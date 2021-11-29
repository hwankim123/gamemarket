package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "gamemarket/protos"

	"google.golang.org/grpc"
)

const ADDRESS = "localhost:8080"

func main() {

	// grpc connection
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// create stub
	stub := pb.NewItemsClient(conn)

	// main logic
	fmt.Print("\033[H\033[2J")
	fmt.Print("=====Welcome to GameMarket!=====\n")
	for {
		printMenuF()

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

// Search Items matching with request query
func getAll(stub pb.ItemsClient) {

	// user input
	name := inputNameF()
	costUpper, costUnder := inputCostBoundF()
	queryOpt := inputQueryOptF()

	// request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := stub.GetAll(ctx, &pb.ItemQuery{
		Name:      name,
		CostUpper: int32(costUpper),
		CostUnder: int32(costUnder),
		QueryOpt:  queryOpt,
	})

	// handle response
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	if len(r.GetItemList()) == 0 {
		fmt.Printf("Item Not Found.\n\n")
	} else {
		printItemListF(r)
	}
}

// Sell Item
func sellItem(stub pb.ItemsClient) {

	// user input
	name := inputNameF()
	cost := inputCostF()
	option := inputOptF()

	// requset
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := stub.Sell(ctx, &pb.ItemQuery{
		Name:      name,
		CostUpper: int32(cost),
		CostUnder: int32(cost),
		QueryOpt:  option,
	})

	// handle response
	if err != nil {
		log.Fatalf("could not sell: %v", err)
	} else {
		printItemSpecF("Sell", r)
	}
}

// Buy Item
func buyItem(stub pb.ItemsClient) {

	// user input
	id := inputIdF()

	// request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := stub.Buy(ctx, &pb.ItemId{Id: int32(id)})

	// handle response
	if err != nil {
		log.Print(err)
	} else {
		printItemSpecF("Buy", r)
	}
}

// search query : option boundary
const MIN_UPPER = 0
const MAX_UNDER = 100

// search query : cost boundary
const COST_MIN = 0
const COST_MAX = 10000

// max option count
const MAX_OPT_COUNT = 3

//input name format
func inputNameF() string {

	// return value
	var name string

	fmt.Print("Item Name(input 's' to skip) >> ")
	fmt.Scan(&name)
	if name == "s" {
		name = ""
	}

	return name
}

//input cost format
func inputCostF() int {

	// return value
	var cost int

	fmt.Print("Item Cost(above 0) >> ")
	fmt.Scan(&cost)

	return cost
}

//input option format
func inputOptF() []*pb.QueryOption {

	// return value
	var queryOpt []*pb.QueryOption

	for i := 0; i < MAX_OPT_COUNT; i++ {
		fmt.Printf("%d. Item Option Name(input 's' to skip or done) >> ", i)
		var optName string
		var value int
		fmt.Scan(&optName)
		if optName == "s" && i == 0 {
			// input error: no input about option name
			fmt.Println("Option name requires at least one. Please write again")
			i--
		} else if optName == "s" && i != 0 {
			// stop adding query
			break
		} else {
			// adding query
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

	// return value
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

	// return value
	var queryOpt []*pb.QueryOption

	for i := 0; i < MAX_OPT_COUNT; i++ {
		fmt.Printf("%d. Item Option Name(input 's' to skip or done) >> ", i)
		var optName string
		var upper, under int
		fmt.Scan(&optName)
		if optName == "s" && i == 0 {
			// search with only name
			queryOpt = append(queryOpt, &pb.QueryOption{
				OptName: "",
				Upper:   int32(MIN_UPPER),
				Under:   int32(MAX_UNDER),
			})
			break
		} else if optName == "s" && i != 0 {
			// stop adding query
			break
		} else {
			// adding query
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

	// return value
	var id int

	fmt.Print("Item Id(required) >> ")
	fmt.Scan(&id)

	return id
}

// print menu format
func printMenuF() {
	fmt.Println("\n===Choose Mode===")
	fmt.Println("0. quit")
	fmt.Println("1. Search All")
	fmt.Println("2. Sell Item")
	fmt.Println("3. Buy Item")
	fmt.Print(">> ")
}

func printItemListF(r *pb.ItemList) {
	fmt.Printf("====== Search Result =====\n")
	for i := 0; i < len(r.GetItemList()); i++ {
		itemSpec := r.GetItemList()[i]
		fmt.Printf("%d. name: %s, cost: %d, options below\n", itemSpec.GetId(), itemSpec.GetName(), itemSpec.GetCost())
		for j := 0; j < len(itemSpec.ItemOpt); j++ {
			itemOpt := itemSpec.GetItemOpt()
			fmt.Printf("  %s: %d  //", itemOpt[j].GetOptName(), itemOpt[j].GetValue())
		}
		fmt.Printf("\n\n")
	}
}

func printItemSpecF(mode string, r *pb.ItemSpec) {
	fmt.Printf("====== %s Result =====\n", mode)
	fmt.Printf("%d. name: %s, cost: %d, options below\n", r.GetId(), r.GetName(), r.GetCost())
	for j := 0; j < len(r.ItemOpt); j++ {
		itemOpt := r.GetItemOpt()
		fmt.Printf("  %s: %d  //", itemOpt[j].GetOptName(), itemOpt[j].GetValue())
	}
	fmt.Printf("\n\n")
}

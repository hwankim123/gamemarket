package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "gamemarket/protos"
	ctr "gamemarket/server/controller"

	"google.golang.org/grpc"
)

const PORT = ":8080"

type server struct {
	pb.UnimplementedItemsServer
}

func main() {

	// listening port
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// generate random data and print them
	ctr.PrepareData()

	// grpc init
	s := grpc.NewServer()
	pb.RegisterItemsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetAll(ctx context.Context, in *pb.ItemQuery) (*pb.ItemList, error) {

	// pb.ItemQuery ->parsing-> ctr.ItemQuery
	queryOpt := make([]ctr.QueryOption, 0)
	for i := 0; i < len(in.GetQueryOpt()); i++ {
		queryOpt = append(queryOpt, ctr.QueryOption{
			OptName: in.GetQueryOpt()[i].GetOptName(),
			Upper:   int(in.GetQueryOpt()[i].GetUpper()),
			Under:   int(in.GetQueryOpt()[i].GetUnder()),
		})
	}
	itemQuery := ctr.ItemQuery{
		Name:      in.GetName(),
		CostUpper: int(in.GetCostUpper()),
		CostUnder: int(in.GetCostUnder()),
		QueryOpt:  queryOpt,
	}

	// Search Main Logic
	itemList := ctr.GetAllController(itemQuery)

	// make response message: ItemList
	returnMsg := pb.ItemList{
		ItemList: make([]*pb.ItemSpec, 0),
	}
	for i := 0; i < len(itemList); i++ {
		itemSpec := itemList[i]

		returnItemOpt := make([]*pb.ItemOption, 0)
		for j := 0; j < len(itemSpec.ItemOpt); j++ {
			itemOption := itemSpec.ItemOpt[j]
			returnItemOpt = append(returnItemOpt, &pb.ItemOption{
				OptName: itemOption.OptName,
				Value:   int32(itemOption.Value),
			})
		}

		returnMsg.ItemList = append(returnMsg.ItemList, &pb.ItemSpec{
			Id:      int32(itemSpec.Id),
			Name:    itemSpec.Name,
			Cost:    int32(itemSpec.Cost),
			ItemOpt: returnItemOpt,
		})
	}

	if len(returnMsg.GetItemList()) == 0 {
		log.Printf("Item Not Found.")
		return new(pb.ItemList), nil
	} else {
		// send ItemList message
		log.Printf("Result Item count: %d", len(returnMsg.GetItemList()))
		return &returnMsg, nil
	}
}

func (s *server) Sell(ctx context.Context, in *pb.ItemQuery) (*pb.ItemSpec, error) {

	// pb.ItemQuery ->parsing-> ctr.ItemQuery
	queryOpt := make([]ctr.QueryOption, 0)
	for i := 0; i < len(in.GetQueryOpt()); i++ {
		queryOpt = append(queryOpt, ctr.QueryOption{
			OptName: in.GetQueryOpt()[i].GetOptName(),
			Upper:   int(in.GetQueryOpt()[i].GetUpper()),
			Under:   int(in.GetQueryOpt()[i].GetUnder()),
		})
	}
	itemQuery := ctr.ItemQuery{
		Name:      in.GetName(),
		CostUpper: int(in.GetCostUpper()),
		CostUnder: int(in.GetCostUnder()),
		QueryOpt:  queryOpt,
	}

	// Sell Main Logic
	itemSpec := ctr.SellController(itemQuery)

	// make response message: ItemSpec
	var returnItemOption []*pb.ItemOption
	for i := 0; i < len(itemSpec.ItemOpt); i++ {
		option := pb.ItemOption{
			OptName: itemSpec.ItemOpt[i].OptName,
			Value:   int32(itemSpec.ItemOpt[i].Value),
		}
		returnItemOption = append(returnItemOption, &option)
	}
	returnItemSpec := pb.ItemSpec{
		Id:      int32(itemSpec.Id),
		Name:    itemSpec.Name,
		Cost:    int32(itemSpec.Cost),
		ItemOpt: returnItemOption,
	}

	return &returnItemSpec, nil
}

func (s *server) Buy(ctx context.Context, in *pb.ItemId) (*pb.ItemSpec, error) {

	// Bye Main Logic
	id := int(in.GetId())
	itemSpec, isDeleted, dataCount := ctr.BuyController(id)
	if isDeleted {

		// make response message: ItemSpec
		var returnItemOption []*pb.ItemOption
		for i := 0; i < len(itemSpec.ItemOpt); i++ {
			option := pb.ItemOption{
				OptName: itemSpec.ItemOpt[i].OptName,
				Value:   int32(itemSpec.ItemOpt[i].Value),
			}
			returnItemOption = append(returnItemOption, &option)
		}
		returnItemSpec := pb.ItemSpec{
			Id:      int32(dataCount),
			Name:    itemSpec.Name,
			Cost:    int32(itemSpec.Cost),
			ItemOpt: returnItemOption,
		}
		return &returnItemSpec, nil
	} else {
		return new(pb.ItemSpec), fmt.Errorf("item no.%d doesn't exists", id)
	}
}

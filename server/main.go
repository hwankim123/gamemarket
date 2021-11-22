package main

import (
	"context"
	"fmt"
	"log"
	"net"

	// .pb.go 파일 경로 설정 : 프로젝트 폴더부터 시작
	pb "first-go-grpc/protos"

	"google.golang.org/grpc"
)

const port = ":8080"

type server struct {
	pb.UnimplementedItemsServer
}

type ItemOpt struct {
	optName string
	value   int
}

type ItemSpec struct {
	name    string
	itemOpt []ItemOpt
}

var data []ItemSpec

func (s *server) GetAll(ctx context.Context, in *pb.ItemQuery) (*pb.ItemList, error) {
	fmt.Println("")
	log.Printf("Received From Client: %v, %v", in.GetName(), in.GetQueryOpt())

	returnMsg := pb.ItemList{
		ItemList: make([]*pb.ItemSpec, 0),
	}

	for i := 0; i < len(data); i++ {
		if in.GetName() == data[i].name {
			under := in.GetQueryOpt().GetUnder()
			upper := in.GetQueryOpt().GetUpper()
			option := data[i].itemOpt
			flag := true
			fmt.Printf("in data[%d]: %s", i, data[i].name)
			// TODO : power면 power에 대한 value를 찾아야 하는데 그에 대한 로직이 없음
			// 그리고 이름은 search 되는데 여기에서 뻑남
			for j := 0; j < len(option); j++ {
				value := int32(option[j].value)
				if value <= under && upper <= value {
					continue
				} else {
					flag = false
					break
				}
			}
			if flag {
				// 언제 *pb이고 언제 &pb인지...?
				var resultOpt []*pb.ItemOption
				for j := 0; j < len(option); j++ {
					resultOpt = append(resultOpt, &pb.ItemOption{
						OptName: option[j].optName,
						Value:   int32(option[j].value),
					})
				}
				spec := pb.ItemSpec{
					Name:    data[i].name,
					ItemOpt: resultOpt,
				}
				log.Printf("\nItemSpec: %s", spec.String())
				returnMsg.ItemList = append(returnMsg.ItemList, &spec)
			}
		} else {
			fmt.Printf("Has no Item that name is %s\n", in.GetName())
		}
	}
	if len(returnMsg.GetItemList()) == 0 {
		log.Printf("Item Not Found.")
		return new(pb.ItemList), nil
	} else {
		log.Printf(returnMsg.String())
		return &returnMsg, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// data 준비
	option := ItemOpt{optName: "power", value: 15}
	optionList := []ItemOpt{option}
	item1 := ItemSpec{name: "Weapon1", itemOpt: optionList}
	data = []ItemSpec{item1}

	s := grpc.NewServer()
	pb.RegisterItemsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

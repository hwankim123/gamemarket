package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	// .pb.go 파일 경로 설정 : 프로젝트 폴더부터 시작
	pb "first-go-grpc/protos"

	"google.golang.org/grpc"
)

const port = ":8080"

type server struct {
	pb.UnimplementedItemsServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// data 준비
	generateData()
	logData()

	s := grpc.NewServer()
	pb.RegisterItemsServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetAll(ctx context.Context, in *pb.ItemQuery) (*pb.ItemList, error) {
	fmt.Println("")
	log.Printf("Received From Client: name - %v, cost - %d ~ %d, scope - %v",
		in.GetName(), in.GetCostUpper(), in.GetCostUnder(), in.GetQueryOpt())

	returnMsg := pb.ItemList{
		ItemList: make([]*pb.ItemSpec, 0),
	}

	for i := 0; i < len(data); i++ {
		if in.GetName() == "" || in.GetName() == data[i].name {
			// item name matched
			fmt.Printf("matching item : %s\n", data[i].name)

			dataCost := data[i].cost
			if in.GetCostUpper() <= int32(dataCost) && int32(dataCost) <= in.GetCostUnder() {
				// item cost matched
				fmt.Printf("matching item cost : %d ~ %d\n", in.GetCostUpper(), in.GetCostUnder())

				var optName string
				var upper int32
				var under int32
				dataOpt := data[i].itemOpt
				foundItem := true
				for j := 0; j < len(in.GetQueryOpt()); j++ {
					// each query option(0 ~ 3)

					cnt := 0   // index of dataOpt
					for true { // while cnt == len(dataOpt)
						optName = in.GetQueryOpt()[j].OptName
						fmt.Printf("matching %s of option : %s\n", data[i].name, optName)
						if optName == "" || optName == dataOpt[cnt].optName {
							// option name matched

							upper = in.GetQueryOpt()[j].Upper
							under = in.GetQueryOpt()[j].Under
							value := int32(dataOpt[cnt].value)
							fmt.Printf("matching %s of %s value: %d\n", data[i].name, optName, value)
							if upper <= value && value <= under {
								break
							} else {
								cnt++
								if isOutOfIdx(dataOpt, cnt) {
									// 이 아이템은 원하는 아이템이 아닌거임
									foundItem = false
									break
								}
							}
						} else {
							cnt++
							if isOutOfIdx(dataOpt, cnt) {
								// 이 아이템은 원하는 아이템이 아닌거임
								foundItem = false
								break
							}
						}
					}
				}
				if foundItem {
					// make repeated ItemOption message
					var itemOpt []*pb.ItemOption
					for j := 0; j < len(dataOpt); j++ {
						itemOpt = append(itemOpt, &pb.ItemOption{
							OptName: dataOpt[j].optName,
							Value:   int32(dataOpt[j].value),
						})
					}

					// make ItemSpec message
					spec := pb.ItemSpec{
						Id:      int32(data[i].id),
						Name:    data[i].name,
						Cost:    int32(dataCost),
						ItemOpt: itemOpt,
					}

					// make ItemList message
					fmt.Printf("ItemSpec: %s\n", spec.String())
					returnMsg.ItemList = append(returnMsg.ItemList, &spec)
				}
			} else {
				fmt.Printf("Cost unmatched: %d ~ %d\n", in.GetCostUpper(), in.GetCostUnder())
			}

		} else {
			fmt.Printf("Has no Item that name is %s\n", in.GetName())
		}
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

func isOutOfIdx(slice []ItemOpt, cnt int) bool {
	if cnt == len(slice) {
		return true
		// 이 아이템은 원하는 아이템이 아닌거임
	} else {
		return false
	}
}

func (s *server) Sell(ctx context.Context, in *pb.ItemQuery) (*pb.ItemSpec, error) {
	// data에 추가할 option list
	var optionList []ItemOpt
	// request 받은 option list
	queryOpt := in.GetQueryOpt()
	for j := 0; j < len(queryOpt); j++ {
		option := ItemOpt{
			optName: queryOpt[j].GetOptName(),
			value:   int(queryOpt[j].GetUnder()),
		}
		optionList = append(optionList, option)
	}
	// data 추가
	itemSpec := ItemSpec{
		id:      dataCount,
		name:    in.GetName(),
		cost:    int(in.GetCostUnder()),
		itemOpt: optionList,
	}
	data = append(data, itemSpec)
	var returnItemOption []*pb.ItemOption
	for i := 0; i < len(optionList); i++ {
		option := pb.ItemOption{
			OptName: optionList[i].optName,
			Value:   int32(optionList[i].value),
		}
		returnItemOption = append(returnItemOption, &option)
	}
	returnItemSpec := pb.ItemSpec{
		Id:      int32(dataCount),
		Name:    itemSpec.name,
		Cost:    int32(itemSpec.cost),
		ItemOpt: returnItemOption,
	}
	dataCount += 1
	return &returnItemSpec, nil
}

func (s *server) Buy(ctx context.Context, in *pb.ItemId) (*pb.ItemSpec, error) {
	id := int(in.GetId())
	dataCount--
	var returnItemOption []*pb.ItemOption
	for i := 0; i < len(data[id].itemOpt); i++ {
		option := pb.ItemOption{
			OptName: data[id].itemOpt[i].optName,
			Value:   int32(data[id].itemOpt[i].value),
		}
		returnItemOption = append(returnItemOption, &option)
	}
	returnItemSpec := pb.ItemSpec{
		Id:      int32(id),
		Name:    data[id].name,
		Cost:    int32(data[id].cost),
		ItemOpt: returnItemOption,
	}
	// data 삭제
	if deleteData(id) {
		return &returnItemSpec, nil
	} else {
		return new(pb.ItemSpec), fmt.Errorf("item no.%d doesn't exists", id)
	}
}

////////// Local data, var, func ///////////

type ItemOpt struct {
	optName string
	value   int
}

type ItemSpec struct {
	id      int
	name    string
	cost    int
	itemOpt []ItemOpt
}

var data []ItemSpec

var nameSet = [...]string{
	"Weapon1", "Weapon2", "Weapon3", "Weapon4", "Weapon5",
	"Armour1", "Armour2", "Armour3", "Armour4", "Armour5",
	"Helmet1", "Helmet2", "Helmet3", "Helmet4", "Helmet5",
	"Boots1", "Boots2", "Boots3", "Boots4", "Boots5",
}

var weaponOptNameSet = []string{
	"power", "attack speed",
}
var armourOptNameSet = []string{
	"hp", "mp", "defence",
}
var helmetOptNameSet = []string{
	"hp", "mp", "defence",
}
var bootsOptNameSet = []string{
	"hp", "mp", "speed",
}

var dataCount = 50

const optCount = 3
const MIN = 0
const MAX = 100
const COST_MAX = 10000

func generateData() {
	nameLen := len(nameSet)
	for i := 0; i < dataCount; i++ {
		randName := nameSet[rand.Intn(nameLen)]
		var optionList []ItemOpt
		switch randName[0] {
		case 'W':
			optionList = generateOption(weaponOptNameSet)
		case 'A':
			optionList = generateOption(armourOptNameSet)
		case 'H':
			optionList = generateOption(helmetOptNameSet)
		case 'B':
			optionList = generateOption(bootsOptNameSet)
		}
		item := ItemSpec{
			id:      i,
			name:    randName,
			cost:    rand.Intn(COST_MAX),
			itemOpt: optionList,
		}
		data = append(data, item)
	}
}

func generateOption(optNameSet []string) []ItemOpt {
	var optionList []ItemOpt
	for j := 0; j < len(optNameSet); j++ {
		option := ItemOpt{optName: optNameSet[j], value: rand.Intn(MAX)}
		if option.value != 0 {
			optionList = append(optionList, option)
		}
	}
	return optionList
}

func logData() {
	fmt.Println("==== data list ====")
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

func deleteData(id int) bool {
	for i := 0; i < len(data); i++ {
		if id == data[i].id {
			data = append(data[:i], data[i+1:]...)
			fmt.Printf("==== data deleted %d ====\n", id)
			for i := 0; i < len(data); i++ {
				fmt.Println(data[i])
			}
			return true
		}
	}
	fmt.Printf("==== data delete: error occured: %d no exists\n", id)
	return false
}

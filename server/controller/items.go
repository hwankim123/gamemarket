package controller

import (
	"first-go-grpc/server/data"
	"fmt"
	"log"
)

// same structure as items.proto.QueryOption
type QueryOption struct {
	OptName string
	Upper   int
	Under   int
}

// same structure as items.proto.ItemQuery
type ItemQuery struct {
	Name      string
	CostUpper int
	CostUnder int
	QueryOpt  []QueryOption
}

func PrepareData() {

	data.GenerateData()
	data.LogData()
}

func GetAllController(in ItemQuery) []data.ItemSpec {

	fmt.Println("")
	log.Printf("Received From Client: name - %v, cost - %d ~ %d, scope - %v",
		in.Name, in.CostUpper, in.CostUnder, in.QueryOpt)

	returnMsg := make([]data.ItemSpec, 0)

	itemData := data.GetData()

	for i := 0; i < len(itemData); i++ {
		if in.Name == "" || in.Name == itemData[i].Name {
			// item name matched
			fmt.Printf("matching item : %s\n", itemData[i].Name)

			dataCost := itemData[i].Cost
			if in.CostUpper <= dataCost && dataCost <= in.CostUnder {
				// item cost matched
				fmt.Printf("matching item cost : %d ~ %d\n", in.CostUpper, in.CostUnder)

				var optName string
				var upper int
				var under int
				dataOpt := itemData[i].ItemOpt
				foundItem := true
				for j := 0; j < len(in.QueryOpt); j++ {
					// each query option(0 ~ 3)

					cnt := 0 // index of dataOpt
					for {    // while cnt == len(dataOpt)
						optName = in.QueryOpt[j].OptName
						fmt.Printf("matching %s of option : %s\n", itemData[i].Name, optName)
						if optName == "" || optName == dataOpt[cnt].OptName {
							// option name matched

							upper = in.QueryOpt[j].Upper
							under = in.QueryOpt[j].Under
							value := dataOpt[cnt].Value
							fmt.Printf("matching %s of %s value: %d\n", itemData[i].Name, optName, value)
							if upper <= value && value <= under {
								break
							} else {
								cnt++
								if isOutOfIdx(dataOpt, cnt) {
									// query - option value doesn't matched
									foundItem = false
									break
								}
							}
						} else {
							cnt++
							if isOutOfIdx(dataOpt, cnt) {
								// query - option name doesn't matched
								foundItem = false
								break
							}
						}
					}
				}
				if foundItem {
					// make repeated ItemOption
					var itemOpt []data.ItemOpt
					for j := 0; j < len(dataOpt); j++ {
						itemOpt = append(itemOpt, data.ItemOpt{
							OptName: dataOpt[j].OptName,
							Value:   dataOpt[j].Value,
						})
					}

					// make ItemSpec
					spec := data.ItemSpec{
						Id:      itemData[i].Id,
						Name:    itemData[i].Name,
						Cost:    dataCost,
						ItemOpt: itemOpt,
					}

					// make ItemList
					fmt.Printf("ItemSpec: %p\n", &spec)
					returnMsg = append(returnMsg, spec)
				}
			} else {
				fmt.Printf("Cost unmatched: %d ~ %d\n", in.CostUpper, in.CostUnder)
			}

		} else {
			fmt.Printf("Has no Item that name is %s\n", in.Name)
		}
	}
	return returnMsg
}

func isOutOfIdx(slice []data.ItemOpt, cnt int) bool {

	if cnt == len(slice) {
		return true
	} else {
		return false
	}
}

// returns ItemSpec, ItemOpt, dataCount
func SellController(in ItemQuery) (data.ItemSpec, int) {

	dataCount := data.GetDataCount()

	var optionList []data.ItemOpt
	queryOpt := in.QueryOpt
	for j := 0; j < len(queryOpt); j++ {
		option := data.ItemOpt{
			OptName: queryOpt[j].OptName,
			Value:   int(queryOpt[j].Under),
		}
		optionList = append(optionList, option)
	}

	// Sell data
	itemSpec := data.ItemSpec{
		Id:      dataCount,
		Name:    in.Name,
		Cost:    int(in.CostUnder),
		ItemOpt: optionList,
	}
	data.SetData(itemSpec)
	data.SetDataCount(dataCount + 1)

	data.LogData()
	return itemSpec, dataCount
}

func BuyController(id int) (data.ItemSpec, bool, int) {

	itemSpec, found := data.DeleteData(id)
	if found {
		dataCount := data.GetDataCount() - 1
		data.SetDataCount(dataCount)
		return itemSpec, found, dataCount
	} else {
		return *new(data.ItemSpec), found, -1
	}
}

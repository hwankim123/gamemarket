package controller

import (
	"fmt"
	"gamemarket/server/data"
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

// Main Logic of Searching Item
func GetAllController(in ItemQuery) []data.ItemSpec {

	fmt.Println("")
	log.Printf("Received From Client: name - %v, cost - %d ~ %d, scope - %v\n",
		in.Name, in.CostUpper, in.CostUnder, in.QueryOpt)

	returnMsg := make([]data.ItemSpec, 0)
	itemData := data.GetAllData()

	for i := 0; i < len(itemData); i++ {
		if in.Name == "" || in.Name == itemData[i].Name {
			// item name matched
			fmt.Printf("item name matched : %s\n", itemData[i].Name)

			dataCost := itemData[i].Cost
			if in.CostUpper <= dataCost && dataCost <= in.CostUnder {
				// item cost matched
				fmt.Printf("%s cost matched : %d ~ %d\n", itemData[i].Name, in.CostUpper, in.CostUnder)
				var optName string
				var upper int
				var under int
				// dataOpt : option of matched item
				dataOpt := itemData[i].ItemOpt

				// flag : false when item option unmatched
				foundItem := true
				for j := 0; j < len(in.QueryOpt); j++ { // for loop each query option(0 ~ 3)
					// index of dataOpt
					cnt := 0
					for { // while cnt == len(dataOpt)
						optName = in.QueryOpt[j].OptName
						if optName == "" || optName == dataOpt[cnt].OptName { // if option name matched
							fmt.Printf("%s's option name matched : %d.%s\n", itemData[i].Name, j, optName)
							upper = in.QueryOpt[j].Upper
							under = in.QueryOpt[j].Under
							// value : value of matched option
							value := dataOpt[cnt].Value

							if upper <= value && value <= under { // if option value matched
								fmt.Printf("%s's option : %s's value %d ~ %d matched : %d\n",
									in.Name, optName, in.CostUpper, in.CostUnder, value)
								break
							} else {
								cnt++
								if isOutOfIdx(dataOpt, cnt) {
									// item unmatched : all the option values of matched item doesn't matched
									fmt.Printf("%s's option %s %d ~ %d : all the option values of matched item doesn't matched\n\n",
										in.Name, in.QueryOpt[j].OptName, in.QueryOpt[j].Upper, in.QueryOpt[j].Under)
									foundItem = false
									break
								}
							}
						} else {
							cnt++
							if isOutOfIdx(dataOpt, cnt) {
								// item unmatched : item name doesn't matched
								fmt.Printf("%s's option %s : option name doesn't matched\n\n",
									in.Name, in.QueryOpt[j].OptName)
								foundItem = false
								break
							}
						}
					}
				}
				if foundItem { // all queries matched
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
					fmt.Printf("%s matched : item id : %d\n\n", spec.Name, spec.Id)
					returnMsg = append(returnMsg, spec)
				}
			} else {
				// item unmatched : item cost doesn't matched
				fmt.Printf("%s's cost %d ~ %d : item cost doesn't matched: %d\n\n",
					in.Name, in.CostUpper, in.CostUnder, dataCost)
			}
		} else {
			// item unmatched : item name doesn't matched
			fmt.Printf("%s : item name doesn't matched: %s\n\n", in.Name, itemData[i].Name)
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

// Main Logic of Searching Item
func SellController(in ItemQuery) data.ItemSpec {

	dataCount := data.GetDataCount()

	// list of Item Option. member value of ItemSpec
	var optionList []data.ItemOpt
	queryOpt := in.QueryOpt
	for j := 0; j < len(queryOpt); j++ {
		option := data.ItemOpt{
			OptName: queryOpt[j].OptName,
			Value:   int(queryOpt[j].Under),
		}
		optionList = append(optionList, option)
	}

	// make ItemSpec
	itemSpec := data.ItemSpec{
		Id:      dataCount,
		Name:    in.Name,
		Cost:    int(in.CostUnder),
		ItemOpt: optionList,
	}

	// append Data
	data.SetData(itemSpec)
	data.SetDataCount(dataCount + 1)
	// print Data
	data.LogData()

	return itemSpec
}

// Main Logic of Buying Item
func BuyController(id int) (data.ItemSpec, bool, int) {

	// delete data if id matched
	itemSpec, found := data.DeleteData(id)

	if found { // id unmatched or out of index
		dataCount := data.GetDataCount() - 1
		data.SetDataCount(dataCount)
		return itemSpec, found, dataCount
	} else {
		return *new(data.ItemSpec), found, -1
	}
}

package data

import (
	"fmt"
	"math/rand"
)

// same structure as items.proto.ItemOption
type ItemOpt struct {
	OptName string
	Value   int
}

// same structure as items.proto.ItemSpec
type ItemSpec struct {
	Id      int
	Name    string
	Cost    int
	ItemOpt []ItemOpt
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

var dataCount = 10

const MIN = 0
const MAX = 100
const COST_MAX = 10000

func GetAllData() []ItemSpec {
	return data
}

func GetDataById(id int) ItemSpec {
	return data[id]
}

func SetData(itemSpec ItemSpec) {
	data = append(data, itemSpec)
}

func GetDataCount() int {
	return dataCount
}

func SetDataCount(c int) {
	dataCount = c
}

// random generate data
func GenerateData() {
	nameLen := len(nameSet)
	for i := 0; i < dataCount; i++ {
		randName := nameSet[rand.Intn(nameLen)]
		var optionList []ItemOpt
		switch randName[0] {
		case 'W':
			optionList = GenerateOption(weaponOptNameSet)
		case 'A':
			optionList = GenerateOption(armourOptNameSet)
		case 'H':
			optionList = GenerateOption(helmetOptNameSet)
		case 'B':
			optionList = GenerateOption(bootsOptNameSet)
		}
		item := ItemSpec{
			Id:      i,
			Name:    randName,
			Cost:    rand.Intn(COST_MAX),
			ItemOpt: optionList,
		}
		data = append(data, item)
	}
}

func GenerateOption(optNameSet []string) []ItemOpt {
	var optionList []ItemOpt
	for j := 0; j < len(optNameSet); j++ {
		option := ItemOpt{OptName: optNameSet[j], Value: rand.Intn(MAX)}
		if option.Value != 0 {
			optionList = append(optionList, option)
		}
	}
	return optionList
}

func LogData() {
	fmt.Println("==== data list ====")
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

func DeleteData(id int) (ItemSpec, bool) {
	for i := 0; i < len(data); i++ {
		if id == data[i].Id {
			returnData := data[i]
			data = append(data[:i], data[i+1:]...)
			fmt.Printf("==== data deleted %d ====\n", id)
			for i := 0; i < len(data); i++ {
				fmt.Println(data[i])
			}
			return returnData, true
		}
	}
	fmt.Printf("==== data delete: error occured: %d no exists\n", id)
	return *new(ItemSpec), false
}

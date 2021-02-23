package main

import "strconv"

type Knapsack struct {
	maxWeight   int
	totalValue  int
	totalWeight int
	numOfItem   int
	itemList    []*Item
}

func NewKnapsack(maxWeight int, maxItem int) *Knapsack {
	itemList := make([]*Item, maxItem)
	knapsack := Knapsack{maxWeight: maxWeight, totalValue: 0, totalWeight: 0, numOfItem: 0, itemList: itemList}
	return &knapsack
}

func (knap *Knapsack) addItem(item *Item) {
	if knap.totalWeight+item.weight <= knap.maxWeight {
		knap.itemList[knap.numOfItem] = item
		knap.totalValue += item.value
		knap.totalWeight += item.weight
		knap.numOfItem++
	}
}

func (knap *Knapsack) showItems() (result string) {
	for i := 0; i < knap.numOfItem; i++ {
		result += knap.itemList[i].name + " "
	}
	return
}

func (knap *Knapsack) toString() (res string) {
	res = ""
	res += strconv.Itoa(len(knap.itemList)) + "\n"
	for _, item := range knap.itemList {
		res += item.name + " " + str(item.value) + " " + str(item.weight) + "\n"
	}
	res += str(knap.maxWeight) + "\n"
	res += "weight: " + str(knap.totalWeight)
	return
}

func str(value int) (res string) {
	res = strconv.Itoa(value)
	return
}

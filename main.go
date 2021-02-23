package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var FILE_NAME string

var numOfItem int
var maxWeight int
var itemList []Item

var threads int

func main() {

	cpus := runtime.NumCPU()
	threads = 8

	fmt.Println("Number of cores: ", cpus)

	// read file
	arg := os.Args[1:]
	FILE_NAME = arg[0]

	if len(arg) == 2 {
		threads, _ = strconv.Atoi(arg[1])
	}

	fmt.Println("Number of thread: ", threads)

	file, err := os.Open(FILE_NAME)
	defer file.Close()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// init vars
	reader := bufio.NewReader(file)
	line, _, err := reader.ReadLine()
	numOfItem, _ = strconv.Atoi(strings.Trim(string(line), " "))

	fmt.Printf("Number of Items: %d\n", numOfItem)
	itemList = make([]Item, numOfItem)
	for index := 0; index < numOfItem; index++ {
		str, _, _ := reader.ReadLine()
		s := strings.ReplaceAll(strings.Trim(string(str), " "), "  ", " ")
		strs := strings.Split(string(s), " ")
		name := strs[0]
		value, _ := strconv.Atoi(string(strs[1]))
		weight, _ := strconv.Atoi(string(strs[2]))
		itemList[index] = Item{name, value, weight}
	}
	line, _, err = reader.ReadLine()
	maxWeight, _ = strconv.Atoi(string(line))

	// test
	start := time.Now()
	bruteForce()
	end := time.Now()
	fmt.Printf("Total runtime: %s\n", end.Sub(start))

}

func bruteForce() {
	cases := 1 << uint(numOfItem)

	//fmt.Printf("cases %d\n", cases)
	// how many cases for a thread
	result := cases / threads
	reminder := cases % threads
	ch := make(chan *Knapsack)
	// equally distribute work
	runs := 0
	for fromCase := 0; fromCase < cases; runs++ {
		toCase := fromCase + result
		if reminder != 0 {
			toCase++
			reminder--
		}
		//fmt.Printf("thread from %d to %d \n", fromCase, toCase)
		go Run(fromCase, toCase, ch)
		fromCase = toCase
	}
	// receive and compare the result
	var bestOne Knapsack
	for r := runs; r > 0; r-- {
		knapsack := <-ch
		if r == runs || knapsack.totalValue > bestOne.totalValue {
			bestOne = *knapsack
		}
	}
	fmt.Printf("The max value is: %d\n", bestOne.totalValue)
	fmt.Println(bestOne.showItems())
}

func Run(fromCase int, toCase int, ch chan *Knapsack) {
	var bestOne Knapsack
	for currCase := fromCase; currCase < toCase; currCase++ {
		knapsack := *NewKnapsack(maxWeight, numOfItem)
		//fmt.Println(knapsack.toString())
		// for all kind of items
		for whichItem := 0; whichItem < numOfItem; whichItem++ {
			var mask int
			mask = int(math.Pow(2, float64(whichItem)))
			// only add items corresponding to the current case.
			if (currCase/mask)%2 != 0 {
				// add this item
				knapsack.addItem(&itemList[whichItem])
			}
		}
		// find the best one
		if fromCase == currCase || knapsack.totalValue > bestOne.totalValue {
			bestOne = knapsack
		}
	}
	// report result
	ch <- &bestOne
	return
}

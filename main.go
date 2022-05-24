package main

import "fmt"

type outputType map[int]*data

func (o *outputType) String() string {
	var s string
	for _, v := range *o {
		s += v.String()
	}
	return s
}

type data struct {
	number    int
	maxProfit int
	orders    []order
}

func (d *data) String() string {
	var s string
	for _, v := range d.orders {
		s += v.String() + ","
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}

	return fmt.Sprintln("Case ", d.number) +
		fmt.Sprintln("Max profit: ", d.maxProfit) +
		fmt.Sprintln("Order: ", s)
}

type order struct {
	bestOrders [][]int
	profit     int
}

func (o *order) String() string {
	var s string
	for _, v := range o.bestOrders {
		s += fmt.Sprint(v)
	}
	return s
}

func main() {
	var z int

	output := outputType{}

	index := 0
	for {
		fmt.Scan(&z)
		if z == 0 {
			break
		}
		var trees [][]int
		for i := 0; i < z; i++ {
			var e int
			fmt.Scan(&e)
			var tree []int
			for j := 0; j < e; j++ {
				var x int
				fmt.Scan(&x)
				tree = append(tree, x)
			}
			trees = append(trees, tree)
		}

		output[index] = makeData(index, trees)
		index++
	}

	fmt.Println(output.String())
}

func makeData(index int, trees [][]int) *data {
	return &data{
		number:    index + 1,
		maxProfit: calculateMaxProfit(calculateOrders(trees)),
		orders:    calculateOrders(trees),
	}
}

func calculateOrders(trees [][]int) []order {
	var orders []order
	for tree := range trees {
		profit, bestOrders := calculateBestOrdersAndProfit(trees[tree])
		orders = append(orders, order{
			bestOrders: bestOrders,
			profit:     profit,
		})
	}
	return orders
}

func calculateBestOrdersAndProfit(tree []int) (int, [][]int) {
	var bestOrders [][]int
	var profit int
	for i := 0; i < len(tree); i++ {
		var orders [][]int
		for j := 0; j < len(tree); j++ {
			if i != j {
				orders = append(orders, []int{tree[i], tree[j]})
			}
		}
		var orderProfit int
		for _, v := range orders {
			orderProfit += v[0] + v[1]
		}
		if orderProfit > profit {
			profit = orderProfit
			bestOrders = orders
		}
	}
	return profit, bestOrders

}

func calculateMaxProfit(orders []order) int {
	var maxProfit int
	for _, v := range orders {
		if v.profit > maxProfit {
			maxProfit = v.profit
		}
	}
	return maxProfit
}

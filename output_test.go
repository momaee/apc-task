package main

import (
	"fmt"
	"testing"
)

func TestOutput(t *testing.T) {
	output := outputType{}
	output[0] = &data{
		number:    1,
		maxProfit: 4,
		orders: []order{
			{
				bestOrders: [][]int{
					{1, 3, 2},
					{2, 3, 1},
				},
			},
		},
	}

	output[1] = &data{
		number:    2,
		maxProfit: 8,
		orders: []order{
			{
				bestOrders: [][]int{
					{1, 2, 1},
					{2, 1, 1},
				},
			},
			{
				bestOrders: [][]int{
					{1, 2},
					{2, 1},
				},
			},
			{
				bestOrders: [][]int{
					{1, 4},
				},
			},
		},
	}

	fmt.Println(output.String())
}

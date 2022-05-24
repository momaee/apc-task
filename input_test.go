package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	tree := []int{
		2, 3, 1,
	}

	profit, bestOrders := calculateBestOrdersAndProfit(tree)
	fmt.Println("profit:", profit)
	fmt.Println("bestOrders:", bestOrders)
	assert.Equal(t, 4, profit)
	assert.Equal(t, [][]int{
		{1, 3, 2},
		{2, 3, 1},
	}, bestOrders)

	trees := [][]int{
		{2, 3, 1},
	}

	maxProfit := calculateMaxProfit(calculateOrders(trees))
	assert.Equal(t, maxProfit, 4)

	trees = [][]int{
		{1, 2, 1},
		{1, 2},
		{1, 4},
	}

	maxProfit = calculateMaxProfit(calculateOrders(trees))
	assert.Equal(t, maxProfit, 8)

}

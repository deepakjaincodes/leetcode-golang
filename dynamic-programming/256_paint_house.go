package main

/*
256. Paint House
There are a row of n houses, each house can be painted with one of the three colors: red, blue or green. The cost of painting each house with a certain color is different. You have to paint all the houses such that no two adjacent houses have the same color.

The cost of painting each house with a certain color is represented by a n x 3 cost matrix. For example, costs[0][0] is the cost of painting house 0 with color red; costs[1][2] is the cost of painting house 1 with color green, and so on... Find the minimum cost to paint all houses.

Note: All costs are positive integers.

Example:

Input: [[17,2,17],[16,16,5],[14,3,19]] Output: 10 Explanation: Paint house 0 into blue, paint house 1 into green, paint house 2 into blue. Minimum cost: 2 + 5 + 3 = 10.
*/

import (
	"fmt"
	"math"
)

func key (i, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func minCostHelp(i, prevSelectedIndex int, costs [][]int, seen map[string]int)int {

	ky := key(i, prevSelectedIndex)

	if value, saw := seen[ky]; saw {
		return value
	}

	min := math.MaxInt32

	for j := 0; j < len(costs[i]); j++ {

		if j != prevSelectedIndex {
			if i < len(costs)-1 {
				min = int(math.Min(float64(min), float64(minCostHelp(i+1, j, costs, seen) + costs[i][j])))
			} else {
				min = int(math.Min(float64(min), float64(costs[i][j])))
			}
		}

	}

	seen[ky] = min
	return min
}

func minCost(costs [][]int) int {

	if len(costs) == 0 {
		return 0
	}

	seen := make(map[string] int, len(costs)*len(costs[0]))

	return minCostHelp(0, -1, costs, seen)

}

func main() {
	fmt.Printf("min cost is %d\n", minCost([][]int{{5, 8, 6}, {19, 14, 13}, {7, 5, 12}, {14, 15, 17}, {3, 20, 10}}))
}
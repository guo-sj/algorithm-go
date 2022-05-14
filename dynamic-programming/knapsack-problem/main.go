// Example code for https://guo-sj.github.io/algorithm/2022/05/14/dynamic-programming.html
package main

import "fmt"

func main() {
	knapsack()
	fmt.Println(getResult())
	getMatrixAndItems()
}

const (
	maxLen       = 5
	knapsackSize = 5
	itemSize     = 5
)

var (
	price       map[string]int      // item's price
	weight      map[string]int      // item's weight
	yAxeMap     map[int]string      // map y axe index to items
	xAxeMap     map[int]int         // map x axe index to knapsack size
	valueToItem map[int][]string    // items of every cell
	matrix      [maxLen][maxLen]int // dynamic programming grid
)

func init() {
	price = make(map[string]int)
	price["guitar"] = 1500
	price["stereo"] = 3000
	price["laptop"] = 2000
	price["iphone"] = 2000
	price["mp3"] = 1000

	weight = make(map[string]int)
	weight["guitar"] = 1
	weight["stereo"] = 4
	weight["laptop"] = 3
	weight["iphone"] = 1
	weight["mp3"] = 1

	yAxeMap = make(map[int]string)
	yAxeMap[0] = "guitar"
	yAxeMap[1] = "stereo"
	yAxeMap[2] = "laptop"
	yAxeMap[3] = "iphone"
	yAxeMap[4] = "mp3"

	xAxeMap = make(map[int]int)
	xAxeMap[0] = 1
	xAxeMap[1] = 2
	xAxeMap[2] = 3
	xAxeMap[3] = 4
	xAxeMap[4] = 5

	valueToItem = make(map[int][]string)
}

// knapsack fills the cells of matrix
func knapsack() {
	for i := 0; i < itemSize; i++ {
		for j := 0; j < knapsackSize; j++ {
			matrix[i][j] = getValue(i, j)
		}
	}
}

// getValue returns correct value of each cell
func getValue(i, j int) int {
	var previousMaxValue, possibleMaxValue int

	if i < 1 {
		previousMaxValue = 0
	} else {
		previousMaxValue = matrix[i-1][j]
	}
	item := yAxeMap[i]
	leftWeight := xAxeMap[j] - weight[item]
	index := i*knapsackSize + j

	// calculate possibleMaxValue
	if leftWeight < 0 {
		// the weight of current item is larger than the current knapsackSize
		// return previousMaxValue
		if i < 1 {
			valueToItem[index] = []string{}
			return previousMaxValue
		}
		previousIndex := (i-1)*knapsackSize + j
		valueToItem[index] = make([]string, len(valueToItem[previousIndex]), cap(valueToItem[previousIndex]))
		copy(valueToItem[index], valueToItem[previousIndex])
		return previousMaxValue
	} else if leftWeight == 0 {
		possibleMaxValue = price[item]
	} else {
		if i < 1 {
			possibleMaxValue = price[item]
		} else {
			possibleMaxValue = price[item] + matrix[i-1][leftWeight-1]
		}
	}

	if possibleMaxValue > previousMaxValue {
		if leftWeight == 0 {
			valueToItem[index] = []string{item}
		} else {
			// copy items of leftWeight cell to current cell's items
			// and append current item to current cell's items
			if i < 1 {
				valueToItem[index] = []string{item}
				return possibleMaxValue
			}
			leftWeightIndex := (i-1)*knapsackSize + leftWeight - 1
			valueToItem[index] = make([]string, len(valueToItem[leftWeightIndex]), cap(valueToItem[leftWeightIndex])+1)
			copy(valueToItem[index], valueToItem[leftWeightIndex])
			valueToItem[index] = append(valueToItem[index], item)
		}
		return possibleMaxValue
	} else {
		// Because previousMaxValue >= possibleMaxValue > 0, the variable i couldn't less than 1.
		previousIndex := (i-1)*knapsackSize + j
		valueToItem[index] = make([]string, len(valueToItem[previousIndex]), cap(valueToItem[previousIndex]))
		copy(valueToItem[index], valueToItem[previousIndex])
		return previousMaxValue
	}
}

// getResult gets final result of the knapsack problem
func getResult() ([]string, int) {
	i := itemSize - 1
	j := knapsackSize - 1
	index := i*knapsackSize + j
	return valueToItem[index], matrix[i][j]
}

// getMatrixAndItems prints matrix and valueToItem
func getMatrixAndItems() {
	i := itemSize - 1
	j := knapsackSize - 1
	length := i*knapsackSize + j
	for i := 0; i < itemSize; i++ {
		fmt.Println(matrix[i])
	}
	for i := 0; i <= length; i++ {
		fmt.Printf("[%d] %v\n", i, valueToItem[i])
	}
}

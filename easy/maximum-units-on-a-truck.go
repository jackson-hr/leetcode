package easy

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	ans := 0
	n := len(boxTypes)
	for i := 0; i < n; i++ {
		maxUnit := 0
		maxUnitIdx := 0
		for j := 0; j < n; j++ {
			if boxTypes[j][1] > maxUnit && boxTypes[j][0] > 0 {
				maxUnit = boxTypes[j][1]
				maxUnitIdx = j
			}
		}
		if truckSize < boxTypes[maxUnitIdx][0] {
			ans += truckSize * maxUnit
			return ans
		}
		ans += boxTypes[maxUnitIdx][0] * maxUnit
		truckSize -= boxTypes[maxUnitIdx][0]
		boxTypes[maxUnitIdx][0] = 0
	}
	return ans
}

func maximumUnits1(boxTypes [][]int, truckSize int) int {
	var ans = 0
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	n := len(boxTypes)
	for i := 0; i < n; i++ {
		if truckSize < boxTypes[i][0] {
			ans += truckSize * boxTypes[i][1]
			break
		}
		ans += boxTypes[i][0] * boxTypes[i][1]
		truckSize -= boxTypes[i][0]
	}
	return ans
}

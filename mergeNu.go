package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	//避免引发panic
	if len(intervals) == 0 {
		return intervals
	}

	//GO标准库用法 按区间起始位置升序，每个区间的第一个数字进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//拿到第一个一维数组
	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := merged[len(merged)-1]
		current := intervals[i]

		if current[0] <= last[1] {
			// 有重叠，合并区间
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			// 无重叠，添加新区间
			merged = append(merged, current)
		}
	}

	return merged
}

func main() {
	// 正确调用方式,注意传入数据要为2维数组，每个数组的最大为2个。
	fmt.Println("合并后区间:", merge([][]int{{1, 2}, {2, 4}, {5, 10}, {15, 18}}))
}

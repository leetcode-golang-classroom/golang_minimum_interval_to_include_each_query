# golang_minimum_interval_to_include_each_query

You are given a 2D integer array `intervals`, where `intervals[i] = [lefti, righti]` describes the `ith` interval starting at `lefti` and ending at `righti` **(inclusive)**. The **size** of an interval is defined as the number of integers it contains, or more formally `righti - lefti + 1`.

You are also given an integer array `queries`. The answer to the `jth` query is the **size of the smallest interval** `i` such that `lefti <= queries[j] <= righti`. If no such interval exists, the answer is `-1`.

Return *an array containing the answers to the queries*.

## Examples

**Example 1:**

```
Input: intervals = [[1,4],[2,4],[3,6],[4,4]], queries = [2,3,4,5]
Output: [3,3,1,4]
Explanation: The queries are processed as follows:
- Query = 2: The interval [2,4] is the smallest interval containing 2. The answer is 4 - 2 + 1 = 3.
- Query = 3: The interval [2,4] is the smallest interval containing 3. The answer is 4 - 2 + 1 = 3.
- Query = 4: The interval [4,4] is the smallest interval containing 4. The answer is 4 - 4 + 1 = 1.
- Query = 5: The interval [3,6] is the smallest interval containing 5. The answer is 6 - 3 + 1 = 4.

```

**Example 2:**

```
Input: intervals = [[2,3],[2,5],[1,8],[20,25]], queries = [2,19,5,22]
Output: [2,-1,4,6]
Explanation: The queries are processed as follows:
- Query = 2: The interval [2,3] is the smallest interval containing 2. The answer is 3 - 2 + 1 = 2.
- Query = 19: None of the intervals contain 19. The answer is -1.
- Query = 5: The interval [2,5] is the smallest interval containing 5. The answer is 5 - 2 + 1 = 4.
- Query = 22: The interval [20,25] is the smallest interval containing 22. The answer is 25 - 20 + 1 = 6.

```

**Constraints:**

- `1 <= intervals.length <= 105`
- `1 <= queries.length <= 105`
- `intervals[i].length == 2`
- `1 <= lefti <= righti <= 107`
- `1 <= queries[j] <= 107`

## 解析

給定一個 2D 陣列 intervals

每個 intervals[i] = [ $left_i, right_i$ ] 代表某個區間 $left_i$ ≤ values ≤ $right_i$

給定一個整數陣列 queries

每個 queries[j] 代表要找的值

對於每個 queries[j] 希望找到一個 intervals[i] = [ $left_i, right_i$ ] 使得 $left_i$ ≤ values ≤ $right_i$

且讓 query_size[j] = intervals[i][1] - intervals[i][0] + 1 最小

題目要求一個演算法 找出給定的 intervals, queries 中所有 query_size 的最小值

另外當 queries[i] 找不到符合的 intervals[j] 則 query_size[i] = -1

這題的困難點在於如何有系統的去找出所有可能的 intervals 並找出最小值

要有順序性的找尋

可以透過把 intervals 根據 $left_i$ 來做由小到大的排序 

有這樣的排序後就可以知道當遇到一個不符合的 intervals 時 ， 其後面的 intervals 都可以不需要再比

同樣地，把 queries 做由小到大的排序

也可以讓 queries[i] 利用之前 ≤ queries[i] 值來做減少搜尋的部份

然而要真找到 size 最小的 interval

需要利用 一個 MinHeap 把 size 跟 left 做由小到大的排序

這樣每次先把有可能的 interval 放入 minHeap

然後當發現 minHeap 最小值 < queries[i] 

則 把 minHeap 掉直到 有 right ≥ queries 

![](https://i.imgur.com/L30mDcY.png)
![](https://i.imgur.com/yLFbBet.png)


這樣的話, 就可以透過時間複雜度 O(nlogn + qlogq)  來做運算 其中 n = len(interval), q = len(queries)

空間複雜度 O(n+q) 因為需要把所有 interval 放入 MinHeap 還有儲存每個 queries[i] 的最小區間大小

## 程式碼
```go
package sol

import (
	"container/heap"
	"sort"
)

type ByStart [][]int

func (a ByStart) Len() int {
	return len(a)
}
func (a ByStart) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}
func (a ByStart) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type Record struct {
	size, right int
}
type MinHeap []Record

func (h *MinHeap) Len() int {
	return len(*h)
}
func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].size < (*h)[j].size ||
		((*h)[i].size == (*h)[j].size && (*h)[i].right < (*h)[j].right)
}
func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *MinHeap) Push(val interface{}) {
	*h = append(*h, val.(Record))
}
func minInterval(intervals [][]int, queries []int) []int {
	// sort intervals, copy queries
	sort.Sort(ByStart(intervals))
	nLen := len(intervals)
	qLen := len(queries)
	bufferQueries := make([]int, qLen)
	result := make([]int, qLen)
	hash := make(map[int]int)
	copy(bufferQueries, queries)
	sort.Ints(bufferQueries)
	pos := 0
	pq := MinHeap{}
	heap.Init(&pq)
	for _, q := range bufferQueries {
		for pos < nLen && intervals[pos][0] <= q {
			heap.Push(&pq, Record{size: intervals[pos][1] - intervals[pos][0] + 1,
				right: intervals[pos][1],
			})
			pos++
		}
		for pq.Len() > 0 && pq[0].right < q {
			heap.Pop(&pq)
		}
		if pq.Len() > 0 {
			hash[q] = pq[0].size
		} else {
			hash[q] = -1
		}
	}
	for pos := 0; pos < qLen; pos++ {
		result[pos] = hash[queries[pos]]
	}
	return result
}
```
## 困難點

1. 需要想出有系統性找尋最小區間 size 的方式
2. 透過排除法來讓 搜訊範圍變小
3. 目前的找法並非很直覺

## Solve Point

- [x]  需要建立一個新的陣列 bufferQueries 來複製 queries
- [x]  建立一個 hashTable hash 來儲存每個 queries 的最小區間大小
- [x]  對 bufferQueries做排序, 對intervals 做以left 排序
- [x]  初始化 pos = 0 代表目前查找到的最後 intervals index, pq = MinHeap
- [x]  由小到大對每個 bufferQueries 內的值 q 做以下檢查
- [x]  while pos ≤ len(intervals) && intervals[pos][0] ≤ q, 則 更新 把 size = intervals[pos][1] - intervals[pos][0] + 1, right = intervals[pos][1] 放入 pq
- [x]  while pq.Len() > 0 && pq[0].right < q, 則更 把 pq  pop 出目前最小的值
- [x]  if pq.Len() > 0 則 hash[q] = pq[0].size , 否則 -1
- [x]  當跑完所有 queries 回傳所有對應 queries 的 sizes
package greedy_algorithm

import (
	"container/heap"
	"sort"
)

/*
题目一：输入一个数组，返回分割的最小代价。
例如：给定数组{10,20,30}，代表一共三个人，整块金条长度为10+20+30=60. 金条要分成10,20,30三个部分。
如果，先把长度60的金条分成10和50，花费60再把长度50的金条分成20和30，花费50 一共花费110铜板。但是如
果，先把长度60的金条分成30和30，花费60再把长度30金条分成10和20，花费30 一共花费90铜板。
解题思路：利用一个小根堆，首次取出两个最小值，然后将相加之后的和放入小根堆中，循环该步骤，直至小根堆中
仅有一个数，将每次产生的中间数累加，得出最终值即为分割的最下代价
*/
type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *minHeap) Min() interface{} {
	return (*h)[0]
}

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

func (h *maxHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *maxHeap) Max() interface{} {
	return (*h)[0]
}

func cutGoldBar(arr []int) int {
	var minHeapInt heap.Interface = new(minHeap)
	var tempArr []int
	var minValueSum int
	var sum int
	for _, v := range arr {
		heap.Push(minHeapInt, v)
	}
	for minHeapInt.Len() > 1 {
		minValueSum = heap.Pop(minHeapInt).(int) + heap.Pop(minHeapInt).(int)
		tempArr = append(tempArr, minValueSum)
		heap.Push(minHeapInt, minValueSum)
	}
	for _, v := range tempArr {
		sum += v
	}
	return sum
}

/*
题目二：输入： 参数1，正数数组costs 参数2，正数数组profits 参数3，正数k 参数4，正数m
costs[i]表示i号项目的花费 profits[i]表示i号项目在扣除花费之后还能挣到的钱(利润) k表示你不能并行、
只能串行的最多做k个项目 m表示你初始的资金
说明：你每做完一个项目，马上获得的收益，可以支持你去做下一个 项目。
输出：你最后获得的最大钱数。
解题思路：以项目花费初始化一个小根堆，首先取出小根堆中满足小于初始资金的项目，然后将取出的项目以项目
利润初始化大根堆，然后在大根堆中取出利润最大项目执行，然后将获取利润加上初始资金，循环执行上述步骤，最后
获取的项目收入即为最大收入
*/
type programInfo struct {
	cost int
	profit int
}
//以项目收益构建构建大根堆
type maxHeapProInfo []programInfo

func (m maxHeapProInfo) Len() int {
	return len(m)
}

func (m maxHeapProInfo) Less(i, j int) bool {
	return m[i].profit > m[j].profit
}

func (m maxHeapProInfo) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeapProInfo) Pop() (v interface{}) {
	value := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return value
}

func (m *maxHeapProInfo) Push(v interface{}) {
	*m = append(*m, v.(programInfo))
}

//以项目成本构建小根堆
type minHeapCostInfo []programInfo

func (c minHeapCostInfo) Len() int {
	return len(c)
}

func (c minHeapCostInfo) Less(i, j int) bool {
	return c[i].cost < c[j].cost
}

func (c minHeapCostInfo) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c *minHeapCostInfo) Push(v interface{}) {
	*c = append(*c, v.(programInfo))
}

func (c *minHeapCostInfo) Pop() interface{} {
	value := (*c)[len(*c)-1]
	*c = (*c)[:len(*c)-1]
	return value
}

func (c *minHeapCostInfo) Min() int {
	return (*c)[0].cost
}

func getProgramProfits(costs []int, profits []int, k, m int) int {
	var projectInfo []programInfo
	var nums int
	for k, v := range costs {
		projectInfo = append(projectInfo, programInfo{v, profits[k]})
	}
	var minCostHeap heap.Interface = new(minHeapCostInfo)
	var maxProfitHeap heap.Interface = new(maxHeapProInfo)
	for _, v := range projectInfo {
		heap.Push(minCostHeap, v)
	}
	//将小根堆中小于本金m的元素全部取出压入按利润构建的大根堆
	for nums < k {
		for minCostHeap.Len() > 0 && minCostHeap.(*minHeapCostInfo).Min() <= m {
			heap.Push(maxProfitHeap, heap.Pop(minCostHeap))
		}
		if maxProfitHeap.Len() == 0 {
			return m
		}
		m += heap.Pop(maxProfitHeap).(programInfo).profit
		nums++
	}
	return m
}

/*
题目三：一个数据流中，随时可以取得中位数
解决思路：用大根堆装排序后的左边部分数据，小根堆装排序后右半部分数据，同时要保证，大根堆与小根堆高度差不能超过1
*/
var maxHeapNum heap.Interface = new(maxHeap)
var minHeapNum heap.Interface = new(minHeap)
func insert(num int) {
	//如果大根堆的长度大于小根堆，则平衡大根堆和小根堆，即压入小根堆一个数，同时，要保证压入小根堆的数是当前数据流中的最大值，可先
	//将num压入大根堆，然后弹出大根堆中的最大值压入小根堆，如果小根堆的长度大于大根堆，则采取相反的操作
	if maxHeapNum.Len() > minHeapNum.Len() {
		heap.Push(maxHeapNum, num)
		heap.Push(minHeapNum, heap.Pop(maxHeapNum))
	} else {
		heap.Push(minHeapNum, num)
		heap.Push(maxHeapNum, heap.Pop(minHeapNum))
	}
}

func getMedianNum() float64 {
	if maxHeapNum.Len() == minHeapNum.Len() {
		return (float64(maxHeapNum.(*maxHeap).Max().(int)+minHeapNum.(*minHeap).Min().(int))) / 2
	} else if maxHeapNum.Len() > minHeapNum.Len() {
		return float64(maxHeapNum.(*maxHeap).Max().(int))
	} else {
		return float64(minHeapNum.(*minHeap).Min().(int))
	}
}

/*
题目四：给定一个字符串类型的数组strs，找到一种拼接方式，使得把所有字符串拼起来之后形成的字符串具有最低的字典序。
解决思路：
*/
func lowestString(strs []string) (res string) {
	if len(strs) == 0 {
		return ""
	}
	sort.Slice(strs, func(i, j int ) bool {
		return strs[i] + strs[j] < strs[j] + strs[i]
	})
	for _, v := range strs {
		res += v
	}
	return res
}

/*
题目五：一些项目要占用一个会议室宣讲，会议室不能同时容纳两个项目的宣讲。 给你每一个项目开始的时间和结束的时间(给你一个数
组，里面 是一个个具体的项目)，你来安排宣讲的日程，要求会议室进行 的宣讲的场次最多。返回这个最多的宣讲场次
解题思路：贪心算法，分析可知，每次选择结束时间最早的会议，依次进行选择，则最后能安排最多的宣讲场次
*/
type preachInfo struct {
	startTime int
	endTime int
}

func mostPreachArrange(preachs []*preachInfo, start int) (res int) {
	sort.Slice(preachs, func(i, j int) bool {
		return preachs[i].endTime < preachs[j].endTime
	})
	for _, v := range preachs {
		if  v.startTime >= start {
			res++
			start = v.endTime
		}
	}
	return res
}

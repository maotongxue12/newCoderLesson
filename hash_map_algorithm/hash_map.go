package hash_map_algorithm

import (
	"math/rand"
	"time"
)

//设计RandomPool结构 【题目】 设计一种结构，在该结构中有如下三个功能： insert(key)：
//将某个key加入到该结构，做到不重复加入。 delete(key)：将原本在结构中的某个key移除。
//getRandom()： 等概率随机返回结构中的任何一个key。
//【要求】 Insert、delete和getRandom方法的时间复杂度都是 O(1)
//思路：用两个map维护该hash结构，

type RandomPool struct {
	keyIndexMap map[string]int
	indexKeyMap map[int]string
	index int	//统计当前map中的key个数
}

func newRandomPool() *RandomPool {
	return &RandomPool{
		keyIndexMap: make(map[string]int),
		indexKeyMap: make(map[int]string),
		index: 0,
	}
}

func (r *RandomPool) insert(key string) {
	if _, ok := r.keyIndexMap[key]; !ok {
		r.keyIndexMap[key] = r.index
		r.indexKeyMap[r.index] = key
		r.index++
	}
}

func (r *RandomPool) getRandom() string {
	if r.index == 0 {
		return ""
	}
	//index
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(r.index)
	return r.indexKeyMap[randomIndex]
}

func (r *RandomPool) delete(key string) {
	if _, ok := r.keyIndexMap[key]; ok {
		r.index--
		lastKey := r.indexKeyMap[r.index]
		curIndex := r.keyIndexMap[key]
		r.keyIndexMap[lastKey] = curIndex
		r.indexKeyMap[curIndex] = lastKey
		delete(r.indexKeyMap, r.index)
		delete(r.keyIndexMap, key)
	}

}

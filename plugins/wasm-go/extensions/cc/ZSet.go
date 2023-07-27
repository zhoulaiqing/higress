package main

type ZSet struct {
	capacity int
	data     map[string]*SkipList
}

func (z *ZSet) ZAdd(key string, score int64) {
	if z.data == nil {
		z.data = make(map[string]*SkipList)
	}

	sl, found := z.data[key]
	if !found {
		sl = NewSkipList()
		z.data[key] = sl
		z.capacity++
	}

	sl.Insert(score)
}

func (z *ZSet) ZCount(key string, min, max int64) int {
	sl, found := z.data[key]
	if !found {
		return 0
	}

	return sl.RangeCount(min, max)
}

func (z *ZSet) ZRange(key string, min, max int64) []int64 {
	sl, found := z.data[key]
	if !found {
		return []int64{}
	}

	return sl.RangeSearch(min, max)
}

func (z *ZSet) ZRemByScore(key string, min, max int64) {
	sl, found := z.data[key]
	if !found {
		return
	}

	sl.RangeDelete(min, max)
}

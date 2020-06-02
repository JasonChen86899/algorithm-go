package graph

import (
	"fmt"
	"testing"
)

func initDisJoinSet() (*DisJoinSet, []*DisJoinSetItem) {
	setNums := uint64(8)
	set := &DisJoinSet{setNums: setNums, father: make(map[*DisJoinSetItem]*DisJoinSetItem)}

	items := make([]*DisJoinSetItem, 0, 8)
	for i := uint64(0); i < setNums; i++ {
		item := &DisJoinSetItem{value:i}
		set.father[item] = item
		items = append(items, item)
	}

	return set, items
}

func TestDisJoinSet_Join(t *testing.T) {
	disJoinSet, items := initDisJoinSet()
	item_0 := items[0]
	item_1 := items[1]
	item_2 := items[2]

	disJoinSet.Join(item_0, item_1)
	disJoinSet.Join(item_1, item_2)
	fmt.Println(disJoinSet.Find(item_0))
}

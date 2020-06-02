package graph

// 并查集 维护一些不相交的集合 支持两种操作 合并两个集合 查询一个元素所处集合
// 初始化时 每个节点的 father都是自身 setNums=节点总数
type DisJoinSetItem struct {
	value interface{}
}

type DisJoinSet struct {
	setNums uint64 // 包含的集合数量
	father  map[*DisJoinSetItem]*DisJoinSetItem
}

func (ds *DisJoinSet) Join(from, to *DisJoinSetItem) {
	ds.setNums--
	ds.father[from] = to
}

func (ds *DisJoinSet) Find(item *DisJoinSetItem) *DisJoinSetItem {
	father := ds.father[item]

	if father == item {
		return father
	} else {
		root := ds.Find(father)
		ds.father[item] = root // 路径压缩
		return root
	}
}

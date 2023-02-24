package wgm

import "sync"

var FindPageOptionSyncPool = sync.Pool{
	New: func() interface{} {
		return new(FindPageOption)
	},
}

func acquireMoney() *FindPageOption {
	return FindPageOptionSyncPool.Get().(*FindPageOption)
}

func releaseMoney(m *FindPageOption) {
	m.selector = nil
	m.fields = nil
	FindPageOptionSyncPool.Put(m)
}

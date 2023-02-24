package wgm

import "sync"

var FindPageOptionSyncPool = sync.Pool{
	New: func() interface{} {
		return new(FindPageOption)
	},
}

func acquireFindPageOption() *FindPageOption {
	return FindPageOptionSyncPool.Get().(*FindPageOption)
}

func releaseFindPageOption(m *FindPageOption) {
	m.selector = nil
	m.fields = nil
	FindPageOptionSyncPool.Put(m)
}

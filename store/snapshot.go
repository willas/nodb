package store

import (
	"github.com/willas/nodb/store/driver"
)

type Snapshot struct {
	driver.ISnapshot
}

func (s *Snapshot) NewIterator() *Iterator {
	it := new(Iterator)
	it.it = s.ISnapshot.NewIterator()

	return it
}

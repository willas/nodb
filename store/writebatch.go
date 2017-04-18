package store

import (
	"github.com/willas/nodb/store/driver"
)

type WriteBatch interface {
	driver.IWriteBatch
}

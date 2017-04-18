package nodb

import (
	"bytes"
	"os"
	"testing"

	"github.com/willas/nodb/config"
	"github.com/willas/nodb/store"
)

func TestDump(t *testing.T) {
	cfgM := new(config.Config)
	cfgM.DataDir = "/tmp/test_ledis_master"

	os.RemoveAll(cfgM.DataDir)

	master, err := Open(cfgM)
	if err != nil {
		t.Fatal(err)
	}

	cfgS := new(config.Config)
	cfgS.DataDir = "/tmp/test_ledis_slave"
	os.RemoveAll(cfgM.DataDir)

	var slave *Nodb
	if slave, err = Open(cfgS); err != nil {
		t.Fatal(err)
	}

	db, _ := master.Select(0)

	db.Set([]byte("a"), []byte("1"))
	db.Set([]byte("b"), []byte("2"))
	db.Set([]byte("c"), []byte("3"))

	if err := master.DumpFile("/tmp/testdb.dump"); err != nil {
		t.Fatal(err)
	}

	if _, err := slave.LoadDumpFile("/tmp/testdb.dump"); err != nil {
		t.Fatal(err)
	}

	it := master.ldb.RangeLimitIterator(nil, nil, store.RangeClose, 0, -1)
	for ; it.Valid(); it.Next() {
		key := it.Key()
		value := it.Value()

		if v, err := slave.ldb.Get(key); err != nil {
			t.Fatal(err)
		} else if !bytes.Equal(v, value) {
			t.Fatal("load dump error")
		}
	}
}

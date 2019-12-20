package store

import (
	"fmt"
	"github.com/tendermint/dex-demo/storeutils"
	"strings"

	dbm "github.com/tendermint/tm-db"

	sdk "github.com/cosmos/cosmos-sdk/store/types"
)

const TablePrefix = "t"

type Table struct {
	db     dbm.DB
	prefix string
}

func NewTable(db dbm.DB, prefix string) *Table {
	return &Table{
		db:     db,
		prefix: fmt.Sprintf("%s/%s", TablePrefix, prefix),
	}
}

func (t *Table) Get(key []byte) []byte {
	return t.db.Get(storeutils.PrefixKeyString(t.prefix, key))
}

func (t *Table) Has(key []byte) bool {
	return t.db.Has(storeutils.PrefixKeyString(t.prefix, key))
}

func (t *Table) Set(key, value []byte) {
	t.db.Set(storeutils.PrefixKeyString(t.prefix, key), value)
}

func (t *Table) Delete(key []byte) {
	t.db.Delete(storeutils.PrefixKeyString(t.prefix, key))
}

func (t *Table) Iterator(start []byte, end []byte, cb IteratorCB) {
	iter := t.db.Iterator(storeutils.PrefixKeyString(t.prefix, start), storeutils.PrefixKeyString(t.prefix, end))
	t.iterate(iter, cb)
}

func (t *Table) ReverseIterator(start []byte, end []byte, cb IteratorCB) {
	iter := t.db.ReverseIterator(storeutils.PrefixKeyString(t.prefix, start), storeutils.PrefixKeyString(t.prefix, end))
	t.iterate(iter, cb)
}

func (t *Table) PrefixIterator(start []byte, cb IteratorCB) {
	start = storeutils.PrefixKeyString(t.prefix, start)
	iter := t.db.Iterator(start, sdk.PrefixEndBytes(start))
	t.iterate(iter, cb)
}

func (t *Table) ReversePrefixIterator(start []byte, cb IteratorCB) {
	start = storeutils.PrefixKeyString(t.prefix, start)
	iter := t.db.ReverseIterator(start, sdk.PrefixEndBytes(start))
	t.iterate(iter, cb)
}

func (t *Table) iterate(iter dbm.Iterator, cb IteratorCB) {
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		k := []byte(strings.TrimPrefix(string(iter.Key()), t.prefix+"/"))
		v := iter.Value()

		if !cb(k, v) {
			return
		}
	}
}

func (t *Table) Substore(prefix string) ArchiveStore {
	return &Table{
		db:     t.db,
		prefix: fmt.Sprintf("%s/%s", t.prefix, prefix),
	}
}

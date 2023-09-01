package wrapper

import (
	cmtdb "github.com/cometbft/cometbft-db"
	cdbm "github.com/cosmos/cosmos-db"
)

type DBWrapper struct {
	cmtdb.DB
}

func NewCosmosDB(db cmtdb.DB) cdbm.DB {
	return &DBWrapper{db}
}

func (db *DBWrapper) Iterator(start, end []byte) (cdbm.Iterator, error) {
	return db.DB.Iterator(start, end)
}

func (db *DBWrapper) ReverseIterator(start, end []byte) (cdbm.Iterator, error) {
	return db.DB.ReverseIterator(start, end)
}

func (db *DBWrapper) NewBatch() cdbm.Batch {
	return NewCosmosBatch(db.DB.NewBatch())
}

func (db *DBWrapper) NewBatchWithSize(size int) cdbm.Batch {
	return NewCosmosBatch(db.DB.NewBatch())
}

type BatchWrapper struct {
	cmtdb.Batch
}

func NewCosmosBatch(batch cmtdb.Batch) cdbm.Batch {
	return &BatchWrapper{batch}
}

func (b *BatchWrapper) GetByteSize() (int, error) {
	return 0, nil
}

package storage

import "github.com/jmoiron/sqlx"

type IStorage interface {
	Book() BookStorageI
}

type storagePg struct {
	db *sqlx.DB
	book BookStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db: db,
		book: NewBookRepo(db),
	}
}

func (s *storagePg) Book() BookStorageI {
	return s.book
}
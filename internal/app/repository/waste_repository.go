package repository

import (
	"github.com/jmoiron/sqlx"
)

type WasteRepository struct {
	db *sqlx.DB
}

func NewWasteRepository(db *sqlx.DB) *WasteRepository {
	return &WasteRepository{
		db: db,
	}
}

var (
	wasteTable = "wastes"
)

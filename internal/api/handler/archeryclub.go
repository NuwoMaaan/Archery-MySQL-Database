package handler

import (
	"database/sql"
)

type ArcherHandler struct {
	db *sql.DB
}

func NewArcherHandler(db *sql.DB) *ArcherHandler {
	return &ArcherHandler{db: db}
}

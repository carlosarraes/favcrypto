package handlers

import "favcrypto/data"

type Handlers struct {
	db *data.DbClient
}

func NewHandlers(db *data.DbClient) *Handlers {
	return &Handlers{db: db}
}

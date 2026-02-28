package handlers

import (
	"github.com/k8s-green/backend/internal/config"
	"github.com/k8s-green/backend/internal/repository"
)

type Handler struct {
	db  *repository.Postgres
	rdb *repository.Redis
	cfg *config.Config
}

func NewHandler(db *repository.Postgres, rdb *repository.Redis, cfg *config.Config) *Handler {
	return &Handler{db: db, rdb: rdb, cfg: cfg}
}

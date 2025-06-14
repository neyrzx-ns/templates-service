package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateTemplate(ctx context.Context) error {
}

func (r *Repository) CreateTemplateGroup(ctx context.Context) error {
	return nil
}

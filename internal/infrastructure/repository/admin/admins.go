package admin

import (
	"context"

	models "github.com/Prohovoch/FollowMe/internal/warehouse/models/administrator"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Creating an interface for our Admin Repository (CRUD operations)
type AdminRepo interface {
	CreateAdmin(ctx context.Context, admin *models.AdminModel) error
	CheckAdminCredetinals(ctx context.Context, id int32) (*models.AdminModel, error)
	UpdateAdminCredetinals(ctx context.Context, admin *models.AdminModel) error
	DeleteAdmin(ctx context.Context, id int32) error
}

// Structure of our concurrency safe pool
type Repository struct {
	db *pgxpool.Pool
}

// Constructor of our concurrency safe pool
func New(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) CreateAdmin(ctx context.Context, admin *models.AdminModel) error {
	return nil
}

func (repo *Repository) CheckAdminCredetinals(ctx context.Context, id int32) (*models.AdminModel, error) {
	return nil, nil
}

func (repo *Repository) UpdateAdminCredetinals(ctx context.Context, admin *models.AdminModel) error {
	return nil
}

func (repo *Repository) DeleteAdmin(ctx context.Context, id int32) error {
	return nil
}

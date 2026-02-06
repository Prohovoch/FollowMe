package admin

import (
	"context"

	mapper "github.com/Prohovoch/FollowMe/internal/infrastructure/repository/mapping/adminsRepoMP"

	models "github.com/Prohovoch/FollowMe/internal/warehouse/models/admin"
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

/*
	Just to achieve a progress with mapping's methods. Will be removed when repo will be ready!
*/

// Initializing an interface
func (repo *Repository) CreateAdmin(ctx context.Context, admin *models.AdminModel) error {
	// will be removed ASAP
	mapper.ToEntity(admin)
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

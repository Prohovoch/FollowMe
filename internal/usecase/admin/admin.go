package admin

import (
	"context"

	repo "github.com/Prohovoch/FollowMe/internal/infrastructure/repository/admin"
	mapper "github.com/Prohovoch/FollowMe/internal/usecase/mapping/adminsServMP"
	"github.com/Prohovoch/FollowMe/internal/warehouse/dto/admin"
)

type AdminService struct {
	repos repo.Repository
}

func New(as repo.Repository) *AdminService {
	return &AdminService{
		repos: as,
	}
}

func (as *AdminService) CreateAdmin(ctx context.Context, dtoAdmin *dto.AdminsRequest) error {
	// Made a fiiler function for commit, changes ASAP
	model := mapper.ToModel(dtoAdmin)
	if err := as.repos.CreateAdmin(ctx, model); err != nil {
		return err
	}
	return nil

}

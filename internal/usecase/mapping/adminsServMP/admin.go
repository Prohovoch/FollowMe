package admins

// Mapping from DTO to db models in use case
import (
	"github.com/Prohovoch/FollowMe/internal/warehouse/dto/admin"
	models "github.com/Prohovoch/FollowMe/internal/warehouse/models/admin"
)

func ToModel(adminReqDTO *dto.AdminsRequest) *models.AdminModel {
	return &models.AdminModel{
		Name:     adminReqDTO.Name,
		Surname:  adminReqDTO.Surname,
		Password: adminReqDTO.Password,
		Email:    adminReqDTO.Email,
	}
}

func ToDTO_NS(adminModel *models.AdminModel) *dto.AdminsResponseNS {
	return &dto.AdminsResponseNS{
		Name:    adminModel.Name,
		Surname: adminModel.Surname,
	}
}

func ToDTO_MailResponse(adminModel *models.AdminModel) *dto.AdminsResponeEmail {
	return &dto.AdminsResponeEmail{
		Email: adminModel.Email,
	}
}

// TODO: Make a hash function on this...
func ToDTO_PassResponse(adminModel *models.AdminModel) *dto.AdminsResponsePassword {
	return &dto.AdminsResponsePassword{
		Password: adminModel.Password,
	}
}

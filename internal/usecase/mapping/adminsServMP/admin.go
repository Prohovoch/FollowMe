package admins

// Mapping from DTO to db models in use case
import (
	"github.com/Prohovoch/FollowMe/internal/warehouse/DTO/admins"
	models "github.com/Prohovoch/FollowMe/internal/warehouse/models/admins"
)

func toModel(adminReqDTO *admins.AdminsRequest) *models.AdminModel {
	return &models.AdminModel{
		Name:     adminReqDTO.Name,
		Surname:  adminReqDTO.Surname,
		Password: adminReqDTO.Password,
		Email:    adminReqDTO.Email,
	}
}

func toDTO_NS(adminModel *models.AdminModel) *admins.AdminsResponseNS {
	return &admins.AdminsResponseNS{
		Name:    adminModel.Name,
		Surname: adminModel.Surname,
	}
}

func toDTO_MailResponse(adminModel *models.AdminModel) *admins.AdminsResponeEmail {
	return &admins.AdminsResponeEmail{
		Email: adminModel.Email,
	}
}

// TODO: Make a hash function on this...
func toDTO_PassResponse(adminModel *models.AdminModel) *admins.AdminsResponsePassword {
	return &admins.AdminsResponsePassword{
		Password: adminModel.Password,
	}
}

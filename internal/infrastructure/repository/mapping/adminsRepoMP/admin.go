package admins

import (
	entity "github.com/Prohovoch/FollowMe/internal/warehouse/entity/admin"
	models "github.com/Prohovoch/FollowMe/internal/warehouse/models/admin"
)

// Create an object named "Admin"

func ToEntity(adminModel *models.AdminModel) *entity.Admin {
	return &entity.Admin{
		Id:       adminModel.Id,
		Name:     adminModel.Name,
		Surname:  adminModel.Surname,
		Password: adminModel.Password,
		Email:    adminModel.Email,
	}
}

// Return a Value Name, Surname to be named in Use Case DTO
func ToModelNS(adminEntity *entity.Admin) *models.AdminModel {
	return &models.AdminModel{
		Name:    adminEntity.Name,
		Surname: adminEntity.Surname,
	}
}

// Same as NS, but with email
func ToModelEmail(adminEntity *entity.Admin) *models.AdminModel {
	return &models.AdminModel{
		Email: adminEntity.Email,
	}
}

// TODO: Think about Password hash...

func ToModelPassword(adminEntity *entity.Admin) *models.AdminModel {
	return &models.AdminModel{
		Password: adminEntity.Password,
	}
}

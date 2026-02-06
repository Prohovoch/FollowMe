package admin

import (
	service "github.com/Prohovoch/FollowMe/internal/usecase/admin"
)

type AdminHandlers struct {
	serv service.AdminService
}

func New(ah service.AdminService) *AdminHandlers {
	return &AdminHandlers{
		serv: ah,
	}
}

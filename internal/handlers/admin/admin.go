package admin

import (
	service "github.com/Prohovoch/FollowMe/internal/usecase/admin"
)

type AdminHandler struct {
	serv service.AdminService
}

func New(ah service.AdminService) *AdminHandler {
	return &AdminHandler{
		serv: ah,
	}
}

// TODO: make CRUD handlers for Admins

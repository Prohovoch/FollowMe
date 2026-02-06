package dto

// Request DTO / WIP
type AdminsRequest struct {
	Name     string
	Surname  string
	Password string
	Email    string
}

// GET (Id)
type AdminsResponeEmail struct {
	Email string
}

type AdminsResponsePassword struct {
	Password string
}

type AdminsResponseNS struct {
	Name    string
	Surname string
}

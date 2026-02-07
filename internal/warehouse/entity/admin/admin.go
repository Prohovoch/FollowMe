package entity

import "github.com/Prohovoch/FollowMe/internal/warehouse/entity/athlete"

// 1: M adnin - athlete
type Admin struct {
	Id       int32
	Name     string
	Surname  string
	Password string
	Email    string
	Athlete  []*athlete.Athlete
}

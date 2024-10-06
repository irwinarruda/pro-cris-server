package models

import (
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type Gender string

const (
	Female Gender = "Female"
	Male   Gender = "Male"
)

func (g Gender) String() string { return string(g) }

func (g *Gender) UnmarshalJSON(b []byte) (err error) {
	return utils.UnmarshalEnum(g, GetGenderString(), b)
}

func GetGenderString() []string {
	return []string{
		Male.String(),
		Female.String(),
	}
}

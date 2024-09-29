package models

type Gender string

const (
	Female Gender = "Female"
	Male   Gender = "Male"
)

func (g Gender) String() string { return string(g) }

func GetGenderString() []string {
	return []string{
		Male.String(),
		Female.String(),
	}
}

package models

type Gender = string

const (
	Female Gender = "Female"
	Male   Gender = "Male"
)

func GetGender() []Gender {
	return []Gender{
		Male,
		Female,
	}
}

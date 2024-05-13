package students

type CreateStudentDTO struct {
	Name              string      `json:"name"`
	BirthDay          *string     `json:"birthDay"`
	DisplayColor      string      `json:"displayColor"`
	Picture           *string     `json:"picture"`
	ParentName        *string     `json:"parentName"`
	ParentPhoneNumber *string     `json:"parentPhoneNumber"`
	HouseAddress      *string     `json:"houseAddress"`
	HouseIdentifier   *string     `json:"hoseInfo"`
	HouseCoordinate   *Coordinate `json:"houseCoordinate"`
	BasePrice         float64     `json:"basePrice"`
}

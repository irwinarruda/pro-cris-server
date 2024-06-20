package students

import (
	"time"

	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type studentEntity struct {
	ID                       int
	Name                     string
	BirthDay                 *string
	DisplayColor             string
	Picture                  *string
	ParentName               *string
	ParentPhoneNumber        *string
	HouseAddress             *string
	HouseIdentifier          *string
	HouseCoordinateLatitude  *float64
	HouseCoordinateLongitude *float64
	BasePrice                float64
	IsDeleted                bool
	CreatedAt                time.Time
	UpdatedAt                time.Time
}

func (s *studentEntity) FromCreateStudent(student CreateStudentDTO) {
	var latitude *float64
	var longitude *float64
	if student.HouseCoordinate != nil {
		latitude = &student.HouseCoordinate.Latitude
		longitude = &student.HouseCoordinate.Longitude
	}
	s.Name = student.Name
	s.BirthDay = student.BirthDay
	s.DisplayColor = student.DisplayColor
	s.Picture = student.Picture
	s.ParentName = student.ParentName
	s.ParentPhoneNumber = student.ParentPhoneNumber
	s.HouseAddress = student.HouseAddress
	s.HouseIdentifier = student.HouseIdentifier
	s.HouseCoordinateLatitude = latitude
	s.HouseCoordinateLongitude = longitude
	s.BasePrice = student.BasePrice
}

func (s *studentEntity) FromUpdateStudent(student UpdateStudentDTO) {
	var latitude *float64
	var longitude *float64
	if student.HouseCoordinate != nil {
		latitude = &student.HouseCoordinate.Latitude
		longitude = &student.HouseCoordinate.Longitude
	}
	s.ID = student.ID
	s.Name = student.Name
	s.BirthDay = student.BirthDay
	s.DisplayColor = student.DisplayColor
	s.Picture = student.Picture
	s.ParentName = student.ParentName
	s.ParentPhoneNumber = student.ParentPhoneNumber
	s.HouseAddress = student.HouseAddress
	s.HouseIdentifier = student.HouseIdentifier
	s.HouseCoordinateLatitude = latitude
	s.HouseCoordinateLongitude = longitude
	s.BasePrice = student.BasePrice
	s.UpdatedAt = time.Now()
}

func (s *studentEntity) ToStudent(routineEntity []routinePlanEntity) Student {
	var coordinate *models.Coordinate
	if s.HouseCoordinateLatitude != nil && s.HouseCoordinateLongitude != nil {
		coordinate = &models.Coordinate{
			Latitude:  *s.HouseCoordinateLatitude,
			Longitude: *s.HouseCoordinateLongitude,
		}
	}
	return Student{
		ID:                s.ID,
		Name:              s.Name,
		BirthDay:          s.BirthDay,
		DisplayColor:      s.DisplayColor,
		Picture:           s.Picture,
		ParentName:        s.ParentName,
		ParentPhoneNumber: s.ParentPhoneNumber,
		HouseAddress:      s.HouseAddress,
		HouseIdentifier:   s.HouseIdentifier,
		HouseCoordinate:   coordinate,
		BasePrice:         s.BasePrice,
		IsDeleted:         s.IsDeleted,
		CreatedAt:         s.CreatedAt,
		UpdatedAt:         s.UpdatedAt,
		Routine: utils.Map(routineEntity, func(rp routinePlanEntity) RoutinePlan {
			return rp.ToRoutinePlan()
		}),
	}
}

type routinePlanEntity struct {
	ID        *int
	IdStudent int
	WeekDay   models.WeekDay
	StartHour int
	Duration  int
	Price     float64
	IsDeleted bool
	CreatedAt time.Time
}

func (r *routinePlanEntity) FromCreateStudentRoutinePlan(routine CreateStudentRoutinePlanDTO, idStudent int, basePrice float64) {
	if routine.Price != nil {
		basePrice = *routine.Price
	}
	r.IdStudent = idStudent
	r.WeekDay = routine.WeekDay
	r.StartHour = routine.StartHour
	r.Duration = routine.Duration
	r.Price = basePrice
	r.CreatedAt = time.Now()
}

func (r *routinePlanEntity) FromUpdateStudentRoutinePlan(routine UpdateStudentRoutinePlanDTO, idStudent int) {
	if routine.ID != nil {
		r.ID = routine.ID
		return
	}
	r.ID = nil
	r.IdStudent = idStudent
	r.WeekDay = *routine.WeekDay
	r.StartHour = *routine.StartHour
	r.Duration = *routine.Duration
	r.Price = *routine.Price
	r.CreatedAt = time.Now()
}

func (r *routinePlanEntity) ToRoutinePlan() RoutinePlan {
	var id *int
	if r.ID != nil {
		id = r.ID
	}
	return RoutinePlan{
		ID:        *id,
		WeekDay:   r.WeekDay,
		StartHour: r.StartHour,
		Duration:  r.Duration,
		Price:     r.Price,
		IsDeleted: r.IsDeleted,
		CreatedAt: r.CreatedAt,
	}
}

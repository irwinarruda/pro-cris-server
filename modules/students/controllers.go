package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

type StudentCtrl struct {
	Db       configs.Db       `inject:"db"`
	Validate configs.Validate `inject:"validate"`
}

func (s StudentCtrl) CreateStudent(c *gin.Context) {
	studentDTO := CreateStudentDTO{}
	err := c.Bind(&studentDTO)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = s.Validate.Struct(studentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	var latitude *float64 = nil
	var longitude *float64 = nil
	if studentDTO.HouseCoordinate != nil {
		latitude = &studentDTO.HouseCoordinate.Latitude
		longitude = &studentDTO.HouseCoordinate.Longitude
	}

	s.Db.Exec(`
    INSERT INTO students(
      name,
      birth_day,
      display_color,
      picture,
      parent_name,
      parent_phone_number,
      house_address,
      house_identifier,
      house_coordinate_latitude,
      house_coordinate_longitude,
      base_price,
      is_deleted
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
  `, studentDTO.Name, studentDTO.BirthDay, studentDTO.DisplayColor, studentDTO.Picture, studentDTO.ParentName, studentDTO.ParentPhoneNumber, studentDTO.HouseAddress, studentDTO.HouseIdentifier, latitude, longitude, studentDTO.BasePrice, false)

	studentsArr := []Student{}
	s.Db.Raw("SELECT * FROM students;").Scan(&studentsArr)
	c.JSON(http.StatusCreated, studentsArr)
}

package students

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proval"
)

var id = 0
var studentsArr = []Student{}

func CreateStudent(c *gin.Context) {
	studentDTO := CreateStudentDTO{}
	err := c.Bind(&studentDTO)
	if err != nil {
		c.String(404, err.Error())
		return
	}

	v := proval.New()
	minVal := v.String("Should be a string").Min(5, "Min 5 items")

	errs := v.ToStringSlice(minVal.Validate(studentDTO.Name))
	if len(errs) > 0 {
		c.JSON(404, errs)
		return
	}

	id++
	student := Student{
		Id:                id,
		Name:              studentDTO.Name,
		BirthDay:          studentDTO.BirthDay,
		DisplayColor:      studentDTO.DisplayColor,
		Picture:           studentDTO.Picture,
		ParentName:        studentDTO.ParentName,
		ParentPhoneNumber: studentDTO.ParentPhoneNumber,
		HouseAddress:      studentDTO.HouseAddress,
		HouseIdentifier:   studentDTO.HouseIdentifier,
		HouseCoordinate:   studentDTO.HouseCoordinate,
		BasePrice:         studentDTO.BasePrice,
		IsDeleted:         false,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	studentsArr = append(studentsArr, student)

	fmt.Println(studentsArr)
	c.Header("Content-Type", "application/json")
	c.JSON(201, studentsArr)
}

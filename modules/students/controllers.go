package students

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

var id = 0
var studentsArr = []Student{}

type StudentCtrl struct {
	Env      configs.Env      `inject:"env"`
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

	id++
	student := studentDTO.ToStudent()
	student.Id = id

	studentsArr = append(studentsArr, student)

	c.JSON(http.StatusCreated, studentsArr)
}

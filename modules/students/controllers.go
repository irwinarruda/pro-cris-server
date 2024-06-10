package students

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(&studentDTO)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	id++
	student := studentDTO.ToStudent()
	student.Id = id

	studentsArr = append(studentsArr, student)

	fmt.Println(studentsArr)
	c.JSON(http.StatusCreated, studentsArr)
}

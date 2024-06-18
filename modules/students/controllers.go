package students

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
)

type StudentCtrl struct {
	Validate configs.Validate `inject:"validate"`
}

func (s StudentCtrl) GetStudents(c *gin.Context) {
	studentRepository := newStudentRepository()
	students := studentRepository.GetAllStudents()
	c.JSON(http.StatusOK, students)
}

func (s StudentCtrl) GetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	studentRepository := newStudentRepository()
	student := studentRepository.GetStudentByID(id)
	c.JSON(http.StatusOK, student)
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

	studentRepository := newStudentRepository()
	id := studentRepository.CreateStudent(studentDTO.ToStudent())
	c.JSON(http.StatusCreated, struct {
		Id int `json:"id"`
	}{Id: id})
}

func (s StudentCtrl) UpdateSudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	studentDTO := UpdateStudentDTO{}
	err = c.Bind(&studentDTO)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = s.Validate.Struct(studentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	studentRepository := newStudentRepository()
	id = studentRepository.UpdateStudent(studentDTO)
	c.JSON(http.StatusCreated, struct {
		Id int `json:"id"`
	}{Id: id})
}

func (s StudentCtrl) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	studentRepository := newStudentRepository()
	studentRepository.DeleteStudentByID(id)
	c.JSON(http.StatusOK, struct {
		Id int `json:"id"`
	}{Id: id})
}

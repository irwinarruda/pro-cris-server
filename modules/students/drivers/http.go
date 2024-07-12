package studentsdrivers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type StudentCtrl struct {
	Validate configs.Validate `inject:"validate"`
}

func NewStudentCtrl() *StudentCtrl {
	return proinject.Resolve(&StudentCtrl{})
}

func (s StudentCtrl) GetStudents(c *gin.Context) {
	studentsDTO := students.GetAllStudentsDTO{}
	studentsDTO.IDUser = c.Value("id_user").(int)
	studentService := students.NewStudentService()
	students := studentService.GetAllStudents(studentsDTO)
	c.JSON(http.StatusOK, students)
}

func (s StudentCtrl) GetStudent(c *gin.Context) {
	idStudent, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError("Invalid student ID.", false, nil))
		return
	}
	studentDTO := students.GetStudentDTO{}
	studentDTO.IDUser = c.Value("id_user").(int)
	studentDTO.ID = idStudent
	studentService := students.NewStudentService()
	student, err := studentService.GetStudentByID(studentDTO)
	if err, ok := err.(utils.AppError); ok {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, student)
}

func (s StudentCtrl) CreateStudent(c *gin.Context) {
	studentDTO := students.CreateStudentDTO{}
	studentDTO.IDUser = c.Value("id_user").(int)
	err := c.Bind(&studentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError("Invalid student data"+err.Error(), false, nil))
		return
	}
	err = s.Validate.Struct(studentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError("Invalid student data"+err.Error(), false, nil))
		return
	}

	studentService := students.NewStudentService()
	id, err := studentService.CreateStudent(studentDTO)
	if err, ok := err.(utils.AppError); ok {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusCreated, struct {
		Id int `json:"id"`
	}{Id: id})
}

func (s StudentCtrl) UpdateSudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError("Invalid student ID.", false, nil))
		return
	}
	studentDTO := students.UpdateStudentDTO{}
	studentDTO.ID = id
	studentDTO.IDUser = c.Value("id_user").(int)
	err = c.Bind(&studentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError("Invalid student data. "+err.Error(), false, nil))
		return
	}
	err = s.Validate.Struct(studentDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError("Invalid student data"+err.Error(), false, nil))
		return
	}
	studentService := students.NewStudentService()
	id, err = studentService.UpdateStudent(studentDTO)
	if err, ok := err.(utils.AppError); ok {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, struct {
		Id int `json:"id"`
	}{Id: id})
}

func (s StudentCtrl) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError("Invalid student ID.", false, nil))
		return
	}
	studentDTO := students.DeleteStudentDTO{}
	studentDTO.ID = id
	studentDTO.IDUser = c.Value("id_user").(int)
	studentService := students.NewStudentService()
	id, err = studentService.DeleteStudent(studentDTO)
	if err, ok := err.(utils.AppError); ok {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, struct {
		Id int `json:"id"`
	}{Id: id})
}

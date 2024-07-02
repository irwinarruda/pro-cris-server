package students

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
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
	studentsDTO := GetAllStudentsDTO{}
	studentsDTO.IDUser = c.Value("id_user").(int)
	studentService := NewStudentService()
	students := studentService.GetAllStudents(studentsDTO)
	c.JSON(http.StatusOK, students)
}

func (s StudentCtrl) GetStudent(c *gin.Context) {
	idStudent, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.NewAppError("Invalid student ID.", false, nil))
		return
	}
	studentDTO := GetStudentDTO{}
	studentDTO.IDUser = c.Value("id_user").(int)
	studentDTO.ID = idStudent
	studentService := NewStudentService()
	student, err := studentService.GetStudentByID(studentDTO)
	if err, ok := err.(utils.AppError); ok {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, student)
}

func (s StudentCtrl) CreateStudent(c *gin.Context) {
	studentDTO := CreateStudentDTO{}
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

	studentService := NewStudentService()
	id := studentService.CreateStudent(studentDTO)
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
	studentDTO := UpdateStudentDTO{}
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
	studentService := NewStudentService()
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
	studentDTO := DeleteStudentDTO{}
	studentDTO.ID = id
	studentDTO.IDUser = c.Value("id_user").(int)
	studentService := NewStudentService()
	id, err = studentService.DeleteStudent(studentDTO)
	if err, ok := err.(utils.AppError); ok {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, struct {
		Id int `json:"id"`
	}{Id: id})
}

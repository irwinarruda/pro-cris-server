package studentsdrivers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type StudentCtrl struct{}

func NewStudentCtrl() *StudentCtrl {
	return proinject.Resolve(&StudentCtrl{})
}

func (s StudentCtrl) GetStudents(c *gin.Context) {
	studentsDTO := students.GetAllStudentsDTO{}
	studentsDTO.IDAccount = c.Value("id_account").(int)
	studentService := students.NewStudentService()
	students := studentService.GetAllStudents(studentsDTO)
	c.JSON(http.StatusOK, students)
}

func (s StudentCtrl) GetStudent(c *gin.Context) {
	idStudent, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleHttpError(c, utils.NewAppError("Invalid student ID.", false, http.StatusBadRequest))
		return
	}
	studentDTO := students.GetStudentDTO{}
	studentDTO.IDAccount = c.Value("id_account").(int)
	studentDTO.ID = idStudent
	studentService := students.NewStudentService()
	student, err := studentService.GetStudentByID(studentDTO)
	if utils.HandleHttpError(c, err) {
		return
	}
	c.JSON(http.StatusOK, student)
}

func (s StudentCtrl) CreateStudent(c *gin.Context) {
	studentDTO := students.CreateStudentDTO{}
	studentDTO.IDAccount = c.Value("id_account").(int)
	err := c.Bind(&studentDTO)
	if err != nil {
		utils.HandleHttpError(c, utils.NewAppError("Invalid student ID.", false, http.StatusBadRequest))
		return
	}

	studentService := students.NewStudentService()
	id, err := studentService.CreateStudent(studentDTO)
	if utils.HandleHttpError(c, err) {
		return
	}
	c.JSON(http.StatusCreated, struct {
		Id int `json:"id"`
	}{Id: id})
}

func (s StudentCtrl) UpdateSudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleHttpError(c, utils.NewAppError("Invalid student ID.", false, http.StatusBadRequest))
		return
	}
	studentDTO := students.UpdateStudentDTO{}
	studentDTO.ID = id
	studentDTO.IDAccount = c.Value("id_account").(int)
	err = c.Bind(&studentDTO)
	if err != nil {
		utils.HandleHttpError(c, utils.NewAppError("Invalid student ID."+err.Error(), false, http.StatusBadRequest))
		return
	}
	studentService := students.NewStudentService()
	id, err = studentService.UpdateStudent(studentDTO)
	if utils.HandleHttpError(c, err) {
		return
	}
	c.JSON(http.StatusOK, struct {
		Id int `json:"id"`
	}{Id: id})
}

func (s StudentCtrl) DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HandleHttpError(c, utils.NewAppError("Invalid student ID.", false, http.StatusBadRequest))
		return
	}
	studentDTO := students.DeleteStudentDTO{}
	studentDTO.ID = id
	studentDTO.IDAccount = c.Value("id_account").(int)
	studentService := students.NewStudentService()
	id, err = studentService.DeleteStudent(studentDTO)
	if utils.HandleHttpError(c, err) {
		return
	}
	c.JSON(http.StatusOK, struct {
		Id int `json:"id"`
	}{Id: id})
}

package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"github.com/irwinarruda/pro-cris-server/tests"
	"github.com/stretchr/testify/assert"
)

func TestStudentServiceHappyPath(t *testing.T) {
	setupTestsStudents()

	var assert = assert.New(t)
	var studentService = students.NewStudentService()

	id1 := studentService.CreateStudent(students.CreateStudentDTO{
		Name:              "John Doe",
		BirthDay:          utils.StringPointer("1990-01-01"),
		DisplayColor:      "#000000",
		Picture:           utils.StringPointer("http://example.com/picture.jpg"),
		ParentName:        utils.StringPointer("Jane Doe"),
		ParentPhoneNumber: utils.StringPointer("1234567890"),
		HouseAddress:      utils.StringPointer("123 Main St"),
		HouseIdentifier:   utils.StringPointer("Apt 1"),
		HouseCoordinate:   &models.Coordinate{Latitude: 20, Longitude: 20},
		BasePrice:         120,
		Routine: []students.CreateStudentRoutinePlanDTO{
			{WeekDay: models.Monday, Duration: 60, StartHour: 8, Price: nil},
			{WeekDay: models.Tuesday, Duration: 60, StartHour: 8, Price: utils.Float64Pointer(100)},
		},
	})
	student1, err := studentService.GetStudentByID(id1)
	assert.NoError(err, "Should return a student with the same ID as the one created\n%v", tests.ExpectString(id1, student1.ID))
	assert.Len(student1.Routine, 2, "Student should have 2 routine plans")
	mondayRoutineId := 0
	for _, routinePlan := range student1.Routine {
		if routinePlan.WeekDay == models.Monday {
			mondayRoutineId = routinePlan.ID
		}
		assert.NotEqual(routinePlan.ID, 0, "Routine plan ID should not be 0")
		assert.Condition(
			func() bool { return routinePlan.WeekDay != models.Monday || routinePlan.Price == 120 },
			"Routine plan price on Monday should be equal to the base price",
		)
		assert.Condition(
			func() bool { return routinePlan.WeekDay != models.Tuesday || routinePlan.Price == 100 },
			"Routine plan price on Tuesday should be equal to 100",
		)
	}

	id2, err := studentService.UpdateStudent(students.UpdateStudentDTO{
		ID:                id1,
		Name:              "Jane Doe Updated",
		BirthDay:          utils.StringPointer("1990-01-02"),
		DisplayColor:      "#FFFFFF",
		Picture:           utils.StringPointer("http://example.com/picture2.jpg"),
		ParentName:        utils.StringPointer("John Doe"),
		ParentPhoneNumber: utils.StringPointer("0987654321"),
		HouseAddress:      utils.StringPointer("456 Main St"),
		HouseIdentifier:   utils.StringPointer("Apt 2"),
		HouseCoordinate:   &models.Coordinate{Latitude: 30, Longitude: 30},
		BasePrice:         200,
		Routine: []students.UpdateStudentRoutinePlanDTO{
			{ID: utils.IntPointer(mondayRoutineId)},
			{ID: nil, WeekDay: utils.StringPointer(models.Friday), StartHour: utils.IntPointer(9), Duration: utils.IntPointer(90)},
		},
	})

	assert.NoError(err, "Should be able to update student")
	assert.Equal(id1, id2, "Should return the same ID as the one updated")

	student2, err := studentService.GetStudentByID(id2)
	assert.NoError(err, "Should return a student with the same ID as the one created\n%v", tests.ExpectString(id2, student2.ID))
	assert.Equal("Jane Doe Updated", student2.Name, "Name should be updated")
	assert.Equal("1990-01-02", *student2.BirthDay, "BirthDay should be updated")
	assert.Equal("#FFFFFF", student2.DisplayColor, "DisplayColor should be updated")
	assert.Equal("http://example.com/picture2.jpg", *student2.Picture, "Picture should be updated")
	assert.Equal("John Doe", *student2.ParentName, "ParentName should be updated")
	assert.Equal("0987654321", *student2.ParentPhoneNumber, "ParentPhoneNumber should be updated")
	assert.Equal("456 Main St", *student2.HouseAddress, "HouseAddress should be updated")
	assert.Equal("Apt 2", *student2.HouseIdentifier, "HouseIdentifier should be updated")
	assert.Equal(30.0, student2.HouseCoordinate.Latitude, "HouseCoordinate Latitude should be updated")
	assert.Equal(30.0, student2.HouseCoordinate.Longitude, "HouseCoordinate Longitude should be updated")
	assert.Equal(200.0, student2.BasePrice, "BasePrice should be updated")
	assert.Len(student2.Routine, 2, "Student should have 2 routine plans")
	for _, routinePlan := range student2.Routine {
		assert.NotEqual(routinePlan.ID, 0, "Routine plan ID should not be 0")
		assert.NotEqual(routinePlan.WeekDay, models.Tuesday, "Routine plan on Tuesday should be deleted")
		assert.Condition(
			func() bool { return routinePlan.WeekDay != models.Monday || routinePlan.Price == 120 },
			"Routine plan price on Monday should be equal to the base price",
		)
		assert.Condition(
			func() bool { return routinePlan.WeekDay != models.Friday || routinePlan.Price == 200 },
			"Routine plan price on Friday should be equal to 200",
		)
	}

	allStudents := studentService.GetAllStudents()
	assert.Len(allStudents, 1, "Should return a list of students with 1 student after creating/updating")

	studentService.DeleteStudent(id2)

	allStudents = studentService.GetAllStudents()
	assert.Len(allStudents, 0, "Should return an empty list of students after deleting")
}

func TestStudentServiceErrorPath(t *testing.T) {
	setupTestsStudents()
	var assert = assert.New(t)
	var studentService = students.NewStudentService()

	_, err := studentService.UpdateStudent(students.UpdateStudentDTO{ID: 7})
	assert.Error(err, "Should return an error when trying to update a student that does not exist")

	_, err = studentService.GetStudentByID(1)
	assert.Error(err, "Should return an error when trying to get a student that does not exist")

	_, err = studentService.DeleteStudent(0)
	assert.Error(err, "Should return an error when trying to delete a student that does not exist")
}

func setupTestsStudents() {
	configs.RegisterInject("env", configs.GetEnv("../../.env"))
	configs.RegisterInject("db", configs.GetDb())
	var studentRepository = configs.ResolveInject(&students.StudentRepository{})
	configs.RegisterInject("students_repository", studentRepository)
	studentRepository.ResetStudents()
}

package integration

import (
	"testing"
	"time"

	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"github.com/irwinarruda/pro-cris-server/tests"
	"github.com/stretchr/testify/assert"
)

type FakeStudentRepository struct {
	Students      []students.Student
	IDStudents    int
	IDRoutinePlan int
}

func (f *FakeStudentRepository) GetAllStudents() []students.Student {
	return f.Students
}

func (f *FakeStudentRepository) GetStudentByID(id int) (students.Student, error) {
	for _, student := range f.Students {
		if student.ID == id {
			return student, nil
		}
	}
	return students.Student{}, utils.NewAppError("Student not found.", true, nil)
}

func (f *FakeStudentRepository) CreateStudent(studentDTO students.CreateStudentDTO) int {
	f.IDStudents++
	student := students.Student{
		ID:                f.IDStudents,
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
	for _, routinePlan := range studentDTO.Routine {
		f.IDRoutinePlan++
		student.Routine = append(student.Routine, students.RoutinePlan{
			ID:        f.IDRoutinePlan,
			StartHour: routinePlan.StartHour,
			Duration:  routinePlan.Duration,
			WeekDay:   routinePlan.WeekDay,
			Price:     *routinePlan.Price,
			CreatedAt: time.Now(),
			IsDeleted: false,
		})
	}
	f.Students = append(f.Students, student)
	return f.IDStudents
}

func (f *FakeStudentRepository) UpdateStudent(studentDTO students.UpdateStudentDTO) (int, error) {
	for i := range f.Students {
		student := &f.Students[i]
		if student.ID == studentDTO.ID {
			student.Name = studentDTO.Name
			student.BirthDay = studentDTO.BirthDay
			student.DisplayColor = studentDTO.DisplayColor
			student.Picture = studentDTO.Picture
			student.ParentName = studentDTO.ParentName
			student.ParentPhoneNumber = studentDTO.ParentPhoneNumber
			student.HouseAddress = studentDTO.HouseAddress
			student.HouseIdentifier = studentDTO.HouseIdentifier
			student.HouseCoordinate = studentDTO.HouseCoordinate
			student.BasePrice = studentDTO.BasePrice
			student.UpdatedAt = time.Now()
			return student.ID, nil
		}
	}
	return 0, utils.NewAppError("Student not found.", true, nil)
}

func (f *FakeStudentRepository) DeleteStudent(id int) (int, error) {
	return 0, nil
}

func (f *FakeStudentRepository) GetRoutineID(idStudent int, excluded ...int) []int {
	routine := []int{}
	for _, student := range f.Students {
		if student.ID == idStudent {
			for _, routinePlan := range student.Routine {
				if !utils.Includes(excluded, routinePlan.ID) {
					routine = append(routine, routinePlan.ID)
				}
			}
		}
	}
	return routine
}

func (f *FakeStudentRepository) CreateRoutine(idStudent int, routinePlan ...students.CreateStudentRoutinePlanDTO) {
	for _, student := range f.Students {
		if student.ID == idStudent {
			for _, r := range routinePlan {
				f.IDRoutinePlan++
				student.Routine = append(student.Routine, students.RoutinePlan{
					ID:        f.IDRoutinePlan,
					StartHour: r.StartHour,
					Duration:  r.Duration,
					WeekDay:   r.WeekDay,
					Price:     *r.Price,
					CreatedAt: time.Now(),
					IsDeleted: false,
				})
			}
		}
	}
}

func (f *FakeStudentRepository) DeleteRoutine(idStudent int, routine ...int) {
	for _, student := range f.Students {
		if student.ID == idStudent {
			for i := 0; i < len(student.Routine); i++ {
				routinePlan := student.Routine[i]
				if utils.Includes(routine, routinePlan.ID) {
					student.Routine = append(student.Routine[:i], student.Routine[i+1:]...)
					i--
				}
			}
		}
	}
}

func TestStudentService(t *testing.T) {
	var assert = assert.New(t)
	configs.GetEnv("../../.env")
	configs.RegisterInject(
		"students_repository",
		&FakeStudentRepository{
			Students:      []students.Student{},
			IDStudents:    0,
			IDRoutinePlan: 0,
		},
	)

	studentsService := students.NewStudentService()
	id1 := studentsService.CreateStudent(students.CreateStudentDTO{
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
		Routine: utils.Map(make([]students.CreateStudentRoutinePlanDTO, 2), func(_ students.CreateStudentRoutinePlanDTO, i int) students.CreateStudentRoutinePlanDTO {
			var weekday models.WeekDay = models.Monday
			var price *float64 = nil
			if i == 1 {
				weekday = models.Tuesday
				price = utils.Float64Pointer(100)
			}
			return students.CreateStudentRoutinePlanDTO{
				WeekDay:   weekday,
				Duration:  60,
				StartHour: 8,
				Price:     price,
			}
		}),
	})
	// Testing the fake repository
	student1, err := studentsService.GetStudentByID(id1)
	assert.NoError(err, "Should return a student with the same ID as the one created\n%v", tests.ExpectString(id1, student1.ID))
	assert.Len(student1.Routine, 2, "Student should have 2 routine plans")
	for _, routinePlan := range student1.Routine {
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

	// Testing update logic
	id2, err := studentsService.UpdateStudent(students.UpdateStudentDTO{
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
			{ID: utils.IntPointer(1)},
			{
				ID:        nil,
				WeekDay:   utils.StringPointer(models.Friday),
				StartHour: utils.IntPointer(9),
				Duration:  utils.IntPointer(90),
			},
		},
	})

	assert.NoError(err, "Should be able to update student")
	assert.Equal(id1, id2, "Should return the same ID as the one updated")

	student2, err := studentsService.GetStudentByID(id2)
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
		assert.Condition(
			func() bool { return routinePlan.WeekDay != models.Monday || routinePlan.Price == 120 },
			"Routine plan price on Monday should be equal to the base price",
		)
		assert.Condition(
			func() bool { return routinePlan.WeekDay != models.Friday || routinePlan.Price == 200 },
			"Routine plan price on Friday should be equal to 200",
		)
	}
}

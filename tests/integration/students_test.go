package integration

import (
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/modules/auth"
	"github.com/irwinarruda/pro-cris-server/modules/students"
	"github.com/irwinarruda/pro-cris-server/shared/constants"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
	"github.com/stretchr/testify/assert"
)

func TestStudentService(t *testing.T) {
	Init()
	var assert = assert.New(t)
	var studentService = students.NewStudentService()

	t.Run("Happy Path", func(t *testing.T) {
		idAccount := beforeEachStudents()

		assert.NotEqual(idAccount, 0, "Should return a valid account id.")

		createStudentDTO := mockCreateStudentDTO(idAccount)
		id1, _ := studentService.CreateStudent(createStudentDTO)

		student1, err := studentService.GetStudentByID(students.GetStudentDTO{IDAccount: idAccount, ID: id1})
		assert.NoError(err, "Should return a student with the same ID as the one created")
		assert.Len(student1.Routine, 2, "Student should have 2 routine plans")
		assert.Equal(models.Male, *student1.Gender, "Should have Male gender")
		assert.Equal(students.PaymentStyleUpfront, student1.PaymentStyle, "Should have Upfront payment style")
		assert.Equal(students.PaymentTypeFixed, student1.PaymentType, "Should have Fixed payment type")
		assert.Equal(float64(2000), *student1.PaymentTypeValue, "Should have 2000 as payment type value")
		assert.Equal(students.SettlementStyleAppointments, student1.SettlementStyle, "Should have Appointments settlement style")
		assert.Equal(10, *student1.SettlementStyleValue, "Should have 10 appointments threshold")
		assert.Nil(student1.SettlementStyleDay, "Should have no settlement day")
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

		updateStudentDTO := mockUpdateStudentDTO(idAccount, id1)
		updateStudentDTO.Routine = append(updateStudentDTO.Routine, students.UpdateStudentRoutinePlanDTO{ID: utils.ToP(mondayRoutineId)})
		id2, err := studentService.UpdateStudent(updateStudentDTO)

		assert.NoError(err, "Should be able to update student")
		assert.Equal(id1, id2, "Should return the same ID as the one updated")

		student2, err := studentService.GetStudentByID(students.GetStudentDTO{IDAccount: idAccount, ID: id2})
		assert.NotEqual(student2.CreatedAt, student2.UpdatedAt, "UpdatedAt should be updated")
		assert.NoError(err, "Should return a student with the same ID aS the one created")
		assert.Equal("Jane Doe Updated", student2.Name, "Name should be updated")
		assert.Equal("1990-01-02", *student2.BirthDay, "BirthDay should be updated")
		assert.Equal("#FFFFFF", student2.DisplayColor, "DisplayColor should be updated")
		assert.Equal("http://example.com/picture2.jpg", *student2.Picture, "Picture should be updated")
		assert.Equal(models.Female, *student2.Gender, "Should have Female gender")
		assert.Equal("John Doe", *student2.ParentName, "ParentName should be updated")
		assert.Equal("0987654321", *student2.ParentPhoneNumber, "ParentPhoneNumber should be updated")
		assert.Equal(students.PaymentStyleLater, student2.PaymentStyle, "Should have Later payment style")
		assert.Equal(students.PaymentTypeVariable, student2.PaymentType, "Should have Variable payment type")
		assert.Nil(student2.PaymentTypeValue, "Should have nil as payment type value")
		assert.Equal(students.SettlementStyleMonthly, student2.SettlementStyle, "Should have Monthly settlement style")
		assert.Equal(1, *student2.SettlementStyleValue, "Should have 10 month threshold")
		assert.Equal(5, *student2.SettlementStyleDay, "Should have day 5th as the settlement day")
		assert.Equal("456 Main St", *student2.HouseAddress, "HouseAddress should be updated")
		assert.Equal("Apt 2", *student2.HouseIdentifier, "HouseIdentifier should be updated")
		assert.Equal(30.0, student2.HouseCoordinate.Latitude, "HouseCoordinate Latitude should be updated")
		assert.Equal(30.0, student2.HouseCoordinate.Longitude, "HouseCoordinate Longitude should be updated")
		assert.Len(student2.Routine, 2, "Student should have 2 routine plans")
		for _, routinePlan := range student2.Routine {
			assert.NotEqual(0, routinePlan.ID, "Routine plan ID should not be 0")
			assert.NotEqual(models.Tuesday, routinePlan.WeekDay, "Routine plan on Tuesday should be deleted")
			assert.Condition(
				func() bool { return routinePlan.WeekDay != models.Monday || routinePlan.Price == 120 },
				"Routine plan price on Monday should be equal to the base price",
			)
			assert.Condition(
				func() bool { return routinePlan.WeekDay != models.Friday || routinePlan.Price == 200 },
				"Routine plan price on Friday should be equal to 200",
			)
		}

		allStudents, _ := studentService.GetAllStudents(students.GetAllStudentsDTO{IDAccount: idAccount})
		assert.Len(allStudents, 1, "Should return a list of students with 1 student after creating/updating")

		studentService.DeleteStudent(students.DeleteStudentDTO{ID: id2, IDAccount: idAccount})

		allStudents, _ = studentService.GetAllStudents(students.GetAllStudentsDTO{IDAccount: idAccount})
		assert.Len(allStudents, 0, "Should return an empty list of students after deleting")

		afterEachStudents()
	})

	t.Run("Error Path", func(t *testing.T) {
		idAccount := beforeEachStudents()

		createStudentDTO := mockCreateStudentDTO(idAccount)
		createStudentDTO.PaymentType = students.PaymentTypeFixed
		createStudentDTO.PaymentTypeValue = nil
		_, err := studentService.CreateStudent(createStudentDTO)
		assert.Error(err, "Should return an error when payment type is Fixed and value is nil")

		createStudentDTO.PaymentType = students.PaymentTypeVariable
		createStudentDTO.SettlementStyle = students.SettlementStyleMonthly
		createStudentDTO.SettlementStyleValue = nil
		_, err = studentService.CreateStudent(createStudentDTO)
		assert.Error(err, "Should return an error when settlement type is Monthly or Weekly and value is nil")

		createStudentDTO.SettlementStyleValue = utils.ToP(1)
		createStudentDTO.SettlementStyleDay = nil
		_, err = studentService.CreateStudent(createStudentDTO)
		assert.Error(err, "Should return an error when settlement type is Monthly or Weekly and day is nil")

		createStudentDTO.SettlementStyle = students.SettlementStyleAppointments
		createStudentDTO.SettlementStyleValue = nil
		createStudentDTO.SettlementStyleDay = nil
		_, err = studentService.CreateStudent(createStudentDTO)
		assert.NoError(err, "Should not return an error when settlement type is Appointments and value/day is nil")

		updateStudentDTO := mockUpdateStudentDTO(idAccount, 1)
		updateStudentDTO.PaymentType = students.PaymentTypeFixed
		updateStudentDTO.PaymentTypeValue = nil
		_, err = studentService.UpdateStudent(updateStudentDTO)
		assert.Error(err, "Should return an error when payment type is Fixed and payment type value is nil")

		_, err = studentService.UpdateStudent(students.UpdateStudentDTO{IDAccount: idAccount, ID: 5})
		assert.Error(err, "Should return an error when trying to update a student that does not exist")

		_, err = studentService.GetStudentByID(students.GetStudentDTO{IDAccount: idAccount, ID: 3})
		assert.Error(err, "Should return an error when trying to get a student that does not exist")

		_, err = studentService.DeleteStudent(students.DeleteStudentDTO{IDAccount: idAccount, ID: 7})
		assert.Error(err, "Should return an error when trying to delete a student that does not exist")

		afterEachStudents()
	})

	t.Run("Overlapping routine plan", func(t *testing.T) {
		idAccount := beforeEachStudents()

		createStudentDTO := mockCreateStudentDTO(idAccount)
		createStudentDTO.Routine = []students.CreateStudentRoutinePlanDTO{
			{WeekDay: models.Thursday, Duration: 60, StartHour: 8, Price: 120},
			{WeekDay: models.Monday, Duration: 60, StartHour: 8, Price: 120},
			{WeekDay: models.Monday, Duration: 60, StartHour: 68, Price: 120},
			{WeekDay: models.Monday, Duration: 4, StartHour: 4, Price: 120},
		}
		idStudent, err := studentService.CreateStudent(createStudentDTO)
		assert.NoError(err, "Should not return when overlapping happens at the beggining and the end appointment hours")

		createStudentDTO.Routine = []students.CreateStudentRoutinePlanDTO{
			{WeekDay: models.Thursday, Duration: 60, StartHour: 8, Price: 120},
			{WeekDay: models.Monday, Duration: 60, StartHour: 8, Price: 120},
			{WeekDay: models.Monday, Duration: 60, StartHour: 67, Price: 120},
		}
		_, err = studentService.CreateStudent(createStudentDTO)
		assert.Error(err, "Should return an error when overlapping routine plan happens")

		createStudentDTO.Routine = []students.CreateStudentRoutinePlanDTO{
			{WeekDay: models.Saturday, Duration: 60, StartHour: 7, Price: 120},
			{WeekDay: models.Saturday.Before(), Duration: constants.Hour24, StartHour: 8, Price: 120},
		}
		_, err = studentService.CreateStudent(createStudentDTO)
		assert.Error(err, "Should return an error when overlapping routine plan from day before happens")

		createStudentDTO.Routine = []students.CreateStudentRoutinePlanDTO{
			{WeekDay: models.Sunday, Duration: constants.Hour24, StartHour: 8, Price: 120},
			{WeekDay: models.Sunday.After(), Duration: 60, StartHour: 7, Price: 120},
		}
		_, err = studentService.CreateStudent(createStudentDTO)
		assert.Error(err, "Should return an error when overlapping routine plan from day after happens")

		createStudentDTO.Routine = []students.CreateStudentRoutinePlanDTO{
			{WeekDay: models.Thursday, Duration: constants.Hour24 + 1, StartHour: 8, Price: 120},
		}
		_, err = studentService.CreateStudent(createStudentDTO)
		assert.Error(err, "Should return an error when duration is greater than 24 hours")

		createStudentDTO.Routine = []students.CreateStudentRoutinePlanDTO{
			{WeekDay: models.Thursday, Duration: 60, StartHour: constants.Hour24 + 1, Price: 120},
		}
		_, err = studentService.CreateStudent(createStudentDTO)
		assert.Error(err, "Should return an error when start hour is greater than 24 hours")

		updateStudentDTO := mockUpdateStudentDTO(idAccount, idStudent)
		updateStudentDTO.Routine = []students.UpdateStudentRoutinePlanDTO{
			{ID: utils.ToP(1)},
			{ID: nil, WeekDay: utils.ToP(models.Thursday), StartHour: utils.ToP(0), Duration: utils.ToP(8), Price: utils.ToP(200.0)},
		}
		_, err = studentService.UpdateStudent(updateStudentDTO)
		assert.NoError(err, "Should not return when updating a routine plan that is not overlapping")

		updateStudentDTO.Routine = []students.UpdateStudentRoutinePlanDTO{
			{ID: utils.ToP(1)},
			{ID: nil, WeekDay: utils.ToP(models.Thursday), StartHour: utils.ToP(50), Duration: utils.ToP(12), Price: utils.ToP(200.0)},
		}
		_, err = studentService.UpdateStudent(updateStudentDTO)
		assert.Error(err, "Should return an error when overlapping happens")

		updateStudentDTO.Routine = []students.UpdateStudentRoutinePlanDTO{
			{ID: utils.ToP(1)},
			{ID: nil, WeekDay: utils.ToP(models.Thursday.Before()), StartHour: utils.ToP(9), Duration: utils.ToP(constants.Hour24), Price: utils.ToP(200.0)},
		}
		_, err = studentService.UpdateStudent(updateStudentDTO)
		assert.Error(err, "Should return an error when overlapping routine plan from day before happens")

		updateStudentDTO.Routine = []students.UpdateStudentRoutinePlanDTO{
			{ID: nil, WeekDay: utils.ToP(models.Thursday), StartHour: utils.ToP(8), Duration: utils.ToP(constants.Hour24 + 1), Price: utils.ToP(200.0)},
		}
		_, err = studentService.UpdateStudent(updateStudentDTO)
		assert.Error(err, "Should return an error when duration is greater than 24 hours")

		updateStudentDTO.Routine = []students.UpdateStudentRoutinePlanDTO{
			{ID: nil, WeekDay: utils.ToP(models.Thursday), StartHour: utils.ToP(constants.Hour24 + 1), Duration: utils.ToP(60), Price: utils.ToP(200.0)},
		}
		_, err = studentService.UpdateStudent(updateStudentDTO)
		assert.Error(err, "Should return an error when start hour is greater than 24 hours")

		afterEachStudents()
	})
}

func mockCreateStudentDTO(idAccount int) students.CreateStudentDTO {
	return students.CreateStudentDTO{
		IDAccount:            idAccount,
		Name:                 "John Doe",
		Gender:               utils.ToP(models.Male),
		BirthDay:             utils.ToP("1990-01-01"),
		DisplayColor:         "#000000",
		Picture:              utils.ToP("http://example.com/picture.jpg"),
		ParentName:           utils.ToP("Jane Doe"),
		ParentPhoneNumber:    utils.ToP("1234567890"),
		PaymentStyle:         students.PaymentStyleUpfront,
		PaymentType:          students.PaymentTypeFixed,
		PaymentTypeValue:     utils.ToP(2000.0),
		SettlementStyle:      students.SettlementStyleAppointments,
		SettlementStyleValue: utils.ToP(10),
		SettlementStyleDay:   nil,
		HouseAddress:         utils.ToP("123 Main St"),
		HouseIdentifier:      utils.ToP("Apt 1"),
		HouseCoordinate:      &models.Coordinate{Latitude: 20, Longitude: 20},
		Routine: []students.CreateStudentRoutinePlanDTO{
			{WeekDay: models.Monday, Duration: 60, StartHour: 8, Price: 120},
			{WeekDay: models.Tuesday, Duration: 60, StartHour: 8, Price: 100},
		},
	}
}

func mockCreateStudentDTO2(idAccount int) students.CreateStudentDTO {
	return students.CreateStudentDTO{
		IDAccount:            idAccount,
		Name:                 "Jane Doe Updated",
		Gender:               utils.ToP(models.Female),
		BirthDay:             utils.ToP("1990-01-02"),
		DisplayColor:         "#FFFFFF",
		Picture:              utils.ToP("http://example.com/picture2.jpg"),
		ParentName:           utils.ToP("John Doe"),
		ParentPhoneNumber:    utils.ToP("0987654321"),
		PaymentStyle:         students.PaymentStyleLater,
		PaymentType:          students.PaymentTypeVariable,
		PaymentTypeValue:     nil,
		SettlementStyle:      students.SettlementStyleMonthly,
		SettlementStyleValue: utils.ToP(1),
		SettlementStyleDay:   utils.ToP(5),
		HouseAddress:         utils.ToP("456 Main St"),
		HouseIdentifier:      utils.ToP("Apt 2"),
		HouseCoordinate:      &models.Coordinate{Latitude: 30, Longitude: 30},
		Routine: []students.CreateStudentRoutinePlanDTO{
			{WeekDay: models.Tuesday, Duration: 120, StartHour: 68, Price: 300},
		},
	}
}

func mockUpdateStudentDTO(idAccount, id int) students.UpdateStudentDTO {
	return students.UpdateStudentDTO{
		IDAccount:            idAccount,
		ID:                   id,
		Name:                 "Jane Doe Updated",
		Gender:               utils.ToP(models.Female),
		BirthDay:             utils.ToP("1990-01-02"),
		DisplayColor:         "#FFFFFF",
		Picture:              utils.ToP("http://example.com/picture2.jpg"),
		ParentName:           utils.ToP("John Doe"),
		ParentPhoneNumber:    utils.ToP("0987654321"),
		PaymentStyle:         students.PaymentStyleLater,
		PaymentType:          students.PaymentTypeVariable,
		PaymentTypeValue:     nil,
		SettlementStyle:      students.SettlementStyleMonthly,
		SettlementStyleValue: utils.ToP(1),
		SettlementStyleDay:   utils.ToP(5),
		HouseAddress:         utils.ToP("456 Main St"),
		HouseIdentifier:      utils.ToP("Apt 2"),
		HouseCoordinate:      &models.Coordinate{Latitude: 30, Longitude: 30},
		Routine: []students.UpdateStudentRoutinePlanDTO{
			{ID: nil, WeekDay: utils.ToP(models.Friday), StartHour: utils.ToP(9), Duration: utils.ToP(90), Price: utils.ToP(200.0)},
		},
	}
}

func beforeEachStudents() int {
	var authRepository = proinject.Get[auth.IAuthRepository]("auth_repository")
	var studentRepository = proinject.Get[students.IStudentRepository]("student_repository")
	authRepository.ResetAuth()
	studentRepository.ResetStudents()

	account, _ := authRepository.CreateAccount(auth.CreateAccountDTO{
		Email:         "john@doe.com",
		Name:          "John Doe",
		Picture:       utils.ToP("https://www.google.com"),
		EmailVerified: false,
		Provider:      auth.LoginProviderGoogle,
	})
	return account.ID
}

func afterEachStudents() {
	var authRepository = proinject.Get[auth.IAuthRepository]("auth_repository")
	var studentRepository = proinject.Get[students.IStudentRepository]("student_repository")
	authRepository.ResetAuth()
	studentRepository.ResetStudents()
}

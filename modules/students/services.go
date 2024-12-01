package students

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/constants"
	"github.com/irwinarruda/pro-cris-server/shared/models"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type StudentService struct {
	Validate           configs.Validate   `inject:"validate"`
	StudentsRepository IStudentRepository `inject:"student_repository"`
}

type IStudentService = *StudentService

func NewStudentService() *StudentService {
	return proinject.Resolve(&StudentService{})
}

func (s *StudentService) GetAllStudents(data GetAllStudentsDTO) ([]Student, error) {
	if err := s.Validate.Struct(data); err != nil {
		return []Student{}, err
	}
	return s.StudentsRepository.GetAllStudents(data)
}

func (s *StudentService) GetStudentByID(data GetStudentDTO) (Student, error) {
	if err := s.Validate.Struct(data); err != nil {
		return Student{}, err
	}
	return s.StudentsRepository.GetStudentByID(data)
}

func (s *StudentService) CreateStudent(student CreateStudentDTO) (int, error) {
	if err := s.Validate.Struct(student); err != nil {
		return 0, err
	}
	if student.PaymentType == models.PaymentTypeFixed && student.PaymentTypeValue == nil {
		return 0, utils.NewAppError("Payment type value is required when payment type is Fixed.", true, http.StatusBadRequest)
	}
	if student.PaymentType != models.PaymentTypeFixed {
		student.PaymentTypeValue = nil
	}
	if student.SettlementStyle != models.SettlementStyleAppointments && student.SettlementStyleValue == nil {
		return 0, utils.NewAppError("Settlement value is required when settlement style is not Appointments.", true, http.StatusBadRequest)
	}
	if student.SettlementStyle != models.SettlementStyleAppointments && student.SettlementStyleDay == nil {
		return 0, utils.NewAppError("Settlement day is required when settlement style is not Appointments.", true, http.StatusBadRequest)
	}
	if student.SettlementStyle != models.SettlementStyleAppointments {
		student.PaymentTypeValue = nil
	}
	for i, routineI := range student.Routine {
		for j, routineJ := range student.Routine {
			if i == j {
				continue
			}
			if routineJ.WeekDay == routineI.WeekDay && utils.IsOverlapping(routineI.StartHour, routineI.Duration, routineJ.StartHour, routineJ.Duration) {
				return 0, utils.NewAppErrors(fmt.Sprintf("Routine plans at %s are overlapping.", routineI.WeekDay), []CreateStudentRoutinePlanDTO{routineI, routineJ}, true, http.StatusBadRequest)
			}
			if routineJ.WeekDay == routineI.WeekDay.Before() && utils.IsOverlapping(routineI.StartHour+constants.Hour24, routineI.Duration, routineJ.StartHour, routineJ.Duration) {
				return 0, utils.NewAppErrors(fmt.Sprintf("Routine plans at %s are overlapping.", routineI.WeekDay), []CreateStudentRoutinePlanDTO{routineI, routineJ}, true, http.StatusBadRequest)
			}
			if routineJ.WeekDay == routineI.WeekDay.After() && utils.IsOverlapping(routineI.StartHour, routineI.Duration, routineJ.StartHour+constants.Hour24, routineJ.Duration) {
				return 0, utils.NewAppErrors(fmt.Sprintf("Routine plans at %s are overlapping.", routineI.WeekDay), []CreateStudentRoutinePlanDTO{routineI, routineJ}, true, http.StatusBadRequest)
			}
		}
	}
	return s.StudentsRepository.CreateStudent(student), nil
}

func (s *StudentService) UpdateStudent(student UpdateStudentDTO) (int, error) {
	if err := s.Validate.Struct(student); err != nil {
		return 0, err
	}
	if student.PaymentType == models.PaymentTypeFixed && student.PaymentTypeValue == nil {
		return 0, utils.NewAppError("Payment type value is required when payment type is Fixed.", true, http.StatusBadRequest)
	}
	if student.PaymentType != models.PaymentTypeFixed {
		student.PaymentTypeValue = nil
	}
	if student.SettlementStyle != models.SettlementStyleAppointments && student.SettlementStyleValue == nil {
		return 0, utils.NewAppError("Settlement value is required when settlement style is not Appointments.", true, http.StatusBadRequest)
	}
	if student.SettlementStyle != models.SettlementStyleAppointments && student.SettlementStyleDay == nil {
		return 0, utils.NewAppError("Settlement day is required when settlement style is not Appointments.", true, http.StatusBadRequest)
	}
	if student.SettlementStyle != models.SettlementStyleAppointments {
		student.PaymentTypeValue = nil
	}
	idStudent, err := s.StudentsRepository.UpdateStudent(student)
	if err != nil {
		return 0, err
	}
	updatedStudent, err := s.StudentsRepository.GetStudentByID(GetStudentDTO{IDAccount: student.IDAccount, ID: idStudent})
	if err != nil {
		return 0, err
	}
	resultingRoutinePlan := []RoutinePlan{}
	mustCreateRoutine := []CreateStudentRoutinePlanDTO{}
	existingRoutine := []int{}
	for _, routinePlan := range student.Routine {
		if routinePlan.ID == nil {
			mustCreateRoutine = append(mustCreateRoutine, CreateStudentRoutinePlanDTO{
				WeekDay:   *routinePlan.WeekDay,
				Duration:  *routinePlan.Duration,
				StartHour: *routinePlan.StartHour,
				Price:     *routinePlan.Price,
			})
			resultingRoutinePlan = append(resultingRoutinePlan, RoutinePlan{
				WeekDay:   *routinePlan.WeekDay,
				StartHour: *routinePlan.StartHour,
				Duration:  *routinePlan.Duration,
				Price:     *routinePlan.Price,
			})
			continue
		}
		index := slices.IndexFunc(updatedStudent.Routine, func(r RoutinePlan) bool {
			return r.ID == *routinePlan.ID
		})
		if index != -1 {
			existingRoutine = append(existingRoutine, updatedStudent.Routine[index].ID)
			resultingRoutinePlan = append(resultingRoutinePlan, RoutinePlan{
				ID:        updatedStudent.Routine[index].ID,
				WeekDay:   updatedStudent.Routine[index].WeekDay,
				StartHour: updatedStudent.Routine[index].StartHour,
				Duration:  updatedStudent.Routine[index].Duration,
				Price:     updatedStudent.Routine[index].Price,
			})
		}
	}
	for i, routineI := range resultingRoutinePlan {
		for j, routineJ := range resultingRoutinePlan {
			if i == j {
				continue
			}
			if routineJ.WeekDay == routineI.WeekDay && utils.IsOverlapping(routineI.StartHour, routineI.Duration, routineJ.StartHour, routineJ.Duration) {
				return 0, utils.NewAppErrors(fmt.Sprintf("Routine plans at %s are overlapping.", routineI.WeekDay), []RoutinePlan{routineI, routineJ}, true, http.StatusBadRequest)
			}
			if routineJ.WeekDay == routineI.WeekDay.Before() && utils.IsOverlapping(routineI.StartHour+constants.Hour24, routineI.Duration, routineJ.StartHour, routineJ.Duration) {
				return 0, utils.NewAppErrors(fmt.Sprintf("Routine plans at %s are overlapping.", routineI.WeekDay), []RoutinePlan{routineI, routineJ}, true, http.StatusBadRequest)
			}
			if routineJ.WeekDay == routineI.WeekDay.After() && utils.IsOverlapping(routineI.StartHour, routineI.Duration, routineJ.StartHour+constants.Hour24, routineJ.Duration) {
				return 0, utils.NewAppErrors(fmt.Sprintf("Routine plans at %s are overlapping.", routineI.WeekDay), []RoutinePlan{routineI, routineJ}, true, http.StatusBadRequest)
			}
		}
	}
	if len(existingRoutine) != len(updatedStudent.Routine) {
		s.StudentsRepository.DeleteAllRoutine(idStudent, existingRoutine...)
	}
	if len(mustCreateRoutine) > 0 {
		s.StudentsRepository.CreateRoutine(idStudent, mustCreateRoutine...)
	}
	return idStudent, nil
}

func (s *StudentService) DeleteStudent(data DeleteStudentDTO) (int, error) {
	if err := s.Validate.Struct(data); err != nil {
		return 0, err
	}
	return s.StudentsRepository.DeleteStudent(data)
}

func (s *StudentService) DoesStudentExists(data DoesStudentExistsDTO) bool {
	if err := s.Validate.Struct(data); err != nil {
		return false
	}
	_, err := s.StudentsRepository.GetStudentByID(GetStudentDTO(data))
	return err == nil
}

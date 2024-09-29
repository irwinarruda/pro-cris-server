package students

import (
	"fmt"
	"net/http"

	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/configs"
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

func (s *StudentService) DoesStudentExists(data DoesStudentExistsDTO) bool {
	if err := s.Validate.Struct(data); err != nil {
		return false
	}
	_, err := s.StudentsRepository.GetStudentByID(GetStudentDTO(data))
	return err == nil
}

func (s *StudentService) CreateStudent(student CreateStudentDTO) (int, error) {
	if err := s.Validate.Struct(student); err != nil {
		return 0, err
	}
	if student.PaymentType == PaymentTypeFixed && student.PaymentTypeValue == nil {
		return 0, utils.NewAppError("Payment type value is required when payment type is Fixed.", true, http.StatusBadRequest)
	}
	if student.PaymentType != PaymentTypeFixed {
		student.PaymentTypeValue = nil
	}
	if student.SettlementStyle != SettlementStyleAppointments && student.SettlementStyleValue == nil {
		return 0, utils.NewAppError("Settlement value is required when settlement style is not Appointments.", true, http.StatusBadRequest)
	}
	if student.SettlementStyle != SettlementStyleAppointments && student.SettlementStyleDay == nil {
		return 0, utils.NewAppError("Settlement day is required when settlement style is not Appointments.", true, http.StatusBadRequest)
	}
	if student.SettlementStyle != SettlementStyleAppointments {
		student.PaymentTypeValue = nil
	}
	for i, routineI := range student.Routine {
		for j, routineJ := range student.Routine {
			if routineI.WeekDay != routineJ.WeekDay || i == j {
				continue
			}
			if utils.IsOverlapping(routineI.StartHour, routineI.Duration, routineJ.StartHour, routineJ.Duration) {
				return 0, utils.NewAppError(fmt.Sprintf("Routine plans at %s are overlapping", routineI.WeekDay), true, http.StatusBadRequest)
			}
		}
	}
	return s.StudentsRepository.CreateStudent(student), nil
}

func (s *StudentService) UpdateStudent(student UpdateStudentDTO) (int, error) {
	if err := s.Validate.Struct(student); err != nil {
		return 0, err
	}
	if student.PaymentType == PaymentTypeFixed && student.PaymentTypeValue == nil {
		return 0, utils.NewAppError("Payment type value is required when payment type is Fixed.", true, http.StatusBadRequest)
	}
	if student.PaymentType != PaymentTypeFixed {
		student.PaymentTypeValue = nil
	}
	if student.SettlementStyle != SettlementStyleAppointments && student.SettlementStyleValue == nil {
		return 0, utils.NewAppError("Settlement value is required when settlement style is not Appointments.", true, http.StatusBadRequest)
	}
	if student.SettlementStyle != SettlementStyleAppointments && student.SettlementStyleDay == nil {
		return 0, utils.NewAppError("Settlement day is required when settlement style is not Appointments.", true, http.StatusBadRequest)
	}
	if student.SettlementStyle != SettlementStyleAppointments {
		student.PaymentTypeValue = nil
	}
	idStudent, err := s.StudentsRepository.UpdateStudent(student)
	if err != nil {
		return 0, err
	}
	mustCreateRoutine := []CreateStudentRoutinePlanDTO{}
	existingRoutine := []int{}
	for _, routinePlan := range student.Routine {
		if routinePlan.ID != nil {
			existingRoutine = append(existingRoutine, *routinePlan.ID)
			continue
		}
		mustCreateRoutine = append(mustCreateRoutine, CreateStudentRoutinePlanDTO{
			WeekDay:   *routinePlan.WeekDay,
			Duration:  *routinePlan.Duration,
			StartHour: *routinePlan.StartHour,
			Price:     *routinePlan.Price,
		})
	}
	shouldDeleteRoutine := s.StudentsRepository.GetRoutineID(idStudent, existingRoutine...)
	if len(shouldDeleteRoutine) > 0 {
		s.StudentsRepository.DeleteRoutine(idStudent, shouldDeleteRoutine...)
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

package students

import (
	"github.com/irwinarruda/pro-cris-server/libs/proinject"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type StudentService struct {
	StudentsRepository IStudentRepository `inject:"students_repository"`
}

func NewStudentService() *StudentService {
	return proinject.Resolve(&StudentService{})
}

func (s *StudentService) GetAllStudents(data GetAllStudentsDTO) []Student {
	return s.StudentsRepository.GetAllStudents(data)
}

func (s *StudentService) GetStudentByID(data GetStudentDTO) (Student, error) {
	return s.StudentsRepository.GetStudentByID(data)
}

func (s *StudentService) CreateStudent(student CreateStudentDTO) (int, error) {
	if student.PaymentType == Fixed && student.PaymentTypeValue == nil {
		return 0, utils.NewAppError("Payment type value is required.", true, nil)
	}
	if student.PaymentType != Fixed {
		student.PaymentTypeValue = nil
	}
	if student.SettlementStyle != Appointments && student.SettlementStyleValue == nil {
		return 0, utils.NewAppError("Settlement value is required.", true, nil)
	}
	if student.SettlementStyle != Appointments && student.SettlementStyleDay == nil {
		return 0, utils.NewAppError("Settlement day is required.", true, nil)
	}
	if student.SettlementStyle != Appointments {
		student.PaymentTypeValue = nil
	}
	return s.StudentsRepository.CreateStudent(student), nil
}

func (s *StudentService) UpdateStudent(student UpdateStudentDTO) (int, error) {
	if student.PaymentType == Fixed && student.PaymentTypeValue == nil {
		return 0, utils.NewAppError("Payment type value is required.", true, nil)
	}
	if student.PaymentType != Fixed {
		student.PaymentTypeValue = nil
	}
	if student.SettlementStyle != Appointments && student.SettlementStyleValue == nil {
		return 0, utils.NewAppError("Settlement value is required.", true, nil)
	}
	if student.SettlementStyle != Appointments && student.SettlementStyleDay == nil {
		return 0, utils.NewAppError("Settlement day is required.", true, nil)
	}
	if student.SettlementStyle != Appointments {
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
	return s.StudentsRepository.DeleteStudent(data)
}

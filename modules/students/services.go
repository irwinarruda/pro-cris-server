package students

import (
	"github.com/irwinarruda/pro-cris-server/shared/configs"
	"github.com/irwinarruda/pro-cris-server/shared/utils"
)

type StudentService struct {
	StudentsRepository IStudentRepository `inject:"students_repository"`
}

func NewStudentService() *StudentService {
	return configs.ResolveInject(&StudentService{})
}

func (s *StudentService) GetAllStudents() []Student {
	return s.StudentsRepository.GetAllStudents()
}

func (s *StudentService) GetStudentByID(id int) (Student, error) {
	return s.StudentsRepository.GetStudentByID(id)
}

func (s *StudentService) CreateStudent(student CreateStudentDTO) int {
	student.Routine = utils.Map(student.Routine, func(s CreateStudentRoutinePlanDTO, _ int) CreateStudentRoutinePlanDTO {
		if s.Price == nil {
			s.Price = &student.BasePrice
		}
		return s
	})
	return s.StudentsRepository.CreateStudent(student)
}

func (s *StudentService) UpdateStudent(student UpdateStudentDTO) (int, error) {
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
		price := &student.BasePrice
		if routinePlan.Price != nil {
			price = routinePlan.Price
		}
		mustCreateRoutine = append(mustCreateRoutine, CreateStudentRoutinePlanDTO{
			WeekDay:   *routinePlan.WeekDay,
			Duration:  *routinePlan.Duration,
			StartHour: *routinePlan.StartHour,
			Price:     price,
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

func (s *StudentService) DeleteStudent(id int) (int, error) {
	return s.StudentsRepository.DeleteStudent(id)
}

package students

import "github.com/irwinarruda/pro-cris-server/shared/utils"

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (s *StudentService) GetAllStudents() []Student {
	studentsRepository := newStudentRepository()
	return studentsRepository.GetAllStudents()
}

func (s *StudentService) GetStudentByID(id int) Student {
	studentsRepository := newStudentRepository()
	return studentsRepository.GetStudentByID(id)
}

func (s *StudentService) CreateStudent(student CreateStudentDTO) int {
	studentsRepository := newStudentRepository()
	return studentsRepository.CreateStudent(student)
}

func (s *StudentService) UpdateStudent(student UpdateStudentDTO) int {
	studentsRepository := newStudentRepository()
	id := studentsRepository.UpdateStudent1(student)
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
			Price:     routinePlan.Price,
		})
	}
	shouldDeleteRoutine := studentsRepository.GetRoutineByIDStudent1(id, &existingRoutine)
	if len(shouldDeleteRoutine) > 0 {
		studentsRepository.DeleteRoutinePlanByIDStudent1(id, utils.Map(
			shouldDeleteRoutine,
			func(rp RoutinePlan) int {
				return rp.ID
			}),
		)
	}

	if len(mustCreateRoutine) > 0 {
		studentsRepository.CreateRoutinePlanByIDStudent1(id, mustCreateRoutine)
	}

	return id
}

func (s *StudentService) DeleteStudent(id int) {
	studentsRepository := newStudentRepository()
	studentsRepository.DeleteStudentByID(id)
}

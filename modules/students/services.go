package students

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
	idStudent := studentsRepository.UpdateStudent(student)
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
	shouldDeleteRoutine := studentsRepository.GetRoutineID(idStudent, existingRoutine...)
	if len(shouldDeleteRoutine) > 0 {
		studentsRepository.DeleteRoutinePlan(idStudent, shouldDeleteRoutine...)
	}

	if len(mustCreateRoutine) > 0 {
		studentsRepository.CreateRoutinePlan(idStudent, mustCreateRoutine...)
	}

	return idStudent
}

func (s *StudentService) DeleteStudent(id int) {
	studentsRepository := newStudentRepository()
	studentsRepository.DeleteStudent(id)
}

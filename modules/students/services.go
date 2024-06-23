package students

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (s *StudentService) GetAllStudents() []Student {
	studentsRepository := NewStudentRepository()
	return studentsRepository.GetAllStudents()
}

func (s *StudentService) GetStudentByID(id int) (Student, error) {
	studentsRepository := NewStudentRepository()
	return studentsRepository.GetStudentByID(id)
}

func (s *StudentService) CreateStudent(student CreateStudentDTO) int {
	studentsRepository := NewStudentRepository()
	return studentsRepository.CreateStudent(student)
}

func (s *StudentService) UpdateStudent(student UpdateStudentDTO) (int, error) {
	studentsRepository := NewStudentRepository()
	idStudent, err := studentsRepository.UpdateStudent(student)
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
			Price:     routinePlan.Price,
		})
	}
	shouldDeleteRoutine := studentsRepository.GetRoutineID(idStudent, existingRoutine...)
	if len(shouldDeleteRoutine) > 0 {
		studentsRepository.DeleteRoutine(idStudent, shouldDeleteRoutine...)
	}
	if len(mustCreateRoutine) > 0 {
		studentsRepository.CreateRoutine(idStudent, mustCreateRoutine...)
	}
	return idStudent, nil
}

func (s *StudentService) DeleteStudent(id int) (int, error) {
	studentsRepository := NewStudentRepository()
	return studentsRepository.DeleteStudent(id)
}

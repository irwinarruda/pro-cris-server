package students

import "fmt"

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (s *StudentService) GetAllStudents() []Student {
	studentsRepository := NewStudentRepository()
	return studentsRepository.GetAllStudents()
}

func (s *StudentService) GetStudentByID(id int) Student {
	studentsRepository := NewStudentRepository()
	student, err := studentsRepository.GetStudentByID(id)
	if err != nil {
		fmt.Println(err)
	}
	return student
}

func (s *StudentService) CreateStudent(student CreateStudentDTO) int {
	studentsRepository := NewStudentRepository()
	id := studentsRepository.CreateStudent(student)
	return id
}

func (s *StudentService) UpdateStudent(student UpdateStudentDTO) int {
	studentsRepository := NewStudentRepository()
	idStudent, _ := studentsRepository.UpdateStudent(student)
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
	return idStudent
}

func (s *StudentService) DeleteStudent(id int) {
	studentsRepository := NewStudentRepository()
	studentsRepository.DeleteStudent(id)
}

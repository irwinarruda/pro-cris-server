package students

type IStudentRepository interface {
	GetAllStudents(data GetAllStudentsDTO) ([]Student, error)
	GetStudentByID(data GetStudentDTO) (Student, error)
	CreateStudent(student CreateStudentDTO) int
	UpdateStudent(student UpdateStudentDTO) (int, error)
	DeleteStudent(data DeleteStudentDTO) (int, error)
	GetRoutineID(idStudent int, excluded ...int) []int
	CreateRoutine(idStudent int, routinePlan ...CreateStudentRoutinePlanDTO)
	DeleteAllRoutine(idStudent int, except ...int)
	ResetStudents()
}

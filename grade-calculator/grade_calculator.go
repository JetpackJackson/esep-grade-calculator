package esepunittests

import "errors"

type GradeCalculator struct {
	submissions []Grade
}

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

var gradeTypeName = map[GradeType]string{
	Assignment: "assignment",
	Exam:       "exam",
	Essay:      "essay",
}

func (gt GradeType) String() string {
	return gradeTypeName[gt]
}

type Grade struct {
	Name  string
	Grade int
	Type  GradeType
}

func NewGradeCalculator() *GradeCalculator {
	return &GradeCalculator{
		submissions: make([]Grade, 0),
	}
}

func (gc *GradeCalculator) GetFinalGrade() string {
	numericalGrade := gc.calculateNumericalGrade()

	if numericalGrade >= 90 {
		return "A"
	} else if numericalGrade >= 80 {
		return "B"
	} else if numericalGrade >= 70 {
		return "C"
	} else if numericalGrade >= 60 {
		return "D"
	}
	return "F"
}

func (gc *GradeCalculator) AddGrade(name string, grade int, gradeType GradeType) error {
	g := Grade{name, grade, gradeType}
	switch gradeType {
	case Assignment:
		gc.submissions = append(gc.submissions, g)
	case Exam:
		gc.submissions = append(gc.submissions, g)
	case Essay:
		gc.submissions = append(gc.submissions, g)
	default:
		return errors.New("Invalid grade type")
	}
	return nil
}

func (gc *GradeCalculator) calculateNumericalGrade() int {
	list_assign := []Grade{}
	list_exam := []Grade{}
	list_essay := []Grade{}
	for _, task := range gc.submissions {
		switch task.Type {
		case Assignment:
			list_assign = append(list_assign, task)
		case Exam:
			list_exam = append(list_exam, task)
		case Essay:
			list_essay = append(list_essay, task)
		}
	}
	assignment_average, _ := computeAverage(list_assign)
	exam_average, _ := computeAverage(list_exam)
	essay_average, _ := computeAverage(list_essay)

	weighted_grade := float64(assignment_average)*.5 + float64(exam_average)*.35 + float64(essay_average)*.15

	return int(weighted_grade)
}

func computeAverage(grades []Grade) (int, error) {
	if len(grades) == 0 {
		return 0, errors.New("Grades cannot be empty!")
	}
	sum := 0

	for _, grade := range grades {
		sum += grade.Grade
	}
	return sum / len(grades), nil
}

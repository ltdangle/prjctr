package rating

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Individual grade.
type Grade float32

// Course grades.
type Rating struct {
	course   string
	grades   []Grade
	gradeAvg Grade
}

func (r *Rating) SetCourse(course string) {
	r.course = course
}

func (r *Rating) AddGrade(grade Grade) {
	r.grades = append(r.grades, grade)
}

func (r *Rating) AddGradeFromString(gradeStr string) error {
	grade, err := strconv.ParseFloat(gradeStr, 32)
	if err != nil {
		return errors.New("Wrong grade format. " + err.Error())
	}

	if grade < 0 {
		return errors.New("Grade cannot be negative.")
	}

	r.AddGrade(Grade(grade))

	return nil
}

func (r *Rating) CalculateGradeAverage() {
	var gradeSum Grade
	for _, grade := range r.grades {
		gradeSum += grade
	}

	r.gradeAvg = gradeSum / Grade(len(r.grades))
}

func (r *Rating) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	// Read course title.
	fmt.Print("Course title: ")
	scanner.Scan()
	r.course = scanner.Text()

	// Read grades in a loop.
	for {
		fmt.Print("Add grade (\"-1\" to finish): ")
		scanner.Scan()
		gradeStr := scanner.Text()

		if gradeStr == "-1" {
			break
		}

		err := r.AddGradeFromString(gradeStr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	r.CalculateGradeAverage()

	fmt.Printf("\nGrade average is %.2f\n", r.gradeAvg)
}

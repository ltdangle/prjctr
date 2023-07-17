package main

type School struct {
	Classes []*Class `json:"classes"`
}

func NewSchool() *School {
	return &School{}
}

func (s *School) AddClass(class *Class) {
	s.Classes = append(s.Classes, class)
}

func (s *School) FindClassById(id int) *Class {
	for _, class := range s.Classes {
		if class.Id == id {
			return class
		}
	}
	return nil
}

func (s *School) FindTeacher(username string) *Teacher {
	for _, class := range s.Classes {
		if class.Teacher.Username == username {
			return class.Teacher
		}
	}
	return nil
}

func (s *School) FindStudentById(studentId int) *Student {
	for _, class := range s.Classes {
		if student := class.FindById(studentId); student != nil {
			return student
		}
	}
	return nil
}

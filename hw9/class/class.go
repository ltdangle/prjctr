package main

type Class struct {
	Id       int        `json:"id"`
	Teacher  *Teacher   `json:"teacher"`
	Students []*Student `json:"students"`
}

func NewClass() *Class {
	return &Class{}
}

func (c *Class) SetTeacher(t *Teacher) {
	t.class = c
	c.Teacher = t
}
func (c *Class) AddStudent(s *Student) {
	s.class = c
	c.Students = append(c.Students, s)
}

func (c *Class) AllStudents() []*Student {
	return c.Students
}

func (c *Class) FindById(id int) *Student {
	for _, student := range c.Students {
		if student.Id == id {
			return student
		}
	}
	return nil
}

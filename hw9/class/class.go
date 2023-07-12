package main

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func NewStudent(id int, name string) *Student {
	return &Student{Id: id, Name: name}
}

type Class struct {
	Students []*Student `json:"students"`
}

func NewClass() *Class {
	return &Class{}
}

func (c *Class) addStudent(s *Student) {
	c.Students = append(c.Students, s)
}

func (c *Class) all() []*Student {
	return c.Students
}

func (c *Class) findById(id int) *Student {
	for _, student := range c.Students {
		if student.Id == id {
			return student
		}
	}
	return nil
}

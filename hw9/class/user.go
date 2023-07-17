package main

// User.
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// Student.
type Student struct {
	User
	class *Class
}

func NewStudent(id int, name string) *Student {
	return &Student{
		User: User{
			Id: id,
		},
	}
}

// Teacher.
type Teacher struct {
	User
	class *Class
}

func NewTeacher(id int, name string) *Teacher {
	return &Teacher{
		User: User{
			Id: id,
		},
	}
}

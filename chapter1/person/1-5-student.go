package person

type Student struct {
	name  string
	score float32
	Age   int
}

// 获取 name
func (s *Student) GetName() string {
	return s.name
}

// 设置 name
func (s *Student) SetName(newName string) {
	s.name = newName
}

// 获取age
func (s *Student) GetAge() int {
	return s.Age
}

// 设置age
func (s *Student) SetAge(newName int) {
	s.Age = newName
}

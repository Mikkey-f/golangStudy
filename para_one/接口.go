package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
	Skill
}

type Skill interface {
	GetSkillName() string
}

type BasketballSkill struct {
}

func (b *BasketballSkill) GetSkillName() string {
	return "basketball"
}

// School 嵌入接口
type School interface {
	Manage
	GetName() string
}

type Manage interface {
	Join(s *Student)
	FindStudentByName(name string) *Student
}

type MiddleSchool struct {
	Name     string
	Addr     string
	Students []*Student
}

func (m *MiddleSchool) Join(s *Student) {
	m.Students = append(m.Students, s)
}

func (m *MiddleSchool) FindStudentByName(name string) *Student {
	for i := range m.Students {
		if m.Students[i].Name == name {
			return m.Students[i]
		}
	}
	return nil
}

func (m *MiddleSchool) GetName() string {
	return m.Name
}

func main() {
	var school School
	school = &MiddleSchool{
		Name: "中学",
		Addr: "幸福小区",
	}
	name := school.GetName()
	fmt.Println(name)

	mikkeyf := &Student{
		Name: "mikkeyf",
		Age:  18,
	}
	mikkeyf.Skill = &BasketballSkill{}
	school.Join(mikkeyf)
	fmt.Println("s = ", school.FindStudentByName(mikkeyf.Name))
	fmt.Println("skill=", mikkeyf.GetSkillName())
}

package main

import (
	"log"
)

// 学校 ====================================
type School struct {
	sid        int
	name       string
	class_room map[int]*ClassRoom
	students   map[int]*Student
}

func (this *School) SetStudent(s *Student) {
	this.students[s.uid] = s
}

func (this *School) GetStudent(uid int) *Student {
	return this.students[uid]
}

func (this *School) SetClassRoom(c *ClassRoom) {
	this.class_room[c.cid] = c
}

func (this *School) GetClassRoom(cid int) *ClassRoom {
	return this.class_room[cid]
}

func (this *School) ChangeClassRoom(s *Student, cid int) {

}

// 教室 ====================================
type ClassRoom struct {
	cid      int
	students map[int]*Student
}

func (this *ClassRoom) SetStudent(s *Student) {
	this.students[s.uid] = s
}

func (this *ClassRoom) GetStudent(uid int) *Student {
	return this.students[uid]
}

// 学生 ====================================
type Student struct {
	uid           int
	class_room_id int
	name          string
	address       string
	age           int
}

func (this *Student) ChangeClassRoom(cid int) {
	this.class_room_id = cid
}

// run
func main() {
	log.Printf("mymap ... ")
	m := make(map[int]*Student)

	var s *Student = new(Student)
	s.uid = 101
	s.class_room_id = 1
	s.name = "Mikle"
	s.address = "红旗南路"
	s.age = 18

	uid := 101
	m[uid] = s

	log.Printf("student: ... ", m[uid].name, m[uid].address, m[uid].age)

}

package main

import (
	"log"
)

type School struct {
	sid        int
	name       string
	class_room map[int]*ClassRoom
}

type ClassRoom struct {
	cid     int
	student map[int]*Student
}

func (c *ClassRoom) Set(s *Student) {
	c.student[s.uid] = s
}

func (c *ClassRoom) Get(uid int) *Student {
	return c.student[uid]
}

type Student struct {
	uid           int
	class_room_id int
	name          string
	address       string
	age           int
}

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

package main

import (
	"log"
)

type Students struct {
	student map[int]*Student
}

func (ss *Students) Set(s *Student) {
	ss.student[s.uid] = s
}

func (ss *Students) Get(uid int) *Student {
	return ss.student[uid]
}

type Student struct {
	uid     int
	gid     int
	name    string
	address string
	age     int
}

func main() {
	log.Printf("mymap ... ")
	m := make(map[int]*Student)

	var s *Student = new(Student)
	s.uid = 101
	s.gid = 1
	s.name = "Mikle"
	s.address = "红旗南路"
	s.age = 18

	uid := 101
	m[uid] = s

	log.Printf("student: ... ", m[uid].name, m[uid].address, m[uid].age)

}

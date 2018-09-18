package main

import (
	"log"
)

type Student struct {
	id      int
	name    string
	address string
	age     int
}

func main() {
	log.Printf("mymap ... ")
	m := make(map[int]*Student)

	var s *Student = new(Student)
	s.id = 101
	s.name = "Mikle"
	s.address = "红旗南路"
	s.age = 18

	m[s.id] = s

	log.Printf("student: ... ", m[s.id].name, m[s.id].address, m[s.id].age)

}

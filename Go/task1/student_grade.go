package main

import "fmt"

type student struct {
	name    string
	courses map[string]float64
}

func (s *student) addCourse(course string, grade float64) {
	s.courses[course] = grade
	fmt.Printf("Course %v added successfuly \n", course)
}

func (s *student) setName(name string) {
	s.name = name
	fmt.Printf("Name %v set successfuly \n", name)
}

func (s *student) calAvarage() float64 {
	avarage := 0.0
	n := len(s.courses)
	for _, value := range s.courses {
		avarage += value
	}

	return (avarage / float64(n))
}

func (s student) printAvarage() {
	for course, result := range s.courses {
		fmt.Printf("%v ......... %v \n", course, result)
	}
	fmt.Printf("Avarage: ......... %v \n", s.calAvarage())
}

func (s *student) editName(name string) {
	s.name = name
	fmt.Printf("Student name updated to %v successfuly. \n", s.name)
}

func (s *student) editCourseScore(course string, score float64) {
	s.courses[course] = score
	fmt.Printf("%v score has been updated to %v \n", course, score)
}

func (s *student) editCourseName(prevName string, newName string) {
	grade := s.courses[prevName]
	delete(s.courses, prevName)
	s.courses[newName] = grade
	fmt.Printf("Update course name %v to %v successfuly \n", prevName, newName)
}

func (s *student) deleteCourse(name string) {
	delete(s.courses, name)
	fmt.Printf("Delete course name %v successfuly \n", name)
}

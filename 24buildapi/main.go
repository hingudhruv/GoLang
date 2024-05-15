package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"pricce"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware or helper
func (c *Course) IsEmpty() bool {
	// return c.CourseName == "" && c.CourseId == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to API creation")
	r := mux.NewRouter()
	courses = append(courses, Course{CourseId: "2", CourseName: "Maths", CoursePrice: 100, Author: &Author{Fullname: "Dhruv", Website: "dhruv.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "English", CoursePrice: 1000, Author: &Author{Fullname: "Mihir", Website: "mihir.com"}})

	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", addOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to the port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers -file
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome to Home page</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One course on ID")
	w.Header().Set("Content-type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses, find matching ID and return responses
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// output := "No course found with"+params["id"]
	json.NewEncoder(w).Encode("NO courses found")
}

func addOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ADD One course ")
	w.Header().Set("Content-type", "application/json")

	//what if body empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("please send some data")
		return
	}

	// what about this {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("NO data inside json")
		return
	}

	// generate unique id
	// append to courses
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update One course ")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	// loop, find ID, remove add again with the same id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One course ")
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	// loop, find ID, remove
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("course Deleted")
			return
		}
	}

}

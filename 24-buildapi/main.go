package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Modal for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to building API in golang")
	// creating router
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2",
		CourseName: "ReactJS", CoursePrice: 299,
		Author: &Author{Fullname: "Vaibhav Kushwaha", Website: "sketchomania.com"}})
	courses = append(courses, Course{CourseId: "4",
		CourseName: "NextJS", CoursePrice: 299,
		Author: &Author{Fullname: "Vaibhav Kushwaha", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	fmt.Printf("Type of params is %T\n", params)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("😶 No Course found with given id: " + params["id"])
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if; body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("😶 No data inside JSON")
		return
	}

	// TODO: check only if title is duplicate
	// loop through, stop when title matches with course.coursename, JSON response "coursename already exists"

	for _, singleCourse := range courses {
		// course.CourseName is from body
		if course.CourseName == singleCourse.CourseName {
			json.NewEncoder(w).Encode("Course name: " + course.CourseName + " already exist...😶")
			json.NewEncoder(w).Encode(singleCourse)
			return
		}
	}

	// generate unique id, string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop through, find id, remove, add with my ID

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// sending json response
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course) // decoding json based on this course
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode("😀 Course with id: " + params["id"] + " updated sucessfully")
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO: send a response when id is not found
	json.NewEncoder(w).Encode("😶 No course fournd with id: " + params["id"])
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop through, find id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			json.NewEncoder(w).Encode("😀 Course with id: " + params["id"] + " deleted sucessfully")
			break
		}
	}
}

func deleteAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all course")
	w.Header().Set("Content-Type", "application/json")
	// TODO: add delete all functionality
}

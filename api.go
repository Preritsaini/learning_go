package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId   string  `json:"courseId"`
	CourseName string  `json:"courseName"`
	CoursePic  int     `json:"coursePic"`
	Author     *Author `json:"authorId"`
}
type Author struct {
	fullname string `'json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {

}
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>WWW<h1>"))
}
func getAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Coutrses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["courseId"] {
			json.NewEncoder(w).Encode(course)
		}
	}

}
func CreateOneCourse(w http.ResponseWriter, r *http.Request) {

}

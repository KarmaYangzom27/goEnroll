package routes

import (
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	port := 8080
	router := mux.NewRouter()
	router.HandleFunc("/student", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{sid}", controller.GetStud).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")
	// Example route (You will later link these to your controller package)
	// router.HandleFunc("/api/test", controller.TestHandler).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT")
	router.HandleFunc("/students", controller.GetAllStuds)

	router.HandleFunc("/signup", controller.Signup).Methods("POST")
	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout)
	router.HandleFunc("/enroll", controller.Enroll).Methods("POST")
	router.HandleFunc("/enroll/{sid}/{cid}", controller.GetEnroll).Methods("GET")
	router.HandleFunc("/enroll/{sid}/{cid}", controller.DeleteEnroll).Methods("DELETE")
	router.HandleFunc("/enrolls", controller.GetEnrolls)

	router.HandleFunc("/course", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")
	router.HandleFunc("/courses", controller.GetAllCourses)
	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)
	log.Println("Application running on port", port)

	// Starting the server
	log.Fatal(http.ListenAndServe(":8080", router))
}

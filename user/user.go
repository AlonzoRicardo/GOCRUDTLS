package user

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const message = "Hello User!"

type User struct {
	gorm.Model
	Name  string
	Email string
}
type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) FetchUsers(w http.ResponseWriter, r *http.Request) {
	msg := "All users succesfully fetched"
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// var users []User
	// db.Find(&users)
	// fmt.Println("{}", users)

	// json.NewEncoder(w).Encode(users)

	// fmt.Fprintf(w, msg)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	msg := "New User Successfully Created"
	// fmt.Println("New User Endpoint Hit")

	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// vars := mux.Vars(r)
	// name := vars["name"]
	// email := vars["email"]

	// fmt.Println(name)
	// fmt.Println(email)

	// db.Create(&User{Name: name, Email: email})

	// fmt.Fprintf(w, msg)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func (h *Handlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	msg := "Successfully Deleted User"
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// vars := mux.Vars(r)
	// name := vars["name"]

	// var user User
	// db.Where("name = ?", name).Find(&user)
	// db.Delete(&user)

	// fmt.Fprintf(w, msg)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func (h *Handlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	msg := "User Succesfully updated"
	// db, err := gorm.Open("sqlite3", "test.db")
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer db.Close()

	// vars := mux.Vars(r)
	// name := vars["name"]
	// email := vars["email"]

	// var user User
	// db.Where("name = ?", name).Find(&user)

	// user.Email = email

	// db.Save(&user)
	// fmt.Fprintf(w, msg)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

// Logger configure logging
func (h *Handlers) Logger(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		handler(w, r)
	}
}

// SetupRoutes To actually enable the routes
func (h *Handlers) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/user/{user}", h.Logger(h.UpdateUser)).Methods("PUT")
	router.HandleFunc("/user/{user}", h.Logger(h.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/user/{user}", h.Logger(h.CreateUser)).Methods("POST")
	router.HandleFunc("/users", h.Logger(h.FetchUsers)).Methods("GET")
}

// NewHandlers Returns a Handler struct
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

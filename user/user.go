package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	uuid_gen "github.com/satori/go.uuid"
)

const message = "Hello User!"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	UUID     string
}
type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) FetchUsers(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	var users []User
	db.Find(&users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {
	body := decodeRequest(r)
	db := connectDB()
	defer db.Close()
	uuid := uuid_gen.NewV4().String()

	db.Create(&User{Name: body.Name, Email: body.Email, Password: body.Password, UUID: uuid})

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New User Successfully Created"))

}

func (h *Handlers) DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()

	vars := mux.Vars(r)
	uuid := vars["UUID"]

	var user User
	db.Where("UUID = ?", uuid).Find(&user)
	db.Delete(&user)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully Deleted User"))
}

func (h *Handlers) UpdateUser(w http.ResponseWriter, r *http.Request) {
	body := decodeRequest(r)
	db := connectDB()
	defer db.Close()

	var user User
	db.Where("UUID = ?", body.UUID).Find(&user)

	user.Name = body.Name
	user.Email = body.Email
	user.Password = body.Password

	db.Save(&user)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
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
	router.HandleFunc("/users", h.Logger(h.FetchUsers)).Methods("GET")
	router.HandleFunc("/user", h.Logger(h.UpdateUser)).Methods("PUT")
	router.HandleFunc("/user", h.Logger(h.CreateUser)).Methods("POST")
	router.HandleFunc("/user/{UUID}", h.Logger(h.DeleteUser)).Methods("DELETE")
}

// InitialMigration migrate user schema
func InitialMigration() {
	db, err := gorm.Open("sqlite3", "db/data/users.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

// NewHandlers Returns a Handler struct
func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}

func connectDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db/data/users.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func decodeRequest(r *http.Request) *User {
	var t User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
	if err != nil {
		panic("failed to decode user creation request")
	}
	return &t
}

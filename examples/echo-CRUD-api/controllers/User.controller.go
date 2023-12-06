package controllers

import (
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/jayantasamaddar/quick-reference-golang/echo-CRUD-api/fakedatabase"
	"github.com/jayantasamaddar/quick-reference-golang/echo-CRUD-api/models"
	"github.com/labstack/echo/v4"
)

type user struct{}
type methods interface {
	CreateUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetUser(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
}

var (
	User  methods = user{}
	mutex         = sync.Mutex{}
	users []*models.User
	db    fakedatabase.Storage = fakedatabase.Storage{Users: users}
)

func (u user) CreateUser(c echo.Context) error {
	mutex.Lock()
	defer mutex.Unlock()

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	// Generate default fields: New User
	user := &models.User{
		ID:       uuid,
		IsActive: true,
	}
	// Bind request body to the user struct
	err = c.Bind(&user)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Store in database
	response := db.AddRecord(user)

	// Send response to client
	return c.JSON(http.StatusCreated, response)
}

// Get All Users
func (u user) GetAllUsers(c echo.Context) error {
	response := db.GetAllRecords()

	// Send response to client
	return c.JSON(http.StatusOK, response)
}

// Get User
func (u user) GetUser(c echo.Context) error {
	// Parse id received from params into UUID
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println(err)
	}

	// Get the User from Database
	response := db.GetRecord(uuid)

	// Send response to client
	return c.JSON(http.StatusOK, response)
}

// Update User
func (u user) UpdateUser(c echo.Context) error {
	// Parse id received from params into UUID
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println(err)
	}

	// Generate default fields: New User
	user := &models.User{
		ID:       uuid,
		IsActive: true,
	}
	// Bind request body to the user struct
	err = c.Bind(&user)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	// Update the User in the Database
	response := db.UpdateRecord(uuid, user)

	// Send response to client
	return c.JSON(http.StatusOK, response)
}

// Delete User
func (u user) DeleteUser(c echo.Context) error {
	// Parse id received from params into UUID
	uuid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		log.Println(err)
	}
	// Get the User from Database
	response := db.DeleteRecord(uuid)

	// Send response to client
	return c.JSON(http.StatusOK, response)
}

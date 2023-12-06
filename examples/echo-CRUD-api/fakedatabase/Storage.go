package fakedatabase

import (
	"github.com/google/uuid"
	"github.com/jayantasamaddar/quick-reference-golang/echo-CRUD-api/models"
)

type Storage struct {
	Users []*models.User `json:"users"`
}

// Add a record to the database
func (db *Storage) AddRecord(data *models.User) *models.User {
	db.Users = append(db.Users, data)
	return db.Users[len(db.Users)-1]
}

// Get all Records
func (db *Storage) GetAllRecords() []*models.User {
	return db.Users
}

// Get a Single Record
func (db *Storage) GetRecord(id uuid.UUID) *models.User {
	for _, user := range db.Users {
		if user.ID == id {
			return user
		}
	}
	return nil
}

// Update Record
func (db *Storage) UpdateRecord(id uuid.UUID, data *models.User) *models.User {
	exists := db.GetRecord(id)
	if exists != nil {
		for i := range db.Users {
			if db.Users[i].ID == id {
				// Do not allow ID to be modified
				data.ID = id
				db.Users[i] = data
				return db.Users[i]
			}
		}
	}
	return nil
}

// Delete Record
func (db *Storage) DeleteRecord(id uuid.UUID) bool {
	for i := range db.Users {
		if db.Users[i].ID == id {
			if i == len(db.Users)-1 {
				db.Users = append(db.Users[:i])
			} else {
				db.Users = append(db.Users[:i], db.Users[:i+1]...)
			}
			return true
		}
	}
	return false
}

package migration

import (
	"github.com/tooth-fairy/core/patient"
	"github.com/tooth-fairy/infrastructure/database"
)

//Migration ... interface
type Migration interface {
	Migrations()
}

type migration struct {
	db *database.Database
}

//Migrations - create tables into databse
func (m *migration) Migrations() {
	g := m.db.GetGormClient()
	g.CreateTable(&patient.Patient{})
}

// New - responsible to create a new migration
func New(db *database.Database) Migration {
	return &migration{
		db: db,
	}
}

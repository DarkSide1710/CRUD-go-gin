package repository

import "github.com/jmoiron/sqlx"

type Repositories struct {
	db *sqlx.DB

	contactDB *contactDB
	taskDB    *taskDB
}

func New(db *sqlx.DB) *Repositories {
	return &Repositories{db: db}
}

func (s *Repositories) Contact() Contact {
	if s.contactDB == nil {
		s.contactDB = newContact(s.db)
	}
	return s.contactDB
}

func (s *Repositories) Task() Task {
	if s.taskDB == nil {
		s.taskDB = newTask(s.db)
	}
	return s.taskDB
}

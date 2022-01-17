package repository

type Repositories struct {
	contactDB *contactDB
	taskDB    *taskDB
}

func New() *Repositories {
	return &Repositories{}
}

func (s *Repositories) Contact() Contact {
	if s.contactDB == nil {
		s.contactDB = newContact()
	}
	return s.contactDB
}

func (s *Repositories) Task() Task {
	if s.taskDB == nil {
		s.taskDB = newTask()
	}
	return s.taskDB
}

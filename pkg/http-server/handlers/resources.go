package handlers

import "DarkSide1710/CRUD-go-gin/pkg/models"

type contactRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Position  string `json:"position" binding:"required"`
}
type taskRequest struct {
	Name      string `json:"name" binding:"required"`
	Status    string `json:"status" binding:"required"`
	Priority  string `json:"priority" binding:"required"`
	CreatedAt string `json:"created_at" binding:"required"`
	CreatedBy string `json:"created_by" binding:"required"`
	DueDate   string `json:"due_date" binding:"required"`
}

//validate
//func (c *contactRequest) validate() bool {
//	if len(c.Email) <= 3 {
//		return false
//	}
//
//	return true
//}

func (c *contactRequest) toModel() models.Contactlist {
	return models.Contactlist{
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Phone:     c.Phone,
		Email:     c.Email,
		Position:  c.Position,
	}
}

func (c *taskRequest) toModel_1() models.Tasklist {
	return models.Tasklist{
		Name:      c.Name,
		Status:    c.Status,
		Priority:  c.Priority,
		CreatedAt: c.CreatedAt,
		CreatedBy: c.CreatedBy,
		DueDate:   c.DueDate,
	}
}

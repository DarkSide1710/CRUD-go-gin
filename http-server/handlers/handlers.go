package handlers

import (
	services "DarkSide1710/CRUD-go-gin/service"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Platform services.Container
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRouter(platform services.Container) *gin.Engine {
	router := gin.Default()

	h := Handlers{Platform: platform}
	//t := HandlersT{Platform: platform}

	cList := router.Group("/contact")
	cList.GET("/", h.GetAllContacts)
	cList.GET("/:id", h.GetContact)
	cList.POST("/", h.CreateContact)
	cList.PUT("/:id", h.UpdateContact)
	cList.DELETE("/:id", h.DeleteContact)

	tList := router.Group("/task")
	tList.GET("/", h.GetAllTasks)
	tList.GET("/:id", h.GetTask)
	tList.POST("/", h.CreateTask)
	tList.PUT("/:id", h.UpdateTask)
	tList.DELETE("/:id", h.DeleteTask)

	router.GET("/", h.HelloWorld)

	return router
}

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

type HandlersT struct {
	Platform services.Container
}

type ResponseT struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRouter(platform services.Container) *gin.Engine {
	router := gin.Default()

	h := Handlers{Platform: platform}

	cList := router.Group("/contact")
	cList.GET("/", h.GetAllContacts)
	cList.GET("/:id", h.GetContact)
	cList.POST("/", h.CreateContact)
	cList.PUT("/:id", h.UpdateContact)
	cList.DELETE("/:id", h.DeleteContact)

	router.GET("/", h.HelloWorld)

	return router
}
func NewrouterT(platform services.Container) *gin.Engine {
	router := gin.Default()

	h := HandlersT{Platform: platform}

	tList := router.Group("/task")
	tList.GET("/", h.GetAllTasks)
	tList.GET("/:id", h.GetTask)
	tList.POST("/", h.CreateTask)
	tList.PUT("/:id", h.UpdateTask)
	tList.DELETE("/:id", h.DeleteTask)

	router.GET("/", h.HelloWorld)

	return router
}

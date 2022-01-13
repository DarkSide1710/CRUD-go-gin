package handlers

import (
	_ "example/web-service-gin/models"
	"example/web-service-gin/service"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Platform *services.Services
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type HandlersT struct {
	Platform *services.ServicesT
}

type ResponseT struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewRouter(platform *services.Services) *gin.Engine {
	router := gin.Default()

	h := Handlers{Platform: platform}

	cList := router.Group("/cList")
	cList.GET("/", h.GetAllContacts)
	cList.GET("/:id", h.GetContact)
	cList.POST("/", h.CreateContact)
	cList.PUT("/:id", h.UpdateContact)
	cList.DELETE("/:id", h.DeleteContact)

	router.GET("/", h.HelloWorld)

	return router
}
func NewrouterT(platform *services.ServicesT) *gin.Engine {
	router := gin.Default()

	h := HandlersT{Platform: platform}

	tList := router.Group("/t_list")
	tList.GET("/", h.GetAllTasks)
	tList.GET("/:id", h.GetTask)
	tList.POST("/", h.CreateTask)
	tList.PUT("/:id", h.UpdateTask)
	tList.DELETE("/:id", h.DeleteTask)

	router.GET("/", h.HelloWorld)

	return router
}

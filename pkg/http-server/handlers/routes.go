package handlers

import (
	_ "DarkSide1710/CRUD-go-gin/docs"
	"DarkSide1710/CRUD-go-gin/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handlers struct {
	Platform services.Container
}

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// @BasePath ./
func NewRouter(platform services.Container) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())

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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

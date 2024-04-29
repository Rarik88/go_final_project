package handler

import (
	"github.com/gin-gonic/gin"
	"github/Rarik88/go_final_project/pkg/api"
	"net/http"
)

type Handler struct {
	api *api.Api
}

func NewHandler(api *api.Api) *Handler {
	return &Handler{api: api}
}

func (h *Handler) Init() *gin.Engine {
	r := gin.Default()

	r.GET("/api/nextdate", h.ND)

	api := r.Group("/api")
	{
		api.POST("/task", h.AddTask)
		api.GET("/task", h.Task)
		api.GET("/tasks", h.Tasks)
		api.PUT("/task", h.UpdateTask)
		api.POST("/task/done", h.TaskDone)
		api.DELETE("/task", h.TaskDelete)
	}

	static := r.Group("/")
	{
		static.StaticFS("./css", http.Dir("./web/css"))
		static.StaticFS("./js", http.Dir("./web/js"))
	}

	r.GET("/", h.indexPage)
	r.StaticFile("/index.html", "./web/index.html")
	r.StaticFile("/login.html", "./web/login.html")
	r.StaticFile("/favicon.ico", "./web/favicon.ico")

	return r

}

func (h *Handler) indexPage(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "./web/index.html")
}

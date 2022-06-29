package routes

import (
	"github.com/PabloCassio/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriaNovoAluno)
	r.DELETE("/alunos/:id", controllers.DeletaAlunoPorId)
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	r.GET("/alunos/cpf/:cpf", controllers.ExibeAlunoPorCpf)
	r.GET("/index", controllers.ExibePaginaIndex)
	r.NoRoute(controllers.RotaNotFound)
	r.Run()

}

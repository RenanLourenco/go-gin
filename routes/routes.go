package routes

import (
	controller "github.com/RenanLourenco/go-gin.git/controllers"
	"github.com/gin-gonic/gin"
)


func HandleRequests(){
	r := gin.Default()
	r.GET("/alunos",controller.ExibeTodosAlunos)
	r.GET("/:nome",controller.Saudacao)
	r.GET("/aluno/:id",controller.ExibePorId)
	r.GET("/aluno/cpf/:cpf",controller.BuscaAlunoPorCPF)
	r.POST("/alunos",controller.CriaAluno)
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	r.PATCH("/alunos/:id", controller.EditaAluno)
	r.Run(":5000")
}
package controller

import (
	"net/http"

	"github.com/RenanLourenco/go-gin.git/database"
	"github.com/RenanLourenco/go-gin.git/models"
	"github.com/gin-gonic/gin"
)


func Saudacao(c *gin.Context){
	nome := c.Param("nome")
	c.JSON(200,gin.H{
		"API diz":"Olá " + nome + ", tranquilo?",
	})
}

func ExibePorId(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno,id)

	if aluno.ID == 0 {
		c.JSON(404,gin.H{
			"msg":"Aluno não encontrado",
		})
		return
	}


	c.JSON(200,aluno)
}

func ExibeTodosAlunos(c *gin.Context){
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200,alunos)
}

func CriaAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno,id)
	c.JSON(http.StatusOK, gin.H{
		"msg":"Aluno deletado com sucesso",
	})
}

func EditaAluno(c *gin.Context){
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno,id)


	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(http.StatusOK,aluno)


}

func BuscaAlunoPorCPF(c *gin.Context){
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(404,gin.H{
			"msg":"Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK,aluno)

}
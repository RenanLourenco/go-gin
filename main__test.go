package main

import (
	"bytes"
	"encoding/json"

	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/RenanLourenco/go-gin.git/controllers"
	"github.com/RenanLourenco/go-gin.git/database"
	"github.com/RenanLourenco/go-gin.git/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var id int
var cpf string

func CriaAlunoMock(){
	aluno := models.Aluno{
		Nome: "Aluno test",
		CPF: "12345678901",
		RG: "123456789",
	}
	database.DB.Create(&aluno)

	id = int(aluno.ID)
	cpf = aluno.CPF
}

func DeletaAlunoMock(){
	var aluno models.Aluno
	database.DB.Delete(&aluno,id)

}


func SetupDasRotasDeTeste() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeDaSaudacaoComParametro(t *testing.T){
	r := SetupDasRotasDeTeste()
	r.GET("/:nome",controller.Saudacao)
	req, _ := http.NewRequest("GET","/renan",nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp,req)

	assert.Equal(t, http.StatusOK, resp.Code,"Status error: valor recebido %d e o esperado era %d", resp.Code, http.StatusOK)
	mockDaResposta := `{"API diz":"Olá renan, tranquilo?"}`
	respostaBody, _ := io.ReadAll(resp.Body)

	assert.Equal(t, mockDaResposta, string(respostaBody))

}

func TestListandoTodosAlunosHandler(t *testing.T){
	database.ConectarDatabase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos",controller.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET","/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp,req)
	assert.Equal(t, http.StatusOK, resp.Code)

}

func TestBuscaAlunoPorCPFHandler(t *testing.T){
	database.ConectarDatabase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/aluno/cpf/:cpf",controller.BuscaAlunoPorCPF)
	req, _ := http.NewRequest("GET","/aluno/cpf/" + cpf,nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp,req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuscaAlunoPorIdHandler(t *testing.T){
	database.ConectarDatabase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/aluno/:id",controller.ExibePorId)
	path := "/aluno/" + strconv.Itoa(id)
	req, _ := http.NewRequest("GET",path,nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp,req)
	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(),&alunoMock)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Aluno test", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
}

func TestDeletaAlunoHandler(t *testing.T){
	database.ConectarDatabase()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controller.DeletaAluno)
	path := "/alunos/" + strconv.Itoa(id)
	req, _ := http.NewRequest("DELETE",path,nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp,req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestEditaAlunoHandler(t *testing.T){
	database.ConectarDatabase()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controller.EditaAluno)
	aluno := models.Aluno{Nome:"Testing patch",CPF: "12312312301",RG: "123123100"}
	valorJson, _ := json.Marshal(aluno)
	path := "/alunos/" + strconv.Itoa(id)
	req,_ := http.NewRequest("PATCH",path,bytes.NewBuffer(valorJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp,req)

	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(),&alunoMock)
	assert.Equal(t, "Testing patch", alunoMock.Nome)
	assert.Equal(t, "12312312301", alunoMock.CPF)
	assert.Equal(t, "123123100", alunoMock.RG)
	



}
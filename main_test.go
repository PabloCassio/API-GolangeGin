package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/PabloCassio/api-go-gin/controllers"
	"github.com/PabloCassio/api-go-gin/database"
	"github.com/PabloCassio/api-go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas

}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "12345678"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)

}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodedaSaudacaoComParametro(t *testing.T) {

	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/pablo", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "A resposta não foi a esperada, deveriam ser iguais.")
	mockDaResposta := `{"api diz:":"E aí pablo, tudo beleza?"}`
	respostaBody, _ := ioutil.ReadAll(resposta.Body)
	assert.Equal(t, mockDaResposta, string(respostaBody))
}

func TestListandoTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "A resposta não foi a esperada, deveriam ser iguais.")

}

func TestBuscaPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.ExibeAlunoPorCpf)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "A resposta não foi a esperada, deveriam ser iguais.")
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.ExibeAlunoPorId)
	req, _ := http.NewRequest("GET", "/alunos/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMock models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)
	assert.Equal(t, "Nome do Aluno Teste", alunoMock.Nome, "os nomes devem ser iguais")
	assert.Equal(t, "12345678901", alunoMock.CPF, "O CPF deve ser igual")
	assert.Equal(t, "12345678", alunoMock.RG, "O RG deve ser igual")
	assert.Equal(t, http.StatusOK, resposta.Code, "A resposta não foi a esperada, deveriam ser iguais.")
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAlunoPorId)
	req, _ := http.NewRequest("DELETE", "/alunos/"+strconv.Itoa(ID), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "A resposta não foi a esperada, deveriam ser iguais.")
}

func TestAtualizaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)
	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "47345678901", RG: "12345678"}
	valorJson, _ := json.Marshal(aluno)
	req, _ := http.NewRequest("PATCH", "/alunos/"+strconv.Itoa(ID), bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "47345678901", alunoMockAtualizado.CPF)
	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
	assert.Equal(t, "12345678", alunoMockAtualizado.RG)
	assert.Equal(t, http.StatusOK, resposta.Code, "A resposta não foi a esperada, deveriam ser iguais.")
}

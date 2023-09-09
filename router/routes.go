package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/bran00/TesteReactGolang/handler"
)

func initializeRoutes(router *gin.Engine, uri string) {
	basePath := "/api/v1"
	mongoURI := uri
	v1 := router.Group(basePath)
	{
		v1.GET("/openings", func(c *gin.Context) {
			// Chame handler.ShowAll para obter a resposta
			response := handler.ShowAll(mongoURI)

			// Verifique se ocorreu um erro na função ShowAll
			if response.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": response.Error.Error()})
				return
			}

			// Atribua os dados retornados pela função à variável data
			data := response.Response
			fmt.Println(data)
			// Envie os dados na resposta JSON
			c.JSON(http.StatusOK, gin.H{
				"message": "Documentos listados com sucesso",
				"Data":    data,
			})
		})

		v1.GET("/opening", func(c *gin.Context) {
			id := c.Query("_ID")
			// Chame handler.ShowDataByID para obter a resposta
			response, err := handler.ShowDataByID(mongoURI, id)

			// Verifique se ocorreu um erro na função ShowDataByID
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Atribua os dados retornados pela função à variável data
			data := response.Data

			// Verifique se a variável data não é nil antes de enviar na resposta JSON
			if data == nil {
				c.JSON(http.StatusNotFound, gin.H{"message": "Vaga não encontrada"})
				return
			}

			// Envie os dados na resposta JSON
			c.JSON(http.StatusOK, gin.H{
				"message": "Vaga listada com sucesso",
				"Data":    data,
			})
		})
	}

	v1.POST("/opening", func(c *gin.Context) {
		// Defina uma estrutura (struct) para receber os dados da solicitação JSON
		var requestData handler.InsertRequest

		// Faça o bind dos dados do corpo da solicitação para a estrutura
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Chame a função de inserção do manipulador com os dados recebidos
		response, err := handler.Insert(mongoURI, requestData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Retorne uma resposta JSON com a mensagem de sucesso ou erro
		c.JSON(http.StatusOK, gin.H{"message": response.Message})
	})

	v1.DELETE("/opening", func(c *gin.Context) {
		id := c.Query("_ID")

		// Chame handler.DeleteDocumentByID para obter a resposta
		err := handler.DeleteDocumentByID(mongoURI, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Vaga removida com sucesso",
			"Data":    err,
		})
	})
}

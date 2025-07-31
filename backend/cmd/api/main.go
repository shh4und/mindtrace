package main

import "github.com/gin-gonic/gin"

func main() {
	// Cria uma instância do Gin com configurações padrão
	r := gin.Default()

	// Define uma rota de teste. Quando alguém acessar a raiz ("/") do seu servidor...
	r.GET("/", func(c *gin.Context) {
		// o server respondera com:
		c.JSON(200, gin.H{
			"message": "O backend está funcionando corretamente!",
		})
	})

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}

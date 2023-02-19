package routers

import (
	"github.com/gin-gonic/gin"
	"mini-project/controllers"
	"mini-project/structs"
	"os"
)

var users = []structs.User{
	{"admin", "password"},
	{"editor", "secret"},
}

func basicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || (username != "admin" && username != "editor") || (password != "password" && password != "secret") {
			c.Header("WWW-Authenticate", `Basic realm="Authorization Required"`)
			c.AbortWithStatus(401)
			return
		}
		c.Next()
	}
}
func StartServer() *gin.Engine {
	r := gin.Default()

	r.GET("/bangun-datar/segitiga-sama-sisi", controllers.Segitiga)
	r.GET("/bangun-datar/persegi", controllers.Persegi)
	r.GET("/bangun-datar/persegi-panjang", controllers.PersegiPanjang)
	r.GET("/bangun-datar/lingkaran", controllers.Lingkaran)
	r.GET("/categories", controllers.GetCat)
	r.POST("/categories", basicAuth(), controllers.PostCat)
	r.PUT("/categories/:id", basicAuth(), controllers.UpdateCat)
	r.DELETE("/categories/:id", basicAuth(), controllers.DeleteCat)
	r.GET("/categories/:id/books", controllers.GetCatId)
	r.GET("/books", controllers.GetBooks)
	r.POST("/books", basicAuth(), controllers.PostBooks)
	r.PUT("/books/:id", basicAuth(), controllers.UpdateBooks)
	r.DELETE("/books/:id", basicAuth(), controllers.DeleteBooks)

	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		panic(err.Error())
	}
	return r
}

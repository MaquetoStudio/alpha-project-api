package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // Création d'un nouveau router avec le middleware par défaut

	// Middleware global (s'applique à toutes les routes)
	//r.Use(SomeGlobalMiddleware())

	// Route pour la page d'accueil ou la santé de l'API
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API en marche",
		})
	})

	// Groupement des routes pour les annonces
	//annonces := r.Group("/annonces")
	//{
	//	annonces.GET("/", GetAllAnnonces)
	//	annonces.POST("/", CreateAnnonce)
	//	annonces.GET("/:id", GetAnnonceByID)
	//	annonces.PUT("/:id", UpdateAnnonce)
	//	annonces.DELETE("/:id", DeleteAnnonce)
	//}
	//
	//// Groupement des routes pour les utilisateurs
	//utilisateurs := r.Group("/utilisateurs")
	//{
	//	utilisateurs.GET("/", GetAllUsers)
	//	utilisateurs.POST("/", CreateUser)
	//	utilisateurs.GET("/:username", GetUserByUsername)
	//	utilisateurs.PUT("/:username", UpdateUser)
	//	utilisateurs.DELETE("/:username", DeleteUser)
	//}
	//
	//// ... ajoutez d'autres groupes de routes ou routes individuelles selon les besoins

	return r
}

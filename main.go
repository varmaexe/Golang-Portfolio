package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vermaexe/go-jwt/controllers"
	"github.com/vermaexe/go-jwt/initializers"
	"github.com/vermaexe/go-jwt/middleware"
	"github.com/vermaexe/go-jwt/weather"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.StaticFS("/templates", http.Dir("./templates/"))
	r.LoadHTMLGlob("./templates/*.html")

	apiRoutes := r.Group("/api")
	{
		apiRoutes.POST("/signup", controllers.Signup)
		apiRoutes.POST("/login", controllers.Login)
		apiRoutes.GET("/logout", controllers.LogOut)
		apiRoutes.GET("/validate", middleware.RequireAuth, controllers.Validate)
		apiRoutes.POST("/contactform", controllers.Contactform)
		apiRoutes.POST("/weather", middleware.RequireAuth, weather.GetWeatherDataFromApi)

		apiRoutes.POST("/posts/create", controllers.PostsCreate)
		apiRoutes.GET("/allposts", controllers.PostsIndex)
		apiRoutes.GET("/postshow", controllers.PostShow)
		apiRoutes.POST("/postsupdate", controllers.PostUpdate)
		apiRoutes.POST("/posts/delete", controllers.PostsDelete)
	}
	viewRoutes := r.Group("/view")
	{
		viewRoutes.GET("/loginpage", controllers.Loginpage)
		viewRoutes.GET("/signuppage", controllers.Signinpage)
		viewRoutes.GET("/", controllers.Indexpage)
		viewRoutes.GET("/skills", controllers.Skillspage)
		viewRoutes.GET("/contact", controllers.Contactpage)
		viewRoutes.GET("/portfolio", controllers.Portfolio)
		viewRoutes.GET("/weatherhome", middleware.RequireAuth, controllers.WeatherHome)
		viewRoutes.GET("/posts/createpost", controllers.PostsCreatePage)
		viewRoutes.GET("/posts/homepage", controllers.PostPage)
		viewRoutes.GET("/postsupdate", controllers.PostUpdatePage)
		viewRoutes.GET("/post/delete", controllers.PostsDeletePage)

	}
	r.Run()
}

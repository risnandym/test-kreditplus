package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"final-project/controllers"
	"final-project/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register-admin", controllers.RegisterAdmin)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	AccountMiddlewareroute := r.Group("/change-password")
	AccountMiddlewareroute.Use(middlewares.PublicMiddleware())
	AccountMiddlewareroute.PATCH("", controllers.UpdatePassword)

	r.GET("/phones", controllers.GetAllPhone)
	r.GET("/phones/:id", controllers.GetPhoneById)
	r.GET("/phones/:id/specs-comments", controllers.GetSpecCommentByPhoneId)

	phonesMiddlewareroute := r.Group("/phones")
	phonesMiddlewareroute.Use(middlewares.AdminMiddleware())
	phonesMiddlewareroute.POST("/", controllers.CreatePhone)
	phonesMiddlewareroute.PATCH("/:id", controllers.UpdatePhone)
	phonesMiddlewareroute.DELETE("/:id", controllers.DeletePhone)

	r.GET("/brands", controllers.GetAllBrand)
	r.GET("/brands/:id", controllers.GetBrandById)
	r.GET("/brands/:id/phones", controllers.GetPhonesByBrandId)

	merkMiddlewareroute := r.Group("/brands")
	merkMiddlewareroute.Use(middlewares.AdminMiddleware())
	merkMiddlewareroute.POST("/", controllers.CreateBrand)
	merkMiddlewareroute.PATCH("/:id", controllers.UpdateBrand)
	merkMiddlewareroute.DELETE("/:id", controllers.DeleteBrand)

	r.GET("/specs", controllers.GetAllSpec)
	r.GET("/specs/:id", controllers.GetSpecById)

	specMiddlewareroute := r.Group("/specs")
	specMiddlewareroute.Use(middlewares.AdminMiddleware())
	specMiddlewareroute.POST("/", controllers.CreateSpec)
	specMiddlewareroute.PATCH("/:id", controllers.UpdateSpec)
	specMiddlewareroute.DELETE("/:id", controllers.DeleteSpec)

	r.GET("/news", controllers.GetAllNews)
	r.GET("/news/:id", controllers.GetNewsById)
	r.GET("/news/:id/comments", controllers.GetCommentByNewsId)

	newsMiddlewareroute := r.Group("/news")
	newsMiddlewareroute.Use(middlewares.AdminMiddleware())
	newsMiddlewareroute.POST("/", controllers.CreateNews)
	newsMiddlewareroute.PATCH("/:id", controllers.UpdateNews)
	newsMiddlewareroute.DELETE("/:id", controllers.DeleteNews)

	r.GET("/comments-phone", controllers.GetAllCommentPhone)
	// r.GET("/comments-phone/:id", controllers.GetCommentPhoneById)

	commentphoneMiddlewareroute := r.Group("/comments-phone")
	commentphoneMiddlewareroute.Use(middlewares.PublicMiddleware())
	commentphoneMiddlewareroute.POST("/", controllers.CreateCommentPhone)
	commentphoneMiddlewareroute.DELETE("/:id", controllers.DeleteCommentPhone)

	r.GET("/comments-news", controllers.GetAllCommentNews)
	// r.GET("/comments-news/:id", controllers.GetCommentNewsById)

	commentnewsMiddlewareroute := r.Group("/comments-news")
	commentnewsMiddlewareroute.Use(middlewares.PublicMiddleware())
	commentnewsMiddlewareroute.POST("/", controllers.CreateCommentNews)
	commentnewsMiddlewareroute.DELETE("/:id", controllers.DeleteCommentNews)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
